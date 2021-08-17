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
	apigateway1 "github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/apigateway/v1"
	"github.com/yandex-cloud/go-sdk/gen/apigateway"
)

func TestServerlessApiGateways(t *testing.T) {
	var serv *grpc.Server
	resource := providertest.ResourceTestData{
		Table: resources.ServerlessApiGateways(),
		Config: client.Config{
			FolderIDs: []string{"test"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			serverlessSvc, serv1, err := createApiGatewayServer()
			serv = serv1
			if err != nil {
				return nil, err
			}
			c := client.NewYandexClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"test"}, nil, nil, &client.Services{
				ApiGateway: serverlessSvc,
			}, nil)
			return c, nil
		},
		Verifiers: []providertest.Verifier{
			providertest.VerifyAtLeastOneRow("yandex_serverless_api_gateways"),
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
	serv.Stop()
}

type FakeApiGatewayServiceServer struct {
	apigateway1.UnimplementedApiGatewayServiceServer
	ApiGateway *apigateway1.ApiGateway
}

func NewFakeApiGatewayServiceServer() (*FakeApiGatewayServiceServer, error) {
	var api_gateway apigateway1.ApiGateway
	faker.SetIgnoreInterface(true)
	err := faker.FakeData(&api_gateway)
	if err != nil {
		return nil, err
	}
	return &FakeApiGatewayServiceServer{ApiGateway: &api_gateway}, nil
}

func (s *FakeApiGatewayServiceServer) List(context.Context, *apigateway1.ListApiGatewayRequest) (*apigateway1.ListApiGatewayResponse, error) {
	return &apigateway1.ListApiGatewayResponse{ApiGateways: []*apigateway1.ApiGateway{s.ApiGateway}}, nil
}

func createApiGatewayServer() (*apigateway.Apigateway, *grpc.Server, error) {
	lis, err := net.Listen("tcp", ":0")

	if err != nil {
		return nil, nil, err
	}

	serv := grpc.NewServer()
	fakeApiGatewayServiceServer, err := NewFakeApiGatewayServiceServer()

	if err != nil {
		return nil, nil, err
	}

	apigateway1.RegisterApiGatewayServiceServer(serv, fakeApiGatewayServiceServer)

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

	return apigateway.NewApigateway(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), serv, nil
}
