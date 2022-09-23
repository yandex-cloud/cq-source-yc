package iam

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
)

func fetchServiceAccounts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	req := &iam.ListServiceAccountsRequest{FolderId: c.MultiplexedResourceId}
	it := c.Services.IAM.ServiceAccount().ServiceAccountIterator(ctx, req)
	for it.Next() {
		res <- it.Value()
	}

	return nil
}
