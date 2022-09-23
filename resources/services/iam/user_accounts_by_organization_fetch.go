package iam

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func fetchUserAccountsByOrganization(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	return fetchUserAccounts(meta.(*client.Client).Services.OrganizationManager.Organization())(ctx, meta, parent, res)
}
