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
	iam1 "github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
	"github.com/yandex-cloud/go-sdk/gen/iam"
)

func TestIAMServiceAccounts(t *testing.T) {
	var serv *grpc.Server
	resource := providertest.ResourceTestData{
		Table: resources.IAMServiceAccounts(),
		Config: client.Config{
			FolderIDs: []string{"testFolder"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			iamSvc, serv1, err := createServiceAccountServer()
			serv = serv1
			if err != nil {
				return nil, err
			}
			c := client.NewYandexClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testFolder"}, nil, &client.Services{
				IAM: iamSvc,
			}, nil)
			return c, nil
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
	serv.Stop()
}

type FakeServiceAccountServiceServer struct {
	iam1.UnimplementedServiceAccountServiceServer
	ServiceAccount *iam1.ServiceAccount
}

func NewFakeServiceAccountServiceServer() (*FakeServiceAccountServiceServer, error) {
	var service_account iam1.ServiceAccount
	faker.SetIgnoreInterface(true)
	err := faker.FakeData(&service_account)
	if err != nil {
		return nil, err
	}
	return &FakeServiceAccountServiceServer{ServiceAccount: &service_account}, nil
}

func (s *FakeServiceAccountServiceServer) List(context.Context, *iam1.ListServiceAccountsRequest) (*iam1.ListServiceAccountsResponse, error) {
	return &iam1.ListServiceAccountsResponse{ServiceAccounts: []*iam1.ServiceAccount{s.ServiceAccount}}, nil
}

func createServiceAccountServer() (*iam.IAM, *grpc.Server, error) {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		return nil, nil, err
	}

	serv := grpc.NewServer()
	fakeServiceAccountServiceServer, err := NewFakeServiceAccountServiceServer()

	if err != nil {
		return nil, nil, err
	}

	iam1.RegisterServiceAccountServiceServer(serv, fakeServiceAccountServiceServer)

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

	return iam.NewIAM(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), serv, nil
}
