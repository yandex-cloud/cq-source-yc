package resourcemanager

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
)

func fetchFolders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	folder, err := c.Services.ResourceManager.Folder().Get(ctx, &resourcemanager.GetFolderRequest{FolderId: c.MultiplexedResourceId})
	if err != nil {
		return err
	}

	res <- folder

	return nil
}
