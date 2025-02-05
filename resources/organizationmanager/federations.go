package organizationmanager

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1/saml"
)

func Federations() *schema.Table {
	return &schema.Table{
		Name:        "yc_organizationmanager_federations",
		Description: `https://yandex.cloud/ru/docs/organization/saml/api-ref/grpc/Federation/list#yandex.cloud.organizationmanager.v1.saml.Federation`,
		Multiplex:   client.OrganizationMultiplex,
		Resolver:    fetchFederations,
		Transform:   client.TransformWithStruct(&saml.Federation{}, client.PrimaryKeyIdTransformer),
	}
}

func fetchFederations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	orgId := c.OrganizationId

	it := c.SDK.OrganizationManagerSAML().Federation().FederationIterator(ctx, &saml.ListFederationsRequest{OrganizationId: orgId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
