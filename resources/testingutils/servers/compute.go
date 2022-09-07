package servers

import (
	"context"
	"net"
	"testing"

	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/yandex-cloud/cq-provider-yandex/resources/testingutils/mocks"
	compute1 "github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	"github.com/yandex-cloud/go-sdk/gen/compute"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func StartComputeServer(t *testing.T, ctx context.Context) (*compute.Compute, error) {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, err
	}

	serv := grpc.NewServer()
	go func() {
		<-ctx.Done()
		serv.Stop()
	}()

	err = registerComputeMocks(t, serv)
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

	return compute.NewCompute(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), nil
}

//go:generate mockgen -destination=../mocks/compute_disk_service_server_mock.go -package=mocks github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1 DiskServiceServer
//go:generate mockgen -destination=../mocks/compute_image_service_server_mock.go -package=mocks github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1 ImageServiceServer
//go:generate mockgen -destination=../mocks/compute_instance_service_server_mock.go -package=mocks github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1 InstanceServiceServer

func registerComputeMocks(t *testing.T, serv *grpc.Server) error {
	ctrl := gomock.NewController(t)
	var err error
	faker.SetIgnoreInterface(true)

	var disk compute1.Disk
	err = faker.FakeData(&disk)
	t.Log("===== DISK:")
	t.Log(disk)
	if err != nil {
		return err
	}
	mDiskServ := mocks.NewMockDiskServiceServer(ctrl)
	mDiskServ.
		EXPECT().
		List(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *compute1.ListDisksRequest) (*compute1.ListDisksResponse, error) {
			if req == nil {
				return nil, status.Errorf(codes.Canceled, "request is nil")
			}
			if req.FolderId != "test-folder-id" {
				return nil, status.Errorf(codes.NotFound, "folder not found")
			}
			return &compute1.ListDisksResponse{Disks: []*compute1.Disk{&disk}}, nil
		}).
		AnyTimes()
	compute1.RegisterDiskServiceServer(serv, mDiskServ)

	var image compute1.Image
	faker.SetIgnoreInterface(true)
	err = faker.FakeData(&image)
	if err != nil {
		return err
	}
	mImageServ := mocks.NewMockImageServiceServer(ctrl)
	mImageServ.
		EXPECT().
		List(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *compute1.ListImagesRequest) (*compute1.ListImagesResponse, error) {
			if req == nil {
				return nil, status.Errorf(codes.Canceled, "request is nil")
			}
			if req.FolderId != "test-folder-id" {
				return nil, status.Errorf(codes.NotFound, "folder not found")
			}
			return &compute1.ListImagesResponse{Images: []*compute1.Image{&image}}, nil
		}).
		AnyTimes()
	compute1.RegisterImageServiceServer(serv, mImageServ)

	var instance compute1.Instance
	faker.SetIgnoreInterface(true)
	err = faker.FakeData(&instance)
	if err != nil {
		return err
	}
	mInstanceServ := mocks.NewMockInstanceServiceServer(ctrl)
	mInstanceServ.
		EXPECT().
		List(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *compute1.ListInstancesRequest) (*compute1.ListInstancesResponse, error) {
			if req == nil {
				return nil, status.Errorf(codes.Canceled, "request is nil")
			}
			if req.FolderId != "test-folder-id" {
				return nil, status.Errorf(codes.NotFound, "folder not found")
			}
			return &compute1.ListInstancesResponse{Instances: []*compute1.Instance{&instance}}, nil
		}).
		AnyTimes()
	compute1.RegisterInstanceServiceServer(serv, mInstanceServ)

	return nil
}
