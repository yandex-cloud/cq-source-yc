package containerregistry

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/cq-source-yc/resources/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/containerregistry/v1"
)

func Registries() *schema.Table {
	return &schema.Table{
		Name:        "yc_containerregistry_registries",
		Description: `https://cloud.yandex.ru/docs/container-registry/api-ref/grpc/registry_service#Registry1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchRegistries,
		Transform:   client.TransformWithStruct(&containerregistry.Registry{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
		Relations: schema.Tables{Repositories(), access.RegistriesAccessBindings()},
	}
}

func fetchRegistries(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.ContainerRegistry().Registry().RegistryIterator(ctx, &containerregistry.ListRegistriesRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
