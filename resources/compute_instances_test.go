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

func TestComputeInstances(t *testing.T) {
	computeSvc, serv, err := createInstanceServer()
	if err != nil {
		t.Fatal(err)
	}
	resource := providertest.ResourceTestData{
		Table: resources.ComputeInstances(),
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
	}
	providertest.TestResource(t, resources.Provider, resource)
	serv.Stop()
}

type FakeInstanceServiceServer struct {
	compute1.UnimplementedInstanceServiceServer
	Instance *compute1.Instance
}

func NewFakeInstanceServiceServer() (*FakeInstanceServiceServer, error) {
	var instance compute1.Instance
	faker.SetIgnoreInterface(true)
	err := faker.FakeData(&instance)
	if err != nil {
		return nil, err
	}
	return &FakeInstanceServiceServer{Instance: &instance}, nil
}

func (s *FakeInstanceServiceServer) List(context.Context, *compute1.ListInstancesRequest) (*compute1.ListInstancesResponse, error) {
	return &compute1.ListInstancesResponse{Instances: []*compute1.Instance{s.Instance}}, nil
}

func createInstanceServer() (*compute.Compute, *grpc.Server, error) {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		return nil, nil, err
	}

	serv := grpc.NewServer()
	fakeInstanceServiceServer, err := NewFakeInstanceServiceServer()

	if err != nil {
		return nil, nil, err
	}

	compute1.RegisterInstanceServiceServer(serv, fakeInstanceServiceServer)

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
