package resources_test

import (
	"context"
	"fmt"
	"net"
	"testing"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	organizationmanager1 "github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1"
	"github.com/yandex-cloud/go-sdk/gen/organizationmanager"
	"google.golang.org/grpc"

	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/cq-provider-yandex/resources"
)

func TestAccessBindingsByOrganization(t *testing.T) {
	accessBindingsByOrganizationServer, serv, err := createAccessBindingsByOrganizationServer()
	if err != nil {
		t.Fatal(err)
	}
	resource := providertest.ResourceTestData{
		Table:  resources.AccessBindingsByOrganization(),
		Config: client.Config{OrganizationIDs: []string{"test"}},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			c := client.NewYandexClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), nil, nil, []string{"test"}, &client.Services{
				OrganizationManager: accessBindingsByOrganizationServer,
			}, nil)
			return c, nil
		},
		Verifiers: []providertest.Verifier{
			providertest.VerifyAtLeastOneRow("yandex_access_bindings_by_organization"),
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
	serv.Stop()
}

type FakeAccessBindingsByOrganizationService struct {
	organizationmanager1.UnimplementedOrganizationServiceServer
}

func (s FakeAccessBindingsByOrganizationService) ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	return &access.ListAccessBindingsResponse{AccessBindings: []*access.AccessBinding{
		{
			RoleId:  "awesome_role",
			Subject: &access.Subject{Id: "1", Type: "userAccount"},
		},
		{
			RoleId:  "another_awesome_role",
			Subject: &access.Subject{Id: "2", Type: "federationAccount"},
		},
		{
			RoleId:  "another_role_but_not_awesome",
			Subject: &access.Subject{Id: "3", Type: "serviceAccount"},
		},
	}}, nil
}

func createAccessBindingsByOrganizationServer() (*organizationmanager.OrganizationManager, *grpc.Server, error) {
	lis, err := net.Listen("tcp", ":0")

	if err != nil {
		return nil, nil, err
	}

	serv := grpc.NewServer()
	fakeOrganizationServiceServer := &FakeAccessBindingsByOrganizationService{}

	organizationmanager1.RegisterOrganizationServiceServer(serv, fakeOrganizationServiceServer)

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

	return organizationmanager.NewOrganizationManager(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), serv, nil
}
