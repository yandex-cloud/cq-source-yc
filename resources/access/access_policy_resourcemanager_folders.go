package access

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
)

func FoldersAccessPolicyBindings() *schema.Table {
	return &schema.Table{
		Name:        "yc_access_policy_bindings_resourcemanager_folders",
		Title:       "YC Access Policy Bindings for Folders",
		Description: `https://yandex.cloud/docs/iam/concepts/access-control/#access-policies`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchFoldersAccessPolicyBindings,
		Transform:   AccessPolicyTransform,
		Columns: schema.ColumnList{
			client.MultiplexedResourceIdColumn,
		},
	}
}

func fetchFoldersAccessPolicyBindings(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	it := c.SDK.ResourceManager().Folder().FolderAccessPolicyBindingsIterator(ctx, &access.ListAccessPolicyBindingsRequest{ResourceId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
