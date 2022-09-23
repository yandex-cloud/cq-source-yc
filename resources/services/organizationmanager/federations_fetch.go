package organizationmanager

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1/saml"
)

func fetchFederations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	req := &saml.ListFederationsRequest{OrganizationId: c.MultiplexedResourceId}
	it := c.Services.OrganizationManagerSAML.Federation().FederationIterator(ctx, req)
	for it.Next() {
		res <- it.Value()
	}

	return nil
}
