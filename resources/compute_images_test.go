package resources_test

import (
	"context"
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
	compute1 "github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	"github.com/yandex-cloud/go-sdk/gen/compute"
)

func TestComputeImages(t *testing.T) {
	computeSvc, serv, err := createImageServer()
	if err != nil {
		t.Fatal(err)
	}
	resource := providertest.ResourceTestData{
		Table: resources.ComputeImages(),
		Config: client.Config{
			FolderIDs: []string{"test"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			c := client.NewYandexClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"test"}, nil, nil, &client.Services{
				Compute: computeSvc,
			}, nil)
			return c, nil
		},
		Verifiers: []providertest.Verifier{
			providertest.VerifyAtLeastOneRow("yandex_compute_images"),
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
	serv.Stop()
}

type FakeImageServiceServer struct {
	compute1.UnimplementedImageServiceServer
	Image *compute1.Image
}

func NewFakeImageServiceServer() (*FakeImageServiceServer, error) {
	var image compute1.Image
	faker.SetIgnoreInterface(true)
	err := faker.FakeData(&image)
	if err != nil {
		return nil, err
	}
	return &FakeImageServiceServer{Image: &image}, nil
}

func (s *FakeImageServiceServer) List(context.Context, *compute1.ListImagesRequest) (*compute1.ListImagesResponse, error) {
	return &compute1.ListImagesResponse{Images: []*compute1.Image{s.Image}}, nil
}

func createImageServer() (*compute.Compute, *grpc.Server, error) {
	lis, err := net.Listen("tcp", ":0")

	if err != nil {
		return nil, nil, err
	}

	serv := grpc.NewServer()
	fakeImageServiceServer, err := NewFakeImageServiceServer()

	if err != nil {
		return nil, nil, err
	}

	compute1.RegisterImageServiceServer(serv, fakeImageServiceServer)

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

	return compute.NewCompute(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), serv, nil
}
