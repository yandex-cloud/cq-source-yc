package organizationmanager

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1"
)

func fetchGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	orgID := parent.Item.(*organizationmanager.Organization).Id

	req := &organizationmanager.ListGroupsRequest{OrganizationId: orgID}
	it := c.Services.OrganizationManager.Group().GroupIterator(ctx, req)
	for it.Next() {
		res <- it.Value()
	}

	return nil
}
