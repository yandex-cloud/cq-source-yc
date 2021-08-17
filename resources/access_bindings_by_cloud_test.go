package resources_test

import (
	"context"
	"fmt"
	"net"
	"testing"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	resourcemanager1 "github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
	"github.com/yandex-cloud/go-sdk/gen/resourcemanager"
	"google.golang.org/grpc"

	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/cq-provider-yandex/resources"
)

func TestAccessBindingsByCloud(t *testing.T) {
	accessBindingsByCloudServer, serv, err := createAccessBindingsByCloudServer()
	if err != nil {
		t.Fatal(err)
	}
	resource := providertest.ResourceTestData{
		Table:  resources.AccessBindingsByCloud(),
		Config: client.Config{CloudIDs: []string{"test"}},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			c := client.NewYandexClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), nil, []string{"test"}, nil, &client.Services{
				ResourceManager: accessBindingsByCloudServer,
			}, nil)
			return c, nil
		},
		Verifiers: []providertest.Verifier{
			providertest.VerifyAtLeastOneRow("yandex_access_bindings_by_cloud"),
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
	serv.Stop()
}

type FakeAccessBindingsByCloudService struct {
	resourcemanager1.UnimplementedCloudServiceServer
}

func (s FakeAccessBindingsByCloudService) ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
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

func createAccessBindingsByCloudServer() (*resourcemanager.ResourceManager, *grpc.Server, error) {
	lis, err := net.Listen("tcp", ":0")

	if err != nil {
		return nil, nil, err
	}

	serv := grpc.NewServer()
	fakeCloudServiceServer := &FakeAccessBindingsByCloudService{}

	resourcemanager1.RegisterCloudServiceServer(serv, fakeCloudServiceServer)

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

	return resourcemanager.NewResourceManager(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), serv, nil
}
