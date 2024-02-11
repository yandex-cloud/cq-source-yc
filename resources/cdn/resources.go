package cdn

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/cdn/v1"
)

func Resources() *schema.Table {
	return &schema.Table{
		Name:        "yc_cdn_resources",
		Description: `https://cloud.yandex.ru/docs/cdn/api-ref/grpc/resource_service#Resource1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchResources,
		Transform:   client.TransformWithStruct(&cdn.Resource{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchResources(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.CDN().Resource().ResourceIterator(ctx, &cdn.ListResourcesRequest{FolderId: c.FolderId})
	for it.Next() {
		value := it.Value()
		res <- value
	}

	return it.Error()
}
