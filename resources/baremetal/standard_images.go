package baremetal

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	baremetal "github.com/yandex-cloud/go-genproto/yandex/cloud/baremetal/v1alpha"
	baremetalsdk "github.com/yandex-cloud/go-sdk/services/baremetal/v1alpha"
)

// StandardImages is a global catalog of vendor-provided images (no folder/cloud
// scope), so the table runs once.
func StandardImages() *schema.Table {
	return &schema.Table{
		Name:        "yc_baremetal_standard_images",
		Description: `https://yandex.cloud/docs/baremetal/api-ref/grpc/StandardImage/list#yandex.cloud.baremetal.v1alpha.StandardImage`,
		Resolver:    fetchStandardImages,
		Transform:   client.TransformWithStruct(&baremetal.StandardImage{}, client.PrimaryKeyIdTransformer),
	}
}

func fetchStandardImages(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	it := baremetalsdk.NewStandardImageClient(c.SDKv2).Iterator(ctx, &baremetal.ListStandardImagesRequest{})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
