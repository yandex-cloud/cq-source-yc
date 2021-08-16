package resources_test

import (
	"context"
	"errors"
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
	resourcemanager1 "github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
	"github.com/yandex-cloud/go-sdk/gen/resourcemanager"
)

func TestResourceManagerFolders(t *testing.T) {
	var serv *grpc.Server
	resource := providertest.ResourceTestData{
		Table: resources.ResourceManagerFolders(),
		Config: client.Config{
			FolderIDs: []string{"testFolder"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			resourcemanagerSvc, serv1, err := createFolderServer()
			serv = serv1
			if err != nil {
				return nil, err
			}
			c := client.NewYandexClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testFolder"}, nil, &client.Services{
				ResourceManager: resourcemanagerSvc,
			}, nil)
			return c, nil
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
	serv.Stop()
}

type FakeFolderServiceServer struct {
	resourcemanager1.UnimplementedFolderServiceServer
	Folder *resourcemanager1.Folder
}

func NewFakeFolderServiceServer() (*FakeFolderServiceServer, error) {
	var Folder resourcemanager1.Folder
	faker.SetIgnoreInterface(true)
	err := faker.FakeData(&Folder)
	if err != nil {
		return nil, err
	}
	Folder.Name = "testFolder"
	return &FakeFolderServiceServer{Folder: &Folder}, nil
}

func (s *FakeFolderServiceServer) Get(_ context.Context, req *resourcemanager1.GetFolderRequest) (*resourcemanager1.Folder, error) {
	if req.FolderId == "testFolder" {
		return s.Folder, nil
	}
	return nil, errors.New("no such Folder")
}

func createFolderServer() (*resourcemanager.ResourceManager, *grpc.Server, error) {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		return nil, nil, err
	}

	serv := grpc.NewServer()
	fakeFolderServiceServer, err := NewFakeFolderServiceServer()

	if err != nil {
		return nil, nil, err
	}

	resourcemanager1.RegisterFolderServiceServer(serv, fakeFolderServiceServer)

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

	return resourcemanager.NewResourceManager(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), serv, nil
}