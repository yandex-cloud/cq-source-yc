package resources_test

import (
	"context"
	"errors"
	"fmt"
	"net"
	"testing"

	"google.golang.org/grpc"

	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/faker/v3"
	"github.com/hashicorp/go-hclog"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/cq-provider-yandex/resources"
	containerregistry1 "github.com/yandex-cloud/go-genproto/yandex/cloud/containerregistry/v1"
	"github.com/yandex-cloud/go-sdk/gen/containerregistry"
)

func TestContainerRegistryScanResults(t *testing.T) {
	containerregistrySvc, serv, err := createScanResultServer()
	if err != nil {
		t.Fatal(err)
	}
	resource := providertest.ResourceTestData{
		Table: resources.ContainerRegistryScanResults(),
		Config: client.Config{
			FolderIDs: []string{"test"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			c := client.NewYandexClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"test"}, nil, nil, &client.Services{
				ContainerRegistry: containerregistrySvc,
			}, nil)
			return c, nil
		},
		Verifiers: []providertest.Verifier{
			providertest.VerifyAtLeastOneRow("yandex_containerregistry_scan_results"),
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
	serv.Stop()
}

type FakeContainerRegistryImageServer struct {
	containerregistry1.UnimplementedImageServiceServer
	Image *containerregistry1.Image
}

func NewFakeContainerRegistryImageServer() (*FakeContainerRegistryImageServer, error) {
	var image containerregistry1.Image
	faker.SetIgnoreInterface(true)
	err := faker.FakeData(&image)
	if err != nil {
		return nil, err
	}
	image.Id = "test"
	return &FakeContainerRegistryImageServer{Image: &image}, nil
}

func (s *FakeContainerRegistryImageServer) List(context.Context, *containerregistry1.ListImagesRequest) (*containerregistry1.ListImagesResponse, error) {
	return &containerregistry1.ListImagesResponse{Images: []*containerregistry1.Image{s.Image}}, nil
}

type FakeScanResultServiceServer struct {
	containerregistry1.UnimplementedScannerServiceServer
	ScanResult *containerregistry1.ScanResult
}

func NewFakeScanResultServiceServer() (*FakeScanResultServiceServer, error) {
	var scanResult containerregistry1.ScanResult
	faker.SetIgnoreInterface(true)
	err := faker.FakeData(&scanResult)
	if err != nil {
		return nil, err
	}
	return &FakeScanResultServiceServer{ScanResult: &scanResult}, nil
}

func (s *FakeScanResultServiceServer) List(_ context.Context, req *containerregistry1.ListScanResultsRequest) (*containerregistry1.ListScanResultsResponse, error) {
	if imageId, ok := req.Id.(*containerregistry1.ListScanResultsRequest_ImageId); ok && imageId.ImageId == "test" {
		return &containerregistry1.ListScanResultsResponse{ScanResults: []*containerregistry1.ScanResult{s.ScanResult}}, nil
	} else {
		return nil, errors.New("scan results not found")
	}
}

func createScanResultServer() (*containerregistry.ContainerRegistry, *grpc.Server, error) {
	lis, err := net.Listen("tcp", "")
	if err != nil {
		return nil, nil, err
	}

	serv := grpc.NewServer()

	fakeContainerRegistryImageServer, err := NewFakeContainerRegistryImageServer()
	if err != nil {
		return nil, nil, err
	}
	containerregistry1.RegisterImageServiceServer(serv, fakeContainerRegistryImageServer)

	fakeScanResultServiceServer, err := NewFakeScanResultServiceServer()
	if err != nil {
		return nil, nil, err
	}
	containerregistry1.RegisterScannerServiceServer(serv, fakeScanResultServiceServer)

	go func() {
		err := serv.Serve(lis)
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())

	if err != nil {
		return nil, nil, err
	}

	return containerregistry.NewContainerRegistry(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), serv, nil
}
