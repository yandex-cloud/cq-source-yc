package k8s

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/k8s/v1"
)

func fetchNodes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	groupID := parent.Item.(*k8s.NodeGroup).Id
	req := &k8s.ListNodeGroupNodesRequest{NodeGroupId: groupID}
	it := c.Services.K8S.NodeGroup().NodeGroupNodesIterator(ctx, req)
	for it.Next() {
		res <- it.Value()
	}

	return nil
}
