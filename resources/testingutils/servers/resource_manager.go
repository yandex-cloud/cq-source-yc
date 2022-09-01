package servers

import (
	"context"
	"net"
	"testing"

	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/yandex-cloud/cq-provider-yandex/resources/testingutils/mocks"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	resourcemanager1 "github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
	"github.com/yandex-cloud/go-sdk/gen/resourcemanager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func StartResourceManagerServer(t *testing.T, ctx context.Context) (*resourcemanager.ResourceManager, error) {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, err
	}

	serv := grpc.NewServer()
	go func() {
		<-ctx.Done()
		serv.Stop()
	}()

	err = registerResourceManagerMocks(t, serv)
	if err != nil {
		return nil, err
	}

	go func() {
		_ = serv.Serve(lis)
	}()

	conn, err := grpc.Dial(lis.Addr().String(), insecure.NewCredentials())
	if err != nil {
		return nil, err
	}

	return resourcemanager.NewResourceManager(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), nil
}

//go:generate mockgen -destination=../mocks/resource_manager_folder_service_server_mock.go -package=mocks github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1 FolderServiceServer
//go:generate mockgen -destination=../mocks/resource_manager_cloud_service_server_mock.go -package=mocks github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1 CloudServiceServer

func registerResourceManagerMocks(t *testing.T, serv *grpc.Server) error {
	ctrl := gomock.NewController(t)
	var err error
	faker.SetIgnoreInterface(true)

	var folder resourcemanager1.Folder
	err = faker.FakeData(&folder)
	if err != nil {
		return err
	}
	mFolderServ := mocks.NewMockFolderServiceServer(ctrl)
	mFolderServ.
		EXPECT().
		List(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *resourcemanager1.ListFoldersRequest) (*resourcemanager1.ListFoldersResponse, error) {
			if req == nil {
				return nil, status.Errorf(codes.Canceled, "request is nil")
			}
			if req.CloudId != "test-cloud-id" {
				return nil, status.Errorf(codes.NotFound, "folder not found")
			}
			return &resourcemanager1.ListFoldersResponse{Folders: []*resourcemanager1.Folder{&folder}}, nil
		}).
		AnyTimes()
	var accessBindingByFolder access.AccessBinding
	err = faker.FakeData(&accessBindingByFolder)
	if err != nil {
		return err
	}
	mFolderServ.
		EXPECT().
		ListAccessBindings(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
			if req == nil {
				return nil, status.Errorf(codes.Canceled, "request is nil")
			}
			return &access.ListAccessBindingsResponse{AccessBindings: []*access.AccessBinding{&accessBindingByFolder}}, nil
		}).
		AnyTimes()
	resourcemanager1.RegisterFolderServiceServer(serv, mFolderServ)

	var cloud resourcemanager1.Cloud
	err = faker.FakeData(&cloud)
	if err != nil {
		return err
	}
	mCloudServ := mocks.NewMockCloudServiceServer(ctrl)
	mCloudServ.
		EXPECT().
		List(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *resourcemanager1.ListCloudsRequest) (*resourcemanager1.ListCloudsResponse, error) {
			if req == nil {
				return nil, status.Errorf(codes.Canceled, "request is nil")
			}
			return &resourcemanager1.ListCloudsResponse{Clouds: []*resourcemanager1.Cloud{&cloud}}, nil
		}).
		AnyTimes()
	var accessBindingByCloud access.AccessBinding
	err = faker.FakeData(&accessBindingByCloud)
	if err != nil {
		return err
	}
	mCloudServ.
		EXPECT().
		ListAccessBindings(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
			if req == nil {
				return nil, status.Errorf(codes.Canceled, "request is nil")
			}
			return &access.ListAccessBindingsResponse{AccessBindings: []*access.AccessBinding{&accessBindingByCloud}}, nil
		}).
		AnyTimes()
	resourcemanager1.RegisterCloudServiceServer(serv, mCloudServ)

	return nil
}
