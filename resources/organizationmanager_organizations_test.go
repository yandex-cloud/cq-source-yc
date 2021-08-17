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
	organizationmanager1 "github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1"
	"github.com/yandex-cloud/go-sdk/gen/organizationmanager"
)

func TestOrganizationManagerOrganizations(t *testing.T) {
	organizationmanagerSvc, serv, err := createOrganizationServer()
	if err != nil {
		t.Fatal(err)
	}
	resource := providertest.ResourceTestData{
		Table: resources.OrganizationManagerOrganizations(),
		Config: client.Config{
			FolderIDs: []string{"testFolder"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			c := client.NewYandexClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testFolder"}, nil, nil, &client.Services{
				OrganizationManager: organizationmanagerSvc,
			}, nil)
			return c, nil
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
	serv.Stop()
}

type FakeOrganizationServiceServer struct {
	organizationmanager1.UnimplementedOrganizationServiceServer
	Organization *organizationmanager1.Organization
}

func NewFakeOrganizationServiceServer() (*FakeOrganizationServiceServer, error) {
	var organization organizationmanager1.Organization
	faker.SetIgnoreInterface(true)
	err := faker.FakeData(&organization)
	if err != nil {
		return nil, err
	}
	return &FakeOrganizationServiceServer{Organization: &organization}, nil
}

func (s *FakeOrganizationServiceServer) List(context.Context, *organizationmanager1.ListOrganizationsRequest) (*organizationmanager1.ListOrganizationsResponse, error) {
	return &organizationmanager1.ListOrganizationsResponse{Organizations: []*organizationmanager1.Organization{s.Organization}}, nil
}

func createOrganizationServer() (*organizationmanager.OrganizationManager, *grpc.Server, error) {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		return nil, nil, err
	}

	serv := grpc.NewServer()
	fakeOrganizationServiceServer, err := NewFakeOrganizationServiceServer()

	if err != nil {
		return nil, nil, err
	}

	organizationmanager1.RegisterOrganizationServiceServer(serv, fakeOrganizationServiceServer)

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

	return organizationmanager.NewOrganizationManager(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), serv, nil
}
