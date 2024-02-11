package mdb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/greenplum/v1"
)

func GreenplumClusters() *schema.Table {
	return &schema.Table{
		Name:        "yc_mdb_greenplum_clusters",
		Description: `https://cloud.yandex.ru/docs/managed-greenplum/api-ref/grpc/cluster_service#Cluster1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchGreenplumClusters,
		Transform:   client.TransformWithStruct(&greenplum.Cluster{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
		Relations: schema.Tables{GreenplumHosts()},
	}
}

func fetchGreenplumClusters(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	folderId := c.FolderId

	it := c.SDK.MDB().Greenplum().Cluster().ClusterIterator(ctx, &greenplum.ListClustersRequest{FolderId: folderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
