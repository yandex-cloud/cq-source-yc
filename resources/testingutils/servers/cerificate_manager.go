package servers

import (
	"context"
	"net"
	"testing"

	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/yandex-cloud/cq-provider-yandex/resources/testingutils/mocks"
	certificatemanager1 "github.com/yandex-cloud/go-genproto/yandex/cloud/certificatemanager/v1"
	"github.com/yandex-cloud/go-sdk/gen/certificatemanager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func StartCertificateManagerServer(t *testing.T, ctx context.Context) (*certificatemanager.CertificateManager, error) {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, err
	}

	serv := grpc.NewServer()
	go func() {
		select {
		case <-ctx.Done():
			serv.Stop()
		}
	}()

	err = registerCertificateManagerMocks(t, serv)
	if err != nil {
		return nil, err
	}

	go serv.Serve(lis)

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return certificatemanager.NewCertificateManager(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), nil
}

//go:generate mockgen -destination=../mocks/certificate_manager_certificate_service_server_mock.go -package=mocks github.com/yandex-cloud/go-genproto/yandex/cloud/certificatemanager/v1 CertificateServiceServer

func registerCertificateManagerMocks(t *testing.T, serv *grpc.Server) error {
	ctrl := gomock.NewController(t)
	var err error
	faker.SetIgnoreInterface(true)

	var certificate certificatemanager1.Certificate
	err = faker.FakeData(&certificate)
	if err != nil {
		return err
	}
	mCertificateServ := mocks.NewMockCertificateServiceServer(ctrl)
	mCertificateServ.
		EXPECT().
		List(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, req *certificatemanager1.ListCertificatesRequest) (*certificatemanager1.ListCertificatesResponse, error) {
			if req == nil {
				return nil, status.Errorf(codes.Canceled, "request is nil")
			}
			if req.FolderId != "test-folder-id" {
				return nil, status.Errorf(codes.NotFound, "folder not found")
			}
			return &certificatemanager1.ListCertificatesResponse{Certificates: []*certificatemanager1.Certificate{&certificate}}, nil
		}).
		AnyTimes()
	certificatemanager1.RegisterCertificateServiceServer(serv, mCertificateServ)

	return nil
}
