package compute

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
)

func Images() *schema.Table {
	return &schema.Table{
		Name:        "yc_compute_images",
		Description: `https://cloud.yandex.ru/docs/compute/api-ref/grpc/image_service#Image2`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchImages,
		Transform:   client.TransformWithStruct(&compute.Image{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.Compute().Image().ImageIterator(ctx, &compute.ListImagesRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
