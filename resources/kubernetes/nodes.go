package kubernetes

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/k8s/v1"
)

func Nodes() *schema.Table {
	return &schema.Table{
		Name:        "yc_kubernetes_nodes",
		Description: ``,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchNodes,
		Transform:   client.TransformWithStruct(&k8s.Node{}, transformers.WithUnwrapStructFields("CloudStatus"), transformers.WithPrimaryKeys("CloudStatus.Id")),
		Columns: schema.ColumnList{
			{
				Name:     "node_group_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}

func fetchNodes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nodeGroup, ok := parent.Item.(*k8s.NodeGroup)
	if !ok {
		return fmt.Errorf("parent is not type of *k8s.NodeGroup: %+v", nodeGroup)
	}

	it := c.SDK.Kubernetes().NodeGroup().NodeGroupNodesIterator(ctx, &k8s.ListNodeGroupNodesRequest{NodeGroupId: nodeGroup.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
