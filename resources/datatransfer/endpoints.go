package datatransfer

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/datatransfer/v1"
)

func Endpoints() *schema.Table {
	return &schema.Table{
		Name:        "yc_datatransfer_endpoints",
		Description: `https://yandex.cloud/ru/docs/data-transfer/api-ref/grpc/endpoint_service#Endpoint1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchEndpoints,
		Transform:   client.TransformWithStruct(&datatransfer.Endpoint{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchEndpoints(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.DataTransfer().Endpoint().EndpointIterator(ctx, &datatransfer.ListEndpointsRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
