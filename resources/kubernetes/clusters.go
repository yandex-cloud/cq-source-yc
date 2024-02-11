package kubernetes

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/k8s/v1"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:        "yc_kubernetes_clusters",
		Description: `https://cloud.yandex.ru/docs/managed-kubernetes/api-ref/grpc/cluster_service#Cluster1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchClusters,
		Transform:   client.TransformWithStruct(&k8s.Cluster{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchClusters(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.Kubernetes().Cluster().ClusterIterator(ctx, &k8s.ListClustersRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
