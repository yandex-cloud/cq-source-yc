package organizationmanager

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1"
)

func Organizations() *schema.Table {
	return &schema.Table{
		Name:        "yc_organizationmanager_organizations",
		Description: `https://cloud.yandex.ru/docs/organization/api-ref/grpc/organization_service#Organization1`,
		Resolver:    fetchOrganizations,
		Multiplex:   client.OrganizationMultiplex,
		Transform:   client.TransformWithStruct(&organizationmanager.Organization{}, client.PrimaryKeyIdTransformer),
	}
}

func fetchOrganizations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	org, err := c.SDK.OrganizationManager().Organization().Get(ctx, &organizationmanager.GetOrganizationRequest{OrganizationId: c.OrganizationId})
	if err != nil {
		return err
	}
	res <- org

	return nil
}
