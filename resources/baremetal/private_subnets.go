package baremetal

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	baremetal "github.com/yandex-cloud/go-genproto/yandex/cloud/baremetal/v1alpha"
	baremetalsdk "github.com/yandex-cloud/go-sdk/services/baremetal/v1alpha"
)

func PrivateSubnets() *schema.Table {
	return &schema.Table{
		Name:        "yc_baremetal_private_subnets",
		Description: `https://yandex.cloud/docs/baremetal/api-ref/grpc/PrivateSubnet/list#yandex.cloud.baremetal.v1alpha.PrivateSubnet`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchPrivateSubnets,
		Transform:   client.TransformWithStruct(&baremetal.PrivateSubnet{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchPrivateSubnets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	it := baremetalsdk.NewPrivateSubnetClient(c.SDKv2).Iterator(ctx, &baremetal.ListPrivateSubnetRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
