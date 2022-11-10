package servers

import (
	"context"
	"net"
	"testing"

	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/yandex-cloud/cq-provider-yandex/resources/testingutils/mocks"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	organizationmanager1 "github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1"
	"github.com/yandex-cloud/go-sdk/gen/organizationmanager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func StartOrganizationManagerServer(t *testing.T, ctx context.Context) (*organizationmanager.OrganizationManager, error) {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, err
	}

	serv := grpc.NewServer()
	go func() {
		<-ctx.Done()
		serv.Stop()
	}()

	err = registerOrganizationManagerMocks(t, serv)
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

	return organizationmanager.NewOrganizationManager(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), nil
}

// go:generate mockgen -destination=../mocks/organization_manager_organization_service_server_mock.go -package=mocks github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1 OrganizationServiceServer
// go:generate mockgen -destination=../mocks/organization_manager_group_service_server_mock.go -package=mocks github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1 GroupServiceServer

func registerOrganizationManagerMocks(t *testing.T, serv *grpc.Server) error {
	ctrl := gomock.NewController(t)
	var err error
	faker.SetIgnoreInterface(true)

	var organization organizationmanager1.Organization
	err = faker.FakeData(&organization)
	if err != nil {
		return err
	}
	mOrganizationServ := mocks.NewMockOrganizationServiceServer(ctrl)
	mOrganizationServ.
		EXPECT().
		List(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *organizationmanager1.ListOrganizationsRequest) (*organizationmanager1.ListOrganizationsResponse, error) {
			if req == nil {
				return nil, status.Errorf(codes.Canceled, "request is nil")
			}
			return &organizationmanager1.ListOrganizationsResponse{Organizations: []*organizationmanager1.Organization{&organization}}, nil
		}).
		AnyTimes()

	var accessBinding access.AccessBinding
	err = faker.FakeData(&accessBinding)
	if err != nil {
		return err
	}
	mOrganizationServ.
		EXPECT().
		ListAccessBindings(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
			if req == nil {
				return nil, status.Errorf(codes.Canceled, "request is nil")
			}
			return &access.ListAccessBindingsResponse{AccessBindings: []*access.AccessBinding{&accessBinding}}, nil
		}).
		AnyTimes()
	organizationmanager1.RegisterOrganizationServiceServer(serv, mOrganizationServ)

	mGroupServ := mocks.NewMockGroupServiceServer(ctrl)
	var group organizationmanager1.Group
	faker.SetIgnoreInterface(true)
	err = faker.FakeData(&group)
	if err != nil {
		return err
	}
	mGroupServ.
		EXPECT().
		List(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *organizationmanager1.ListGroupsRequest) (*organizationmanager1.ListGroupsResponse, error) {
			if req == nil {
				return nil, status.Errorf(codes.Canceled, "request is nil")
			}
			return &organizationmanager1.ListGroupsResponse{Groups: []*organizationmanager1.Group{&group}}, nil
		}).
		AnyTimes()

	organizationmanager1.RegisterGroupServiceServer(serv, mGroupServ)

	return nil
}
