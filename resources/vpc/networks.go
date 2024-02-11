package vpc

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
)

func Networks() *schema.Table {
	return &schema.Table{
		Name:        "yc_vpc_networks",
		Description: `https://cloud.yandex.ru/docs/vpc/api-ref/grpc/network_service#Network1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchNetworks,
		Transform:   client.TransformWithStruct(&vpc.Network{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchNetworks(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.VPC().Network().NetworkIterator(ctx, &vpc.ListNetworksRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
