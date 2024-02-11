package compute

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
)

func Disks() *schema.Table {
	return &schema.Table{
		Name:        "yc_compute_disks",
		Description: `https://cloud.yandex.ru/docs/compute/api-ref/grpc/disk_service#Disk1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchDisks,
		Transform:   client.TransformWithStruct(&compute.Disk{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
		Relations: schema.Tables{DisksImages()},
	}
}

func fetchDisks(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.Compute().Disk().DiskIterator(ctx, &compute.ListDisksRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
