package baremetal

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/baremetal/v2"
	baremetalsdk "github.com/yandex-cloud/go-sdk/services/baremetal/v2"
)

// Images is backed by the v2 BootImage service: the v1alpha Image resource
// (user-uploaded images) was renamed to BootImage in v2.
func Images() *schema.Table {
	return &schema.Table{
		Name:        "yc_baremetal_images",
		Description: `https://yandex.cloud/docs/baremetal/api-ref/grpc/BootImage/list#yandex.cloud.baremetal.v2.BootImage`,
		Multiplex:   client.FolderMultiplex(client.ServiceBaremetal),
		Resolver:    fetchImages,
		Transform:   client.TransformWithStruct(&baremetal.BootImage{}, client.PrimaryKeyIdTransformers("BootImageId")...),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchImages(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	it := baremetalsdk.NewBootImageClient(c.SDKv2).BootImagesIterator(ctx, &baremetal.ListBootImagesRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
