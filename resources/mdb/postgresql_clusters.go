package mdb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/postgresql/v1"
)

func PostgreSQLClusters() *schema.Table {
	return &schema.Table{
		Name:        "yc_mdb_postgresql_clusters",
		Description: `https://cloud.yandex.ru/docs/managed-postgresql/api-ref/grpc/cluster_service#Cluster1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchPostgreSQLClusters,
		Transform:   client.TransformWithStruct(&postgresql.Cluster{}, client.PrimaryKeyIdTransformer),
		Relations:   schema.Tables{PostgreSQLDatabases(), PostgreSQLUsers(), PostgreSQLHosts()},
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchPostgreSQLClusters(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	folderId := c.FolderId

	it := c.SDK.MDB().PostgreSQL().Cluster().ClusterIterator(ctx, &postgresql.ListClustersRequest{FolderId: folderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
