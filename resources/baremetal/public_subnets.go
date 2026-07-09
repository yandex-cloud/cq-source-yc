package baremetal

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/baremetal/v2"
	baremetalsdk "github.com/yandex-cloud/go-sdk/services/baremetal/v2"
)

func PublicSubnets() *schema.Table {
	return &schema.Table{
		Name:        "yc_baremetal_public_subnets",
		Description: `https://yandex.cloud/docs/baremetal/api-ref/grpc/PublicSubnet/list#yandex.cloud.baremetal.v2.PublicSubnet`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchPublicSubnets,
		Transform:   client.TransformWithStruct(&baremetal.PublicSubnet{}, client.PrimaryKeyIdTransformers("PublicSubnetId")...),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchPublicSubnets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	it := baremetalsdk.NewPublicSubnetClient(c.SDKv2).PublicSubnetsIterator(ctx, &baremetal.ListPublicSubnetsRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
