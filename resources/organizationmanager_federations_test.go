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
	saml1 "github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1/saml"
	"github.com/yandex-cloud/go-sdk/gen/organizationmanager/saml"
)

func TestOrganizationManagerFederations(t *testing.T) {
	samlSvc, serv, err := createFederationsServer()
	if err != nil {
		t.Fatal(err)
	}
	resource := providertest.ResourceTestData{
		Table: resources.OrganizationManagerFederations(),
		Config: client.Config{
			CloudIDs: []string{"test"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			c := client.NewYandexClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), nil, nil, []string{"test"}, &client.Services{
				OrganizationManagerSAML: samlSvc,
			}, nil)
			return c, nil
		},
		Verifiers: []providertest.Verifier{
			providertest.VerifyAtLeastOneRow("yandex_organizationmanager_federations"),
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
	serv.Stop()
}

type FakeFederationsServiceServer struct {
	saml1.UnimplementedFederationServiceServer
	Federation *saml1.Federation
}

func NewFakeFederationsServiceServer() (*FakeFederationsServiceServer, error) {
	var federation saml1.Federation
	faker.SetIgnoreInterface(true)
	err := faker.FakeData(&federation)
	if err != nil {
		return nil, err
	}
	return &FakeFederationsServiceServer{Federation: &federation}, nil
}

func (s *FakeFederationsServiceServer) List(context.Context, *saml1.ListFederationsRequest) (*saml1.ListFederationsResponse, error) {
	return &saml1.ListFederationsResponse{Federations: []*saml1.Federation{s.Federation}}, nil
}

func createFederationsServer() (*saml.OrganizationManagerSAML, *grpc.Server, error) {
	lis, err := net.Listen("tcp", ":0")

	if err != nil {
		return nil, nil, err
	}

	serv := grpc.NewServer()
	fakeFederationsServiceServer, err := NewFakeFederationsServiceServer()

	if err != nil {
		return nil, nil, err
	}

	saml1.RegisterFederationServiceServer(serv, fakeFederationsServiceServer)

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

	return saml.NewOrganizationManagerSAML(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), serv, nil
}
