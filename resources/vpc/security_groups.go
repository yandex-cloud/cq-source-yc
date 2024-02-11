package vpc

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
)

func SecurityGroups() *schema.Table {
	return &schema.Table{
		Name:        "yc_vpc_security_groups",
		Description: `https://cloud.yandex.ru/docs/vpc/api-ref/grpc/security_group_service#SecurityGroup1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchSecurityGroups,
		Transform:   client.TransformWithStruct(&vpc.SecurityGroup{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchSecurityGroups(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.VPC().SecurityGroup().SecurityGroupIterator(ctx, &vpc.ListSecurityGroupsRequest{FolderId: c.FolderId})

	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
