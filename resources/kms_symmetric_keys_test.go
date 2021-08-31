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
	kmsSvc, serv, err := createSymmetricKeyServer()
	if err != nil {
		t.Fatal(err)
	}
	resource := providertest.ResourceTestData{
		Table: resources.KMSSymmetricKeys(),
		Config: client.Config{
			FolderIDs: []string{"test"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			c := client.NewYandexClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"test"}, nil, nil, &client.Services{
				KMS: kmsSvc,
			}, nil)
			return c, nil
		},
		Verifiers: []providertest.Verifier{
			providertest.VerifyAtLeastOneRow("yandex_kms_symmetric_keys"),
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
	lis, err := net.Listen("tcp", ":0")

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

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())

	if err != nil {
		return nil, nil, err
	}

	return kms.NewKMS(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), serv, nil
}
