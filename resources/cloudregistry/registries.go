package cloudregistry

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	xaccess "github.com/yandex-cloud/cq-source-yc/resources/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/cloudregistry/v1"
)

func Registries() *schema.Table {
	base := "cloudregistry_registries"
	return &schema.Table{
		Name:        "yc_" + base,
		Description: `https://yandex.cloud/ru/docs/cloud-registry/api-ref/grpc/Registry/list#yandex.cloud.cloudregistry.v1.Registry`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchRegistries,
		Transform:   client.TransformWithStruct(&cloudregistry.Registry{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
		Relations: schema.Tables{
			xaccess.NewTable(base, fetchRegistryAccessBindings),
			IpPermissions(),
		},
	}
}

func fetchRegistries(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	it := c.SDK.CloudRegistry().Registry().RegistryIterator(ctx, &cloudregistry.ListRegistriesRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}

func fetchRegistryAccessBindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	registry, ok := parent.Item.(*cloudregistry.Registry)
	if !ok {
		return fmt.Errorf("parent in not type of *cloudregistry.Registry: %+v", registry)
	}

	it := c.SDK.CloudRegistry().Registry().RegistryAccessBindingsIterator(ctx, &access.ListAccessBindingsRequest{ResourceId: registry.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
