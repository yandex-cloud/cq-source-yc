package compute

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
)

func fetchImages(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	req := &compute.ListImagesRequest{FolderId: c.MultiplexedResourceId}
	it := c.Services.Compute.Image().ImageIterator(ctx, req)
	for it.Next() {
		res <- it.Value()
	}

	return nil
}
