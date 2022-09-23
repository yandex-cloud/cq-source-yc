package k8s

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/k8s/v1"
)

func fetchNodeGroups(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	req := &k8s.ListNodeGroupsRequest{FolderId: c.MultiplexedResourceId}
	it := c.Services.K8S.NodeGroup().NodeGroupIterator(ctx, req)
	for it.Next() {
		res <- it.Value()
	}

	return nil
}
