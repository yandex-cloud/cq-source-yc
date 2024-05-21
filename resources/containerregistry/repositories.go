package containerregistry

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/cq-source-yc/resources/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/containerregistry/v1"
)

func Repositories() *schema.Table {
	return &schema.Table{
		Name:        "yc_containerregistry_repositories",
		Description: `https://cloud.yandex.ru/docs/container-registry/api-ref/grpc/repository_service#Repository2`,
		Resolver:    fetchRepositories,
		Transform:   client.TransformWithStruct(&containerregistry.Repository{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
			schema.Column{
				Name:     "registry_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
		Relations: schema.Tables{access.RepositoriesAccessBindings()},
	}
}

func fetchRepositories(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	registry, ok := parent.Item.(*containerregistry.Registry)
	if !ok {
		return fmt.Errorf("parent is not type of *containerregistry.Registry: %+v", registry)
	}

	it := c.SDK.ContainerRegistry().Repository().RepositoryIterator(ctx, &containerregistry.ListRepositoriesRequest{RegistryId: registry.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
