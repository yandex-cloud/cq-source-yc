package vpc

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
)

func Gateways() *schema.Table {
	return &schema.Table{
		Name:        "yc_vpc_gateways",
		Description: `https://cloud.yandex.ru/docs/vpc/api-ref/grpc/gateway_service#Gateway1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchGateways,
		Transform:   client.TransformWithStruct(&vpc.Gateway{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchGateways(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.VPC().Gateway().GatewayIterator(ctx, &vpc.ListGatewaysRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
