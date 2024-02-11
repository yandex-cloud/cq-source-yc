package mdb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/opensearch/v1"
)

func OpenSearchClusters() *schema.Table {
	return &schema.Table{
		Name:        "yc_mdb_opensearch_clusters",
		Description: `https://cloud.yandex.ru/docs/managed-opensearch/api-ref/grpc/cluster_service#Cluster1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchOpenSearchClusters,
		Transform:   client.TransformWithStruct(&opensearch.Cluster{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
		Relations: schema.Tables{OpensearchHosts(), OpenSearchAuthSettings()},
	}
}

func fetchOpenSearchClusters(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.MDB().OpenSearch().Cluster().ClusterIterator(ctx, &opensearch.ListClustersRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
