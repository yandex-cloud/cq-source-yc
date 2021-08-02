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
	vpc1 "github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
	"github.com/yandex-cloud/go-sdk/gen/vpc"
)

func TestVPCNetworks(t *testing.T) {
	var serv *grpc.Server
	resource := providertest.ResourceTestData{
		Table: resources.VPCNetworks(),
		Config: client.Config{
			FolderIDs: []string{"testFolder"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			vpcSvc, serv1, err := createNetworkServer()
			serv = serv1
			if err != nil {
				return nil, err
			}
			c := client.NewYandexClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testFolder"}, &client.Services{
				VPC: vpcSvc,
			})
			return c, nil
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
	serv.Stop()
}

type FakeNetworkServiceServer struct {
	vpc1.UnimplementedNetworkServiceServer
	Network *vpc1.Network
}

func NewFakeNetworkServiceServer() (*FakeNetworkServiceServer, error) {
	var network vpc1.Network
	faker.SetIgnoreInterface(true)
	err := faker.FakeData(&network)
	if err != nil {
		return nil, err
	}
	return &FakeNetworkServiceServer{Network: &network}, nil
}

func (s *FakeNetworkServiceServer) List(context.Context, *vpc1.ListNetworksRequest) (*vpc1.ListNetworksResponse, error) {
	return &vpc1.ListNetworksResponse{Networks: []*vpc1.Network{s.Network}}, nil
}

func createNetworkServer() (*vpc.VPC, *grpc.Server, error) {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		return nil, nil, err
	}

	serv := grpc.NewServer()
	fakeNetworkServiceServer, err := NewFakeNetworkServiceServer()

	if err != nil {
		return nil, nil, err
	}

	vpc1.RegisterNetworkServiceServer(serv, fakeNetworkServiceServer)

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

	return vpc.NewVPC(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), serv, nil
}
