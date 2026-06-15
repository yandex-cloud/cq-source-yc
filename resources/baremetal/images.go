package baremetal

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	baremetal "github.com/yandex-cloud/go-genproto/yandex/cloud/baremetal/v1alpha"
	baremetalsdk "github.com/yandex-cloud/go-sdk/services/baremetal/v1alpha"
)

func Images() *schema.Table {
	return &schema.Table{
		Name:        "yc_baremetal_images",
		Description: `https://yandex.cloud/docs/baremetal/api-ref/grpc/Image/list#yandex.cloud.baremetal.v1alpha.Image`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchImages,
		Transform:   client.TransformWithStruct(&baremetal.Image{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchImages(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	it := baremetalsdk.NewImageClient(c.SDKv2).Iterator(ctx, &baremetal.ListImagesRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
