package compute

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
)

func DisksImages() *schema.Table {
	return &schema.Table{
		Name: "yc_compute_disks_images",
		Description: `This table is exact copy of [yc_compute_images](yc_compute_images.md), but contains images used in [yc_compute_disks](yc_compute_disks.md)
https://cloud.yandex.ru/docs/compute/api-ref/grpc/image_service#Image2`,
		Resolver:  fetchDisksImages,
		Transform: client.TransformWithStruct(&compute.Image{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchDisksImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	disk, ok := parent.Item.(*compute.Disk)
	if !ok {
		return fmt.Errorf("parent is not type of *compute.Disk")
	}

	imageId := disk.GetSourceImageId()
	if imageId == "" {
		return nil // Disk source is not Image, but rather Snapshot
	}

	image, err := c.SDK.Compute().Image().Get(ctx, &compute.GetImageRequest{ImageId: imageId})
	if err != nil {
		return err
	}

	res <- image
	return nil
}
