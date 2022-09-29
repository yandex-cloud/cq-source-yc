package access_bindings

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func fetchByOrganization(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	return fetchAccessBindings(func(c *client.Client) accessBindingsClient {
		return c.Services.OrganizationManager.Organization()
	})(ctx, meta, parent, res)
}
