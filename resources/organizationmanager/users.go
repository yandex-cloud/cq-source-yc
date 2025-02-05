package organizationmanager

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:        "yc_organizationmanager_users",
		Description: `https://cloud.yandex.ru/docs/organization/api-ref/grpc/user_service#OrganizationUser`,
		Multiplex:   client.OrganizationMultiplex,
		Resolver:    fetchUsers,
		Transform:   client.TransformWithStruct(&organizationmanager.ListMembersResponse_OrganizationUser{}, transformers.WithUnwrapStructFields("SubjectClaims"), transformers.WithPrimaryKeys("SubjectClaims.Sub")),
		Columns: schema.ColumnList{
			client.OrganiztionIdColumn,
		},
	}
}

func fetchUsers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	orgId := c.OrganizationId

	it := c.SDK.OrganizationManager().User().UserMembersIterator(ctx, &organizationmanager.ListMembersRequest{OrganizationId: orgId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
