package access

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/containerregistry/v1"
)

func RepositoriesAccessBindings() *schema.Table {
	return &schema.Table{
		Name:        "yc_access_bindings_containerregistry_repositories",
		Description: ``,
		Resolver:    fetchRepositoriesAccessBindings,
		Transform:   Transform,
		Columns: schema.ColumnList{
			client.ParentIdColumn,
		},
	}
}

func fetchRepositoriesAccessBindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	repository, ok := parent.Item.(*containerregistry.Repository)
	if !ok {
		return fmt.Errorf("parent is not type of *containerregistry.Repository: %+v", repository)
	}

	it := c.SDK.ContainerRegistry().Repository().RepositoryAccessBindingsIterator(ctx, &access.ListAccessBindingsRequest{ResourceId: repository.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
