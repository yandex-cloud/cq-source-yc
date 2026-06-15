package trino

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/trino/v1"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:        "yc_trino_clusters",
		Description: `https://yandex.cloud/docs/managed-trino/api-ref/grpc/Cluster/list#yandex.cloud.trino.v1.Cluster`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchClusters,
		Transform:   client.TransformWithStruct(&trino.Cluster{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
		Relations: schema.Tables{
			Catalogs(),
		},
	}
}

func fetchClusters(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	it := c.SDK.Trino().Cluster().ClusterIterator(ctx, &trino.ListClustersRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
