package resources

import (
	"context"

	"github.com/yandex-cloud/cq-provider-yandex/tools"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
)

func ComputeImages() *schema.Table {
	table, err := tools.GenerateTable(
		tools.WithTableName("yandex_computes_images"),
		tools.WithProtoFile("Image", "yandex/cloud/compute/v1/image.proto", "cloudapi"),
		tools.WithResolver(fetchComputeImages),
	)
	if err != nil {
		return &schema.Table{}
	}
	return table
}

func fetchComputeImages(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	locations := []string{c.FolderId}

	for _, f := range locations {
		req := &compute.ListImagesRequest{FolderId: f}
		it := c.Services.Compute.Image().ImageIterator(ctx, req)
		for it.Next() {
			res <- it.Value()
		}
	}
	return nil
}
