package resources

import (
	"context"
	"github.com/yandex-cloud/cq-provider-yandex/tools"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
)

func ComputeImages() *schema.Table {
	gen, err := tools.NewTableGenerator(
		"yandex_compute_images",
		"Compute",
		"Image",
		"resources/proto/image.proto",
		tools.GetCommonDefaultColumns("image"),
		tools.IgnoredColumns{},
		fetchComputeImages,
	)
	if err != nil {
		return nil
	}
	table, err := gen.Generate()
	if err != nil {
		return nil
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
