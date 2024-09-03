package resourcemanager

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
)

func Folders() *schema.Table {
	return &schema.Table{
		Name:        "yc_resourcemanager_folders",
		Description: `https://cloud.yandex.ru/docs/resource-manager/api-ref/grpc/folder_service#Folder1`,
		Resolver:    fetchFolders,
		Multiplex:   client.FolderMultiplex,
		Transform:   client.TransformWithStruct(&resourcemanager.Folder{}, client.PrimaryKeyIdTransformer),
	}
}

func fetchFolders(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	folder, err := c.SDK.ResourceManager().Folder().Get(ctx, &resourcemanager.GetFolderRequest{FolderId: c.FolderId})
	if err != nil {
		return err
	}
	res <- folder

	return nil
}
