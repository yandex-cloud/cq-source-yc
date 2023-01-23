package servers

import (
	"context"
	"net"
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/yandex-cloud/cq-provider-yandex/resources/testingutils/mocks"
	vpc1 "github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
	"github.com/yandex-cloud/go-sdk/gen/vpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func StartVpcServer(t *testing.T, ctx context.Context) (*vpc.VPC, error) {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, err
	}

	serv := grpc.NewServer()
	go func() {
		<-ctx.Done()
		serv.Stop()
	}()

	err = registerVpcMocks(t, serv)
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

	return vpc.NewVPC(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), nil
}

//go:generate mockgen -destination=../mocks/vpc_address_service_server_mock.go -package=mocks github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1 AddressServiceServer
//go:generate mockgen -destination=../mocks/vpc_network_service_server_mock.go -package=mocks github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1 NetworkServiceServer
//go:generate mockgen -destination=../mocks/vpc_security_group_service_server_mock.go -package=mocks github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1 SecurityGroupServiceServer
//go:generate mockgen -destination=../mocks/vpc_subnet_service_server_mock.go -package=mocks github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1 SubnetServiceServer

func registerVpcMocks(t *testing.T, serv *grpc.Server) error {
	ctrl := gomock.NewController(t)
	var err error

	var address vpc1.Address
	err = faker.FakeObject(&address)
	if err != nil {
		return err
	}
	addressServ := mocks.NewMockAddressServiceServer(ctrl)
	addressServ.
		EXPECT().
		List(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *vpc1.ListAddressesRequest) (*vpc1.ListAddressesResponse, error) {
			if req == nil {
				return nil, status.Errorf(codes.Canceled, "request is nil")
			}
			if req.FolderId != "test-folder-id" {
				return nil, status.Errorf(codes.NotFound, "folder not found")
			}
			return &vpc1.ListAddressesResponse{Addresses: []*vpc1.Address{&address}}, nil
		}).
		AnyTimes()
	vpc1.RegisterAddressServiceServer(serv, addressServ)

	var network vpc1.Network
	err = faker.FakeObject(&network)
	if err != nil {
		return err
	}
	networkServ := mocks.NewMockNetworkServiceServer(ctrl)
	networkServ.
		EXPECT().
		List(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *vpc1.ListNetworksRequest) (*vpc1.ListNetworksResponse, error) {
			if req == nil {
				return nil, status.Errorf(codes.Canceled, "request is nil")
			}
			if req.FolderId != "test-folder-id" {
				return nil, status.Errorf(codes.NotFound, "folder not found")
			}
			return &vpc1.ListNetworksResponse{Networks: []*vpc1.Network{&network}}, nil
		}).
		AnyTimes()
	vpc1.RegisterNetworkServiceServer(serv, networkServ)

	var securityGroup vpc1.SecurityGroup
	err = faker.FakeObject(&securityGroup)
	if err != nil {
		return err
	}
	securityGroupServ := mocks.NewMockSecurityGroupServiceServer(ctrl)
	securityGroupServ.
		EXPECT().
		List(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *vpc1.ListSecurityGroupsRequest) (*vpc1.ListSecurityGroupsResponse, error) {
			if req == nil {
				return nil, status.Errorf(codes.Canceled, "request is nil")
			}
			if req.FolderId != "test-folder-id" {
				return nil, status.Errorf(codes.NotFound, "folder not found")
			}
			return &vpc1.ListSecurityGroupsResponse{SecurityGroups: []*vpc1.SecurityGroup{&securityGroup}}, nil
		}).
		AnyTimes()
	vpc1.RegisterSecurityGroupServiceServer(serv, securityGroupServ)

	var subnet vpc1.Subnet
	err = faker.FakeObject(&subnet)
	if err != nil {
		return err
	}
	subnetServ := mocks.NewMockSubnetServiceServer(ctrl)
	subnetServ.
		EXPECT().
		List(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *vpc1.ListSubnetsRequest) (*vpc1.ListSubnetsResponse, error) {
			if req == nil {
				return nil, status.Errorf(codes.Canceled, "request is nil")
			}
			if req.FolderId != "test-folder-id" {
				return nil, status.Errorf(codes.NotFound, "folder not found")
			}
			return &vpc1.ListSubnetsResponse{Subnets: []*vpc1.Subnet{&subnet}}, nil
		}).
		AnyTimes()
	vpc1.RegisterSubnetServiceServer(serv, subnetServ)

	return nil
}
