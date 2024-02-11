package access

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/containerregistry/v1"
)

func RegistriesAccessBindings() *schema.Table {
	return &schema.Table{
		Name:        "yc_access_bindings_containerregistry_registries",
		Description: ``,
		Resolver:    fetchRegistriesAccessBindings,
		Transform:   Transform,
		Columns: schema.ColumnList{
			client.ParentIdColumn,
		},
	}
}

func fetchRegistriesAccessBindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	registry, ok := parent.Item.(*containerregistry.Registry)
	if !ok {
		return fmt.Errorf("parent is not type of *containerregistry.Registry: %+v", registry)
	}

	it := c.SDK.ContainerRegistry().Registry().RegistryAccessBindingsIterator(ctx, &access.ListAccessBindingsRequest{ResourceId: registry.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
