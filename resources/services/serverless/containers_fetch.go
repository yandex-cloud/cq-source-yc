package serverless

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/containers/v1"
)

func fetchContainers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.Services.Serverless.Containers().Container().ContainerIterator(ctx,
		&containers.ListContainersRequest{FolderId: c.MultiplexedResourceId},
	)
	for it.Next() {
		res <- it.Value()
	}

	return nil
}
