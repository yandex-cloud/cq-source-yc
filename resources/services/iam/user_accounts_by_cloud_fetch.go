package iam

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func fetchUserAccountsByCloud(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	return fetchUserAccounts(meta.(*client.Client).Services.ResourceManager.Cloud())(ctx, meta, parent, res)

}
