package serverless

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/cq-source-yc/resources/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/containers/v1"
)

func ContainersContainers() *schema.Table {
	return &schema.Table{
		Name:        "yc_serverless_containers_containers",
		Description: `https://cloud.yandex.ru/docs/serverless-containers/containers/api-ref/grpc/container_service#Container1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchContainersContainers,
		Transform:   client.TransformWithStruct(&containers.Container{}, client.PrimaryKeyIdTransformer),
		Relations:   schema.Tables{access.ServerlessContainersAccessBindings()},
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchContainersContainers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.Serverless().Containers().Container().ContainerIterator(ctx, &containers.ListContainersRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
