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
	resourcemanager1 "github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
	"github.com/yandex-cloud/go-sdk/gen/resourcemanager"
)

func TestResourceManagerClouds(t *testing.T) {
	var serv *grpc.Server
	resource := providertest.ResourceTestData{
		Table: resources.ResourceManagerClouds(),
		Config: client.Config{
			CloudIDs: []string{"testCloud"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			resourcemanagerSvc, serv1, err := createCloudServer()
			serv = serv1
			if err != nil {
				return nil, err
			}
			c := client.NewYandexClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testCloud"}, nil, &client.Services{
				ResourceManager: resourcemanagerSvc,
			}, nil)
			return c, nil
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
	serv.Stop()
}

type FakeCloudServiceServer struct {
	resourcemanager1.UnimplementedCloudServiceServer
	Cloud *resourcemanager1.Cloud
}

func NewFakeCloudServiceServer() (*FakeCloudServiceServer, error) {
	var Cloud resourcemanager1.Cloud
	faker.SetIgnoreInterface(true)
	err := faker.FakeData(&Cloud)
	if err != nil {
		return nil, err
	}
	Cloud.Name = "testCloud"
	return &FakeCloudServiceServer{Cloud: &Cloud}, nil
}

func (s *FakeCloudServiceServer) Get(_ context.Context, req *resourcemanager1.GetCloudRequest) (*resourcemanager1.Cloud, error) {
	if req.CloudId == "testCloud" {
		return s.Cloud, nil
	}
	return nil, errors.New("no such Cloud")
}

func createCloudServer() (*resourcemanager.ResourceManager, *grpc.Server, error) {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		return nil, nil, err
	}

	serv := grpc.NewServer()
	fakeCloudServiceServer, err := NewFakeCloudServiceServer()

	if err != nil {
		return nil, nil, err
	}

	resourcemanager1.RegisterCloudServiceServer(serv, fakeCloudServiceServer)

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

	return resourcemanager.NewResourceManager(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), serv, nil
}
