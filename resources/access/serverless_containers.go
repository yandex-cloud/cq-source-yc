package access

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/containers/v1"
)

func ServerlessContainersAccessBindings() *schema.Table {
	return &schema.Table{
		Name:      "yc_access_bindings_serverless_containers",
		Resolver:  fetchServerlessContainersAccessBindings,
		Transform: Transform,
		Columns: schema.ColumnList{
			client.ParentIdColumn,
		},
	}
}

func fetchServerlessContainersAccessBindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	container, ok := parent.Item.(*containers.Container)
	if !ok {
		return fmt.Errorf("parent is not type of *containers.Container: %+v", container)
	}

	it := c.SDK.Serverless().Containers().Container().ContainerAccessBindingsIterator(ctx, &access.ListAccessBindingsRequest{ResourceId: container.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
