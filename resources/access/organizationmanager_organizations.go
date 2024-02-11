package access

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
)

func OrganizationsAccessBindings() *schema.Table {
	return &schema.Table{
		Name:      "yc_access_bindings_organizationmanager_organizations",
		Resolver:  fetchOrganizationsAccessBindings,
		Multiplex: client.OrganizationMultiplex,
		Transform: Transform,
		Columns: schema.ColumnList{
			client.MultiplexedResourceIdColumn,
		},
	}
}

func fetchOrganizationsAccessBindings(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	orgId := c.OrganizationId

	it := c.SDK.OrganizationManager().Organization().OrganizationAccessBindingsIterator(ctx, &access.ListAccessBindingsRequest{ResourceId: orgId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
