package compute

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
)

func Instances() *schema.Table {
	return &schema.Table{
		Name:        "yc_compute_instances",
		Description: `https://cloud.yandex.ru/docs/compute/api-ref/grpc/instance_service#Instance`,
		Resolver:    fetchInstances,
		Transform:   client.TransformWithStruct(&compute.Instance{}, client.PrimaryKeyIdTransformer),
		Multiplex:   client.FolderMultiplex,
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchInstances(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	folderId := c.FolderId

	it := c.SDK.Compute().Instance().InstanceIterator(ctx, &compute.ListInstancesRequest{FolderId: folderId})
	for it.Next() {
		value := it.Value()
		res <- value
	}

	return it.Error()
}
