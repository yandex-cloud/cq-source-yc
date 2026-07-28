package baremetal

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/baremetal/v2"
	baremetalsdk "github.com/yandex-cloud/go-sdk/services/baremetal/v2"
)

// StandardImages is backed by the v2 Image service: the v1alpha StandardImage
// resource was renamed to Image in v2, and the list became folder-scoped.
func StandardImages() *schema.Table {
	return &schema.Table{
		Name:        "yc_baremetal_standard_images",
		Description: `https://yandex.cloud/docs/baremetal/api-ref/grpc/Image/list#yandex.cloud.baremetal.v2.Image`,
		Multiplex:   client.FolderMultiplex(client.ServiceBaremetal),
		Resolver:    fetchStandardImages,
		Transform:   client.TransformWithStruct(&baremetal.Image{}, client.PrimaryKeyIdTransformers("ImageId")...),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchStandardImages(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	it := baremetalsdk.NewImageClient(c.SDKv2).ImagesIterator(ctx, &baremetal.ListImagesRequest{FolderId: "baremetal-standard-images"})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
