package access

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
)

func FoldersAccessBindings() *schema.Table {
	return &schema.Table{
		Name:        "yc_access_bindings_resourcemanager_folders",
		Description: `https://cloud.yandex.ru/docs/resource-manager/api-ref/grpc/folder_service#AccessBinding`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchFoldersAccessBindings,
		Transform:   Transform,
		Columns: schema.ColumnList{
			client.MultiplexedResourceIdColumn,
		},
	}
}

func fetchFoldersAccessBindings(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	folderId := c.FolderId

	it := c.SDK.ResourceManager().Folder().FolderAccessBindingsIterator(ctx, &access.ListAccessBindingsRequest{ResourceId: folderId})
	for it.Next() {
		value := it.Value()
		res <- value
	}

	return it.Error()
}
