package servers

import (
	"context"
	"net"
	"testing"

	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/yandex-cloud/cq-provider-yandex/resources/testingutils/mocks"
	containerregistry1 "github.com/yandex-cloud/go-genproto/yandex/cloud/containerregistry/v1"
	"github.com/yandex-cloud/go-sdk/gen/containerregistry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func StartContainerRegistryServer(t *testing.T, ctx context.Context) (*containerregistry.ContainerRegistry, error) {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, err
	}

	serv := grpc.NewServer()
	go func() {
		<-ctx.Done()
		serv.Stop()
	}()

	err = registerContainerRegistryMocks(t, serv)
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

	return containerregistry.NewContainerRegistry(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), nil
}

//go:generate mockgen -destination=../mocks/container_registry_image_service_server_mock.go -package=mocks -mock_names=ImageServiceServer=MockContainerRegistryImageServiceServer github.com/yandex-cloud/go-genproto/yandex/cloud/containerregistry/v1 ImageServiceServer
//go:generate mockgen -destination=../mocks/container_registry_scan_result_service_server_mock.go -package=mocks github.com/yandex-cloud/go-genproto/yandex/cloud/containerregistry/v1 ScannerServiceServer

func registerContainerRegistryMocks(t *testing.T, serv *grpc.Server) error {
	ctrl := gomock.NewController(t)
	var err error
	faker.SetIgnoreInterface(true)

	var image containerregistry1.Image
	err = faker.FakeData(&image)
	if err != nil {
		return err
	}
	mImageServ := mocks.NewMockContainerRegistryImageServiceServer(ctrl)
	mImageServ.
		EXPECT().
		List(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *containerregistry1.ListImagesRequest) (*containerregistry1.ListImagesResponse, error) {
			if req == nil {
				return nil, status.Errorf(codes.Canceled, "request is nil")
			}
			if req.FolderId != "test-folder-id" {
				return nil, status.Errorf(codes.NotFound, "folder not found")
			}
			return &containerregistry1.ListImagesResponse{Images: []*containerregistry1.Image{&image}}, nil
		}).
		AnyTimes()
	containerregistry1.RegisterImageServiceServer(serv, mImageServ)

	var scanResult containerregistry1.ScanResult
	err = faker.FakeData(&scanResult)
	if err != nil {
		return err
	}
	mScannerServ := mocks.NewMockScannerServiceServer(ctrl)
	mScannerServ.
		EXPECT().
		List(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *containerregistry1.ListScanResultsRequest) (*containerregistry1.ListScanResultsResponse, error) {
			if req == nil {
				return nil, status.Errorf(codes.Canceled, "request is nil")
			}
			return &containerregistry1.ListScanResultsResponse{ScanResults: []*containerregistry1.ScanResult{&scanResult}}, nil
		}).
		AnyTimes()
	containerregistry1.RegisterScannerServiceServer(serv, mScannerServ)

	return nil
}
