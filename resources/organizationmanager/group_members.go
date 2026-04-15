package organizationmanager

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1"
)

func GroupMembers() *schema.Table {
	return &schema.Table{
		Name:        "yc_organizationmanager_group_members",
		Description: `https://yandex.cloud/ru/docs/organization/api-ref/grpc/Group/listMembers#yandex.cloud.organizationmanager.v1.GroupMember`,
		Resolver:    nil,
		Transform:   client.TransformWithStruct(&organizationmanager.GroupMember{}, transformers.WithPrimaryKeys("SubjectId")),
		Columns: schema.ColumnList{
			client.ParentIdColumn,
		},
	}
}

func fetchGroupMemebrs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	group, ok := parent.Item.(*organizationmanager.Group)
	if !ok {
		return fmt.Errorf("parent is not type of *organizationmanager.Group: %+v", group)
	}

	it := c.SDK.OrganizationManager().Group().GroupMembersIterator(ctx, &organizationmanager.ListGroupMembersRequest{GroupId: group.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
