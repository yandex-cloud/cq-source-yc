package access_bindings

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
)

func fetchByOrganization(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	_client := c.Services.OrganizationManager.Organization()

	for {
		resp, err := _client.ListAccessBindings(ctx, &access.ListAccessBindingsRequest{
			ResourceId: c.MultiplexedResourceId,
		})
		if err != nil {
			return err
		}

		res <- resp.GetAccessBindings()

		if resp.GetNextPageToken() == "" {
			return nil
		}
	}
}
