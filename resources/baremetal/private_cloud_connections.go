package baremetal

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/baremetal/v2"
	baremetalsdk "github.com/yandex-cloud/go-sdk/services/baremetal/v2"
)

func PrivateCloudConnections() *schema.Table {
	return &schema.Table{
		Name:        "yc_baremetal_private_cloud_connections",
		Description: `https://yandex.cloud/docs/baremetal/api-ref/grpc/PrivateCloudConnection/list#yandex.cloud.baremetal.v2.PrivateCloudConnection`,
		Multiplex:   client.FolderMultiplex(client.ServiceBaremetal),
		Resolver:    fetchPrivateCloudConnections,
		Transform:   client.TransformWithStruct(&baremetal.PrivateCloudConnection{}, client.PrimaryKeyIdTransformers("PrivateCloudConnectionId")...),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchPrivateCloudConnections(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	it := baremetalsdk.NewPrivateCloudConnectionClient(c.SDKv2).PrivateCloudConnectionsIterator(ctx, &baremetal.ListPrivateCloudConnectionsRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
