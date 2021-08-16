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

func TestVPCSecurityGroups(t *testing.T) {
	vpcSvc, serv, err := createSecurityGroupServer()
	if err != nil {
		t.Fatal(err)
	}
	resource := providertest.ResourceTestData{
		Table: resources.VPCSecurityGroups(),
		Config: client.Config{
			FolderIDs: []string{"testFolder"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			c := client.NewYandexClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testFolder"}, nil, &client.Services{
				VPC: vpcSvc,
			}, nil)
			return c, nil
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
	serv.Stop()
}

type FakeSecurityGroupServiceServer struct {
	vpc1.UnimplementedSecurityGroupServiceServer
	SecurityGroup *vpc1.SecurityGroup
}

func NewFakeSecurityGroupServiceServer() (*FakeSecurityGroupServiceServer, error) {
	var security_group vpc1.SecurityGroup
	faker.SetIgnoreInterface(true)
	err := faker.FakeData(&security_group)
	if err != nil {
		return nil, err
	}
	return &FakeSecurityGroupServiceServer{SecurityGroup: &security_group}, nil
}

func (s *FakeSecurityGroupServiceServer) List(context.Context, *vpc1.ListSecurityGroupsRequest) (*vpc1.ListSecurityGroupsResponse, error) {
	return &vpc1.ListSecurityGroupsResponse{SecurityGroups: []*vpc1.SecurityGroup{s.SecurityGroup}}, nil
}

func createSecurityGroupServer() (*vpc.VPC, *grpc.Server, error) {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		return nil, nil, err
	}

	serv := grpc.NewServer()
	fakeSecurityGroupServiceServer, err := NewFakeSecurityGroupServiceServer()

	if err != nil {
		return nil, nil, err
	}

	vpc1.RegisterSecurityGroupServiceServer(serv, fakeSecurityGroupServiceServer)

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
