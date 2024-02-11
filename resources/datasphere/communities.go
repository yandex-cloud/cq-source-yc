package datasphere

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/cq-source-yc/resources/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/datasphere/v2"
)

func Communities() *schema.Table {
	return &schema.Table{
		Name:        "yc_datasphere_communities",
		Description: `https://cloud.yandex.ru/docs/datasphere/api-ref/grpc/community_service#Community3`,
		Multiplex:   client.OrganizationMultiplex,
		Resolver:    fetchCommunities,
		Transform:   client.TransformWithStruct(&datasphere.Community{}, client.PrimaryKeyIdTransformer),
		Relations:   schema.Tables{Projects(), access.DatasphereCommunitiesBindings()},
	}
}

func fetchCommunities(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.Datasphere().Community().CommunityIterator(ctx, &datasphere.ListCommunitiesRequest{OrganizationId: c.OrganizationId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
