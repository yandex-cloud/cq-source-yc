package access_bindings

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func fetchByCloud(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	return fetchAccessBindings(func(c *client.Client) accessBindingsClient {
		return c.Services.ResourceManager.Cloud()
	})(ctx, meta, parent, res)
}
