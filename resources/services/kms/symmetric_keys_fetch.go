package kms

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/kms/v1"
)

func fetchSymmetricKeys(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	req := &kms.ListSymmetricKeysRequest{FolderId: c.MultiplexedResourceId}
	it := c.Services.KMS.SymmetricKey().SymmetricKeyIterator(ctx, req)
	for it.Next() {
		res <- it.Value()
	}

	return nil
}
