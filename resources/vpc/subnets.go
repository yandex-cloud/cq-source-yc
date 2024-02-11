package vpc

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
)

func Subnets() *schema.Table {
	return &schema.Table{
		Name:        "yc_vpc_subnets",
		Description: `https://cloud.yandex.ru/docs/vpc/api-ref/grpc/subnet_service#Subnet1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchSubnets,
		Transform:   client.TransformWithStruct(&vpc.Subnet{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchSubnets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.VPC().Subnet().SubnetIterator(ctx, &vpc.ListSubnetsRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
