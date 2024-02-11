package access

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/datasphere/v2"
)

func DatasphereCommunitiesBindings() *schema.Table {
	return &schema.Table{
		Name:        "yc_access_bindings_datasphere_communities",
		Description: ``,
		Resolver:    fetchDatasphereCommunitiesBindings,
		Transform:   Transform,
		Columns: schema.ColumnList{
			client.ParentIdColumn,
		},
	}
}

func fetchDatasphereCommunitiesBindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	community, ok := parent.Item.(*datasphere.Community)
	if !ok {
		return fmt.Errorf("parent is not type of *datasphere.Community: %+v", community)
	}

	it := c.SDK.Datasphere().Community().CommunityAccessBindingsIterator(ctx, &access.ListAccessBindingsRequest{ResourceId: community.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
