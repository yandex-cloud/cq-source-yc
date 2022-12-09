package servers

import (
	"context"
	"net"
	"testing"

	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/yandex-cloud/cq-provider-yandex/resources/testingutils/mocks"
	lockbox1 "github.com/yandex-cloud/go-genproto/yandex/cloud/lockbox/v1"
	lockbox "github.com/yandex-cloud/go-sdk/gen/lockboxsecret"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func StartLockboxServer(t *testing.T, ctx context.Context) (*lockbox.LockboxSecret, error) {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, err
	}

	serv := grpc.NewServer()
	go func() {
		<-ctx.Done()
		serv.Stop()
	}()

	err = registerLockboxMocks(t, serv)
	if err != nil {
		return nil, err
	}

	go func() {
		_ = serv.Serve(lis)
	}()

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return lockbox.NewLockboxSecret(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), nil
}

//go:generate mockgen -destination=../mocks/lockbox_secret_server_mock.go -package=mocks github.com/yandex-cloud/go-genproto/yandex/cloud/lockbox/v1 SecretServiceServer

func registerLockboxMocks(t *testing.T, serv *grpc.Server) error {
	ctrl := gomock.NewController(t)
	var err error
	faker.SetIgnoreInterface(true)

	var secret lockbox1.Secret
	err = faker.FakeData(&secret)
	if err != nil {
		return err
	}
	mLockboxSecret := mocks.NewMockSecretServiceServer(ctrl)
	mLockboxSecret.
		EXPECT().
		List(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *lockbox1.ListSecretsRequest) (*lockbox1.ListSecretsResponse, error) {
			if req == nil {
				return nil, status.Errorf(codes.Canceled, "request is nil")
			}
			if req.FolderId != "test-folder-id" {
				return nil, status.Errorf(codes.NotFound, "folder not found")
			}
			return &lockbox1.ListSecretsResponse{Secrets: []*lockbox1.Secret{&secret}}, nil
		}).
		AnyTimes()
	lockbox1.RegisterSecretServiceServer(serv, mLockboxSecret)
	return nil
}