package baremetal

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	baremetal "github.com/yandex-cloud/go-genproto/yandex/cloud/baremetal/v1alpha"
	baremetalsdk "github.com/yandex-cloud/go-sdk/services/baremetal/v1alpha"
)

func Vrfs() *schema.Table {
	return &schema.Table{
		Name:        "yc_baremetal_vrfs",
		Description: `https://yandex.cloud/docs/baremetal/api-ref/grpc/Vrf/list#yandex.cloud.baremetal.v1alpha.Vrf`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchVrfs,
		Transform:   client.TransformWithStruct(&baremetal.Vrf{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchVrfs(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	it := baremetalsdk.NewVrfClient(c.SDKv2).Iterator(ctx, &baremetal.ListVrfRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
