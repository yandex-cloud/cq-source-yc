package access

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
)

func CloudsAccessBindings() *schema.Table {
	return &schema.Table{
		Name:        "yc_access_bindings_resourcemanager_clouds",
		Description: `https://cloud.yandex.ru/docs/resource-manager/api-ref/grpc/cloud_service#AccessBinding`,
		Multiplex:   client.CloudMultiplex,
		Resolver:    fetchCloudsAccessBindings,
		Transform:   Transform,
		Columns: schema.ColumnList{
			client.MultiplexedResourceIdColumn,
		},
	}
}

func fetchCloudsAccessBindings(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	cloudId := c.CloudId

	it := c.SDK.ResourceManager().Cloud().CloudAccessBindingsIterator(ctx, &access.ListAccessBindingsRequest{ResourceId: cloudId})
	for it.Next() {
		value := it.Value()
		res <- value
	}

	return it.Error()
}
