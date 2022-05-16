package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1"
)

func OrganizationManagerOrganizations() *schema.Table {
	return &schema.Table{
		Name:        "yandex_organizationmanager_organizations",
		Resolver:    fetchOrganizationManagerOrganizations,
		Multiplex:   client.EmptyMultiplex,
		IgnoreError: client.IgnoreErrorHandler,
		Columns: []schema.Column{
			{
				Name:            "id",
				Type:            schema.TypeString,
				Description:     "ID of the resource.",
				Resolver:        client.ResolveResourceId,
				CreationOptions: schema.ColumnCreationOptions{NotNull: true, Unique: true},
			},
			{
				Name:        "created_at",
				Type:        schema.TypeTimestamp,
				Description: "",
				Resolver:    client.ResolveAsTime,
			},
			{
				Name:        "name",
				Type:        schema.TypeString,
				Description: "Name of the organization. 3-63 characters long.",
				Resolver:    schema.PathResolver("Name"),
			},
			{
				Name:        "description",
				Type:        schema.TypeString,
				Description: "Description of the organization. 0-256 characters long.",
				Resolver:    schema.PathResolver("Description"),
			},
			{
				Name:        "title",
				Type:        schema.TypeString,
				Description: "Display name of the organization. 0-256 characters long.",
				Resolver:    schema.PathResolver("Title"),
			},
		},
	}
}

func fetchOrganizationManagerOrganizations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	req := &organizationmanager.ListOrganizationsRequest{}
	it := c.Services.OrganizationManager.Organization().OrganizationIterator(ctx, req)
	for it.Next() {
		res <- it.Value()
	}

	return nil
}
