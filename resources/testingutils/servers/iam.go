package servers

import (
	"context"
	"net"
	"testing"

	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/yandex-cloud/cq-provider-yandex/resources/testingutils/mocks"
	iam1 "github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
	"github.com/yandex-cloud/go-sdk/gen/iam"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func StartIamServer(t *testing.T, ctx context.Context) (*iam.IAM, error) {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, err
	}

	serv := grpc.NewServer()
	go func() {
		<-ctx.Done()
		serv.Stop()
	}()

	err = registerIamMocks(t, serv)
	if err != nil {
		return nil, err
	}

	go func() {
		_ = serv.Serve(lis)
	}()

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return iam.NewIAM(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), nil
}

//go:generate mockgen -destination=../mocks/iam_service_account_service_server_mock.go -package=mocks github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1 ServiceAccountServiceServer
//go:generate mockgen -destination=../mocks/iam_user_account_service_server_mock.go -package=mocks github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1 UserAccountServiceServer

func registerIamMocks(t *testing.T, serv *grpc.Server) error {
	ctrl := gomock.NewController(t)
	var err error
	faker.SetIgnoreInterface(true)

	var serviceAccount iam1.ServiceAccount
	err = faker.FakeData(&serviceAccount)
	if err != nil {
		return err
	}
	mServiceAccount := mocks.NewMockServiceAccountServiceServer(ctrl)
	mServiceAccount.
		EXPECT().
		List(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *iam1.ListServiceAccountsRequest) (*iam1.ListServiceAccountsResponse, error) {
			if req == nil {
				return nil, status.Errorf(codes.Canceled, "request is nil")
			}
			if req.FolderId != "test-folder-id" {
				return nil, status.Errorf(codes.NotFound, "folder not found")
			}
			return &iam1.ListServiceAccountsResponse{ServiceAccounts: []*iam1.ServiceAccount{&serviceAccount}}, nil
		}).
		AnyTimes()
	iam1.RegisterServiceAccountServiceServer(serv, mServiceAccount)

	var userAccount iam1.UserAccount
	err = faker.FakeData(&userAccount)
	if err != nil {
		return err
	}
	mUserAccount := mocks.NewMockUserAccountServiceServer(ctrl)
	mUserAccount.
		EXPECT().
		Get(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *iam1.GetUserAccountRequest) (*iam1.UserAccount, error) {
			if req == nil {
				return nil, status.Errorf(codes.Canceled, "request is nil")
			}
			return &userAccount, nil
		}).
		AnyTimes()
	iam1.RegisterUserAccountServiceServer(serv, mUserAccount)

	return nil
}
