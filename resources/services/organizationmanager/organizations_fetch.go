package organizationmanager

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1"
)

func fetchOrganizations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	org, err := c.Services.OrganizationManager.Organization().Get(ctx,
		&organizationmanager.GetOrganizationRequest{OrganizationId: c.MultiplexedResourceId},
	)
	if err != nil {
		return err
	}

	res <- org
	return nil
}
