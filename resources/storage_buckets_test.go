package resources_test

//
//import (
//	"context"
//	"fmt"
//	"net"
//	"testing"
//
//	"github.com/aws/aws-sdk-go/service/s3"
//	"google.golang.org/grpc"
//
//	"github.com/cloudquery/cq-provider-sdk/logging"
//	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
//	"github.com/cloudquery/cq-provider-sdk/provider/schema"
//	"github.com/cloudquery/faker/v3"
//	"github.com/hashicorp/go-hclog"
//	"github.com/yandex-cloud/cq-provider-yandex/client"
//	"github.com/yandex-cloud/cq-provider-yandex/resources"
//	storage1 "github.com/yandex-cloud/go-genproto/yandex/cloud/storage/v1"
//	"github.com/yandex-cloud/go-sdk/gen/storage"
//)
//
//func TestStorageBuckets(t *testing.T) {
//	storageSvc, serv, err := createBucketsServer()
//	if err != nil {
//		t.Fatal(err)
//	}
//	resource := providertest.ResourceTestData{
//		Table: resources.StorageBuckets(),
//		Config: client.Config{
//			FolderIDs: []string{"test"},
//		},
//		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
//			c := client.NewYandexClient(logging.New(&hclog.LoggerOptions{
//				Level: hclog.Warn,
//			}), []string{"test"}, nil, nil, nil, nil)
//			return c, nil
//		},
//	}
//	providertest.TestResource(t, resources.Provider, resource)
//	serv.Stop()
//}
//
//type FakeBucketsServiceServer struct {
//	storage1.UnimplementedBucketsServiceServer
//	Buckets *storage1.Buckets
//}
//
//func NewFakeBucketsServiceServer() (*FakeBucketsServiceServer, error) {
//	var buckets storage1.Buckets
//	faker.SetIgnoreInterface(true)
//	err := faker.FakeData(&buckets)
//	if err != nil {
//		return nil, err
//	}
//	return &FakeBucketsServiceServer{Buckets: &buckets}, nil
//}
//
//func (s *FakeBucketsServiceServer) List(context.Context, *storage1.ListBucketsRequest) (*storage1.ListBucketsResponse, error) {
//	return &storage1.ListBucketsResponse{Buckets: []*storage1.Buckets{s.Buckets}}, nil
//}
//
//func createBucketsServer() (*storage.Storage, *grpc.Server, error) {
//	s3.New()
//
//
//	lis, err := net.Listen("tcp", ":0")
//
//	if err != nil {
//		return nil, nil, err
//	}
//
//	serv := grpc.NewServer()
//	fakeBucketsServiceServer, err := NewFakeBucketsServiceServer()
//
//	if err != nil {
//		return nil, nil, err
//	}
//
//	storage1.RegisterBucketsServiceServer(serv, fakeBucketsServiceServer)
//
//	go func() {
//		err := serv.Serve(lis)
//		if err != nil {
//			fmt.Println(err.Error())
//		}
//	}()
//
//	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
//
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return storage.NewStorage(
//		func(ctx context.Context) (*grpc.ClientConn, error) {
//			return conn, nil
//		},
//	), serv, nil
//}
