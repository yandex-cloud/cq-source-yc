package vpc

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
)

func RouteTables() *schema.Table {
	return &schema.Table{
		Name:        "yc_vpc_route_tables",
		Description: `https://cloud.yandex.ru/docs/vpc/api-ref/grpc/route_table_service#RouteTable1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchRouteTables,
		Transform:   client.TransformWithStruct(&vpc.RouteTable{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchRouteTables(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.VPC().RouteTable().RouteTableIterator(ctx, &vpc.ListRouteTablesRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
