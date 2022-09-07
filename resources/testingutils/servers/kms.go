package servers

import (
	"context"
	"net"
	"testing"

	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/yandex-cloud/cq-provider-yandex/resources/testingutils/mocks"
	kms1 "github.com/yandex-cloud/go-genproto/yandex/cloud/kms/v1"
	"github.com/yandex-cloud/go-sdk/gen/kms"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func StartKmsServer(t *testing.T, ctx context.Context) (*kms.KMS, error) {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, err
	}

	serv := grpc.NewServer()
	go func() {
		<-ctx.Done()
		serv.Stop()
	}()

	err = registerKmsMocks(t, serv)
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

	return kms.NewKMS(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), nil
}

//go:generate mockgen -destination=../mocks/kms_symmetric_key_service_server_mock.go -package=mocks github.com/yandex-cloud/go-genproto/yandex/cloud/kms/v1 SymmetricKeyServiceServer

func registerKmsMocks(t *testing.T, serv *grpc.Server) error {
	ctrl := gomock.NewController(t)
	var err error
	faker.SetIgnoreInterface(true)

	var symmetricKey kms1.SymmetricKey
	err = faker.FakeData(&symmetricKey)
	if err != nil {
		return err
	}
	mSymmetricKey := mocks.NewMockSymmetricKeyServiceServer(ctrl)
	mSymmetricKey.
		EXPECT().
		List(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *kms1.ListSymmetricKeysRequest) (*kms1.ListSymmetricKeysResponse, error) {
			if req == nil {
				return nil, status.Errorf(codes.Canceled, "request is nil")
			}
			if req.FolderId != "test-folder-id" {
				return nil, status.Errorf(codes.NotFound, "folder not found")
			}
			return &kms1.ListSymmetricKeysResponse{Keys: []*kms1.SymmetricKey{&symmetricKey}}, nil
		}).
		AnyTimes()
	kms1.RegisterSymmetricKeyServiceServer(serv, mSymmetricKey)

	return nil
}
