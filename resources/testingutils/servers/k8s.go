package servers

import (
	"context"
	"net"
	"testing"

	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/yandex-cloud/cq-provider-yandex/resources/testingutils/mocks"
	k8s1 "github.com/yandex-cloud/go-genproto/yandex/cloud/k8s/v1"
	k8s "github.com/yandex-cloud/go-sdk/gen/kubernetes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func StartK8SServer(t *testing.T, ctx context.Context) (*k8s.Kubernetes, error) {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, err
	}

	serv := grpc.NewServer()
	go func() {
		<-ctx.Done()
		serv.Stop()
	}()

	err = registerK8SMocks(t, serv)
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

	return k8s.NewKubernetes(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), nil
}

//go:generate mockgen -destination=../mocks/k8s_cluster_service_server_mock.go -package=mocks github.com/yandex-cloud/go-genproto/yandex/cloud/k8s/v1 ClusterServiceServer

func registerK8SMocks(t *testing.T, serv *grpc.Server) error {
	ctrl := gomock.NewController(t)
	var err error
	faker.SetIgnoreInterface(true)

	var cluster k8s1.Cluster
	err = faker.FakeData(&cluster)
	if err != nil {
		return err
	}
	mClusterServ := mocks.NewMockClusterServiceServer(ctrl)
	mClusterServ.
		EXPECT().
		List(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *k8s1.ListClustersRequest) (*k8s1.ListClustersResponse, error) {
			if req == nil {
				return nil, status.Errorf(codes.Canceled, "request is nil")
			}
			if req.FolderId != "test-folder-id" {
				return nil, status.Errorf(codes.NotFound, "folder not found")
			}
			return &k8s1.ListClustersResponse{Clusters: []*k8s1.Cluster{&cluster}}, nil
		}).
		AnyTimes()
	k8s1.RegisterClusterServiceServer(serv, mClusterServ)

	return nil
}
