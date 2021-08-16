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

func TestComputeDisks(t *testing.T) {
	computeSvc, serv, err := createDiskServer()
	if err != nil {
		t.Fatal(err)
	}
	resource := providertest.ResourceTestData{
		Table: resources.ComputeDisks(),
		Config: client.Config{
			FolderIDs: []string{"testFolder"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			c := client.NewYandexClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testFolder"}, nil, &client.Services{
				Compute: computeSvc,
			}, nil)
			return c, nil
		},
		Verifiers: []providertest.Verifier{
			providertest.VerifyNoEmptyColumnsExcept("yandex_compute_disks", "source_source_image_id", "source_source_snapshot_id"),
			providertest.VerifyOneOf("yandex_compute_disks", "source_source_image_id", "source_source_snapshot_id"),
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
	serv.Stop()
}

type FakeDiskServiceServer struct {
	compute1.UnimplementedDiskServiceServer
	Disk *compute1.Disk
}

func NewFakeDiskServiceServer() (*FakeDiskServiceServer, error) {
	var disk compute1.Disk
	faker.SetIgnoreInterface(true)
	err := faker.FakeData(&disk)
	if err != nil {
		return nil, err
	}
	return &FakeDiskServiceServer{Disk: &disk}, nil
}

func (s *FakeDiskServiceServer) List(context.Context, *compute1.ListDisksRequest) (*compute1.ListDisksResponse, error) {
	return &compute1.ListDisksResponse{Disks: []*compute1.Disk{s.Disk}}, nil
}

func createDiskServer() (*compute.Compute, *grpc.Server, error) {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		return nil, nil, err
	}

	serv := grpc.NewServer()
	fakeDiskServiceServer, err := NewFakeDiskServiceServer()

	if err != nil {
		return nil, nil, err
	}

	compute1.RegisterDiskServiceServer(serv, fakeDiskServiceServer)

	go func() {
		err := serv.Serve(lis)
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		return nil, nil, err
	}

	return compute.NewCompute(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), serv, nil
}
