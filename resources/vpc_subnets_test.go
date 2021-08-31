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

func TestVPCSubnets(t *testing.T) {
	vpcSvc, serv, err := createSubnetServer()
	if err != nil {
		t.Fatal(err)
	}
	resource := providertest.ResourceTestData{
		Table: resources.VPCSubnets(),
		Config: client.Config{
			FolderIDs: []string{"test"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			c := client.NewYandexClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"test"}, nil, nil, &client.Services{
				VPC: vpcSvc,
			}, nil)
			return c, nil
		},
		Verifiers: []providertest.Verifier{
			providertest.VerifyAtLeastOneRow("yandex_vpc_subnets"),
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
	serv.Stop()
}

type FakeSubnetServiceServer struct {
	vpc1.UnimplementedSubnetServiceServer
	Subnet *vpc1.Subnet
}

func NewFakeSubnetServiceServer() (*FakeSubnetServiceServer, error) {
	var subnet vpc1.Subnet
	faker.SetIgnoreInterface(true)
	err := faker.FakeData(&subnet)
	if err != nil {
		return nil, err
	}
	return &FakeSubnetServiceServer{Subnet: &subnet}, nil
}

func (s *FakeSubnetServiceServer) List(context.Context, *vpc1.ListSubnetsRequest) (*vpc1.ListSubnetsResponse, error) {
	return &vpc1.ListSubnetsResponse{Subnets: []*vpc1.Subnet{s.Subnet}}, nil
}

func createSubnetServer() (*vpc.VPC, *grpc.Server, error) {
	lis, err := net.Listen("tcp", ":0")

	if err != nil {
		return nil, nil, err
	}

	serv := grpc.NewServer()
	fakeSubnetServiceServer, err := NewFakeSubnetServiceServer()

	if err != nil {
		return nil, nil, err
	}

	vpc1.RegisterSubnetServiceServer(serv, fakeSubnetServiceServer)

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

	return vpc.NewVPC(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), serv, nil
}
