package servers

import (
	"context"
	"net"
	"testing"

	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/yandex-cloud/cq-provider-yandex/resources/testingutils/mocks"
	apigateway1 "github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/apigateway/v1"
	"github.com/yandex-cloud/go-sdk/gen/apigateway"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func StartApiGatewayServer(t *testing.T, ctx context.Context) (*apigateway.Apigateway, error) {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, err
	}

	serv := grpc.NewServer()
	go func() {
		<-ctx.Done()
		serv.Stop()
	}()

	err = registerApiGatewayMocks(t, serv)
	if err != nil {
		return nil, err
	}

	go func() {
		_ = serv.Serve(lis)
	}()

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return apigateway.NewApigateway(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), nil
}

//go:generate mockgen -destination=../mocks/api_gateway_service_server_mock.go -package=mocks github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/apigateway/v1 ApiGatewayServiceServer

func registerApiGatewayMocks(t *testing.T, serv *grpc.Server) error {
	ctrl := gomock.NewController(t)
	var err error
	faker.SetIgnoreInterface(true)

	var apigateway apigateway1.ApiGateway
	err = faker.FakeData(&apigateway)
	if err != nil {
		return err
	}
	mApiGateway := mocks.NewMockApiGatewayServiceServer(ctrl)
	mApiGateway.
		EXPECT().
		List(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *apigateway1.ListApiGatewayRequest) (*apigateway1.ListApiGatewayResponse, error) {
			if req == nil {
				return nil, status.Errorf(codes.Canceled, "request is nil")
			}
			if req.FolderId != "test-folder-id" {
				return nil, status.Errorf(codes.NotFound, "folder not found")
			}
			return &apigateway1.ListApiGatewayResponse{ApiGateways: []*apigateway1.ApiGateway{&apigateway}}, nil
		}).
		AnyTimes()
	apigateway1.RegisterApiGatewayServiceServer(serv, mApiGateway)

	return nil
}
