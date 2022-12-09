package lockbox

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/lockbox/v1"
)

func fetchSecrets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	req := &lockbox.ListSecretsRequest{FolderId: c.MultiplexedResourceId}
	it := c.Services.LockboxSecret.Secret().SecretIterator(ctx, req)
	for it.Next() {
		value := it.Value()
		res <- value
	}

	return nil
}
