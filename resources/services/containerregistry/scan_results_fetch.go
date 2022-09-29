package containerregistry

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/containerregistry/v1"
	"golang.org/x/sync/errgroup"
)

func fetchScanResults(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	g := errgroup.Group{}
	ch := make(chan *containerregistry.Image)

	g.Go(func() error {
		defer close(ch)
		req := &containerregistry.ListImagesRequest{FolderId: c.MultiplexedResourceId}
		it := c.Services.ContainerRegistry.Image().ImageIterator(ctx, req)
		for it.Next() {
			ch <- it.Value()
		}
		return nil
	})

	g.Go(func() error {
		for image := range ch {
			req := &containerregistry.ListScanResultsRequest{
				Id: &containerregistry.ListScanResultsRequest_ImageId{
					ImageId: image.Id,
				},
			}
			it := c.Services.ContainerRegistry.Scanner().ScannerIterator(ctx, req)
			for it.Next() {
				res <- it.Value()
			}
		}
		return nil
	})

	return g.Wait()
}
