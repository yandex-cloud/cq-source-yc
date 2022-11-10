package organizationmanager

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1"
)

func fetchOrganizations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	req := &organizationmanager.ListOrganizationsRequest{}
	it := c.Services.OrganizationManager.Organization().OrganizationIterator(ctx, req)
	for it.Next() {
		res <- it.Value()
	}

	return nil
}
