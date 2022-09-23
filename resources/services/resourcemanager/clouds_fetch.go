package resourcemanager

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
)

func fetchClouds(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	cloud, err := c.Services.ResourceManager.Cloud().Get(ctx, &resourcemanager.GetCloudRequest{CloudId: c.MultiplexedResourceId})
	if err != nil {
		return err
	}

	res <- cloud

	return nil
}
