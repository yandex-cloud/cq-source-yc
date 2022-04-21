package servers

import (
	"context"
	"net"
	"testing"

	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/yandex-cloud/cq-provider-yandex/resources/testingutils/mocks"
	organizationmanager1 "github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1/saml"
	organizationmanager "github.com/yandex-cloud/go-sdk/gen/organizationmanager/saml"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func StartOrganizationManagerSAMLServer(t *testing.T, ctx context.Context) (*organizationmanager.OrganizationManagerSAML, error) {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, err
	}

	serv := grpc.NewServer()
	go func() {
		<-ctx.Done()
		serv.Stop()
	}()

	err = registerOrganizationManagerSAMLMocks(t, serv)
	if err != nil {
		return nil, err
	}

	go func() {
		_ = serv.Serve(lis)
	}()

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return organizationmanager.NewOrganizationManagerSAML(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), nil
}

//go:generate mockgen -destination=../mocks/organization_manager_saml_federation_service_server_mock.go -package=mocks github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1/saml FederationServiceServer

func registerOrganizationManagerSAMLMocks(t *testing.T, serv *grpc.Server) error {
	ctrl := gomock.NewController(t)
	var err error
	faker.SetIgnoreInterface(true)

	var federation organizationmanager1.Federation
	err = faker.FakeData(&federation)
	if err != nil {
		return err
	}
	mFederationServ := mocks.NewMockFederationServiceServer(ctrl)
	mFederationServ.
		EXPECT().
		List(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *organizationmanager1.ListFederationsRequest) (*organizationmanager1.ListFederationsResponse, error) {
			if req == nil {
				return nil, status.Errorf(codes.Canceled, "request is nil")
			}
			return &organizationmanager1.ListFederationsResponse{Federations: []*organizationmanager1.Federation{&federation}}, nil
		}).
		AnyTimes()
	organizationmanager1.RegisterFederationServiceServer(serv, mFederationServ)

	return nil
}
