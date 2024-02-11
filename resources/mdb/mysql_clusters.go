package mdb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mysql/v1"
)

func MySQLClusters() *schema.Table {
	return &schema.Table{
		Name:        "yc_mdb_mysql_clusters",
		Description: `https://cloud.yandex.ru/docs/managed-mysql/api-ref/grpc/cluster_service#Cluster1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchMySQLClusters,
		Transform:   client.TransformWithStruct(&mysql.Cluster{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
		Relations: schema.Tables{
			MySQLDatabases(),
			MySQLHosts(),
			MySQLUsers(),
		},
	}
}

func fetchMySQLClusters(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.MDB().MySQL().Cluster().ClusterIterator(ctx, &mysql.ListClustersRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
