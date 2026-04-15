package organizationmanager

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1"
)

func Groups() *schema.Table {
	return &schema.Table{
		Name:        "yc_organizationmanager_groups",
		Description: `https://yandex.cloud/ru/docs/organization/api-ref/grpc/Group/list#yandex.cloud.organizationmanager.v1.Group`,
		Multiplex:   client.OrganizationMultiplex,
		Resolver:    fetchGroups,
		Transform:   client.TransformWithStruct(&organizationmanager.Group{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.OrganiztionIdColumn,
		},
		Relations: schema.Tables{GroupMembers()},
	}
}

func fetchGroups(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.OrganizationManager().Group().GroupIterator(ctx, &organizationmanager.ListGroupsRequest{OrganizationId: c.OrganizationId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
