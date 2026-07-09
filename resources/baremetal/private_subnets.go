package baremetal

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/baremetal/v2"
	baremetalsdk "github.com/yandex-cloud/go-sdk/services/baremetal/v2"
)

func PrivateSubnets() *schema.Table {
	return &schema.Table{
		Name:        "yc_baremetal_private_subnets",
		Description: `https://yandex.cloud/docs/baremetal/api-ref/grpc/PrivateSubnet/list#yandex.cloud.baremetal.v2.PrivateSubnet`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchPrivateSubnets,
		Transform:   client.TransformWithStruct(&baremetal.PrivateSubnet{}, client.PrimaryKeyIdTransformers("PrivateSubnetId")...),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchPrivateSubnets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	it := baremetalsdk.NewPrivateSubnetClient(c.SDKv2).PrivateSubnetsIterator(ctx, &baremetal.ListPrivateSubnetsRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
