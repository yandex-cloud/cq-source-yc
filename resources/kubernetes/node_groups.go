package kubernetes

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/k8s/v1"
)

func NodeGroups() *schema.Table {
	return &schema.Table{
		Name:        "yc_kubernetes_node_groups",
		Description: ``,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchNodeGroups,
		Transform:   client.TransformWithStruct(&k8s.NodeGroup{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
		Relations: schema.Tables{Nodes()},
	}
}

func fetchNodeGroups(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.Kubernetes().NodeGroup().NodeGroupIterator(ctx, &k8s.ListNodeGroupsRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
