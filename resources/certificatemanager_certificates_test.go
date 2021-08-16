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
	certificatemanager1 "github.com/yandex-cloud/go-genproto/yandex/cloud/certificatemanager/v1"
	"github.com/yandex-cloud/go-sdk/gen/certificatemanager"
)

func TestCertificateManagerCertificates(t *testing.T) {
	certificatemanagerSvc, serv, err := createCertificateServer()
	if err != nil {
		t.Fatal(err)
	}
	resource := providertest.ResourceTestData{
		Table: resources.CertificateManagerCertificates(),
		Config: client.Config{
			FolderIDs: []string{"testFolder"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			c := client.NewYandexClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testFolder"}, nil, nil, &client.Services{
				CertificateManager: certificatemanagerSvc,
			}, nil)
			return c, nil
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
	serv.Stop()
}

type FakeCertificateServiceServer struct {
	certificatemanager1.UnimplementedCertificateServiceServer
	Certificate *certificatemanager1.Certificate
}

func NewFakeCertificateServiceServer() (*FakeCertificateServiceServer, error) {
	var certificate certificatemanager1.Certificate
	faker.SetIgnoreInterface(true)
	err := faker.FakeData(&certificate)
	if err != nil {
		return nil, err
	}
	return &FakeCertificateServiceServer{Certificate: &certificate}, nil
}

func (s *FakeCertificateServiceServer) List(context.Context, *certificatemanager1.ListCertificatesRequest) (*certificatemanager1.ListCertificatesResponse, error) {
	return &certificatemanager1.ListCertificatesResponse{Certificates: []*certificatemanager1.Certificate{s.Certificate}}, nil
}

func createCertificateServer() (*certificatemanager.CertificateManager, *grpc.Server, error) {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		return nil, nil, err
	}

	serv := grpc.NewServer()
	fakeCertificateServiceServer, err := NewFakeCertificateServiceServer()

	if err != nil {
		return nil, nil, err
	}

	certificatemanager1.RegisterCertificateServiceServer(serv, fakeCertificateServiceServer)

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

	return certificatemanager.NewCertificateManager(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), serv, nil
}
