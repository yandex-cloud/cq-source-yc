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
	kms1 "github.com/yandex-cloud/go-genproto/yandex/cloud/kms/v1"
	"github.com/yandex-cloud/go-sdk/gen/kms"
)

func TestKMSSymmetricKeys(t *testing.T) {
	var serv *grpc.Server
	resource := providertest.ResourceTestData{
		Table: resources.KMSSymmetricKeys(),
		Config: client.Config{
			FolderIDs: []string{"testFolder"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			kmsSvc, serv1, err := createSymmetricKeyServer()
			serv = serv1
			if err != nil {
				return nil, err
			}
			c := client.NewYandexClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testFolder"}, nil, &client.Services{
				KMS: kmsSvc,
			}, nil)
			return c, nil
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
	serv.Stop()
}

type FakeSymmetricKeyServiceServer struct {
	kms1.UnimplementedSymmetricKeyServiceServer
	SymmetricKey *kms1.SymmetricKey
}

func NewFakeSymmetricKeyServiceServer() (*FakeSymmetricKeyServiceServer, error) {
	var symmetric_key kms1.SymmetricKey
	faker.SetIgnoreInterface(true)
	err := faker.FakeData(&symmetric_key)
	if err != nil {
		return nil, err
	}
	return &FakeSymmetricKeyServiceServer{SymmetricKey: &symmetric_key}, nil
}

func (s *FakeSymmetricKeyServiceServer) List(context.Context, *kms1.ListSymmetricKeysRequest) (*kms1.ListSymmetricKeysResponse, error) {
	return &kms1.ListSymmetricKeysResponse{Keys: []*kms1.SymmetricKey{s.SymmetricKey}}, nil
}

func createSymmetricKeyServer() (*kms.KMS, *grpc.Server, error) {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		return nil, nil, err
	}

	serv := grpc.NewServer()
	fakeSymmetricKeyServiceServer, err := NewFakeSymmetricKeyServiceServer()

	if err != nil {
		return nil, nil, err
	}

	kms1.RegisterSymmetricKeyServiceServer(serv, fakeSymmetricKeyServiceServer)

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

	return kms.NewKMS(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), serv, nil
}
