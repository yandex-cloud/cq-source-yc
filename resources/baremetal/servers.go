package baremetal

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	baremetal "github.com/yandex-cloud/go-genproto/yandex/cloud/baremetal/v1alpha"
	baremetalsdk "github.com/yandex-cloud/go-sdk/services/baremetal/v1alpha"
)

func Servers() *schema.Table {
	return &schema.Table{
		Name:        "yc_baremetal_servers",
		Description: `https://yandex.cloud/docs/baremetal/api-ref/grpc/Server/list#yandex.cloud.baremetal.v1alpha.Server`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchServers,
		Transform:   client.TransformWithStruct(&baremetal.Server{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchServers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	it := baremetalsdk.NewServerClient(c.SDKv2).Iterator(ctx, &baremetal.ListServerRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
