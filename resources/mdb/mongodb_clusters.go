package mdb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mongodb/v1"
)

func MongoDBClusters() *schema.Table {
	return &schema.Table{
		Name:        "yc_mdb_mongodb_clusters",
		Description: `https://cloud.yandex.ru/docs/managed-mongodb/api-ref/grpc/cluster_service#Cluster1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchMongoDBClusters,
		Transform:   client.TransformWithStruct(&mongodb.Cluster{}, client.PrimaryKeyIdTransformer),
		Relations: schema.Tables{
			MongoDBDatabases(),
			MongoDBHosts(),
			MongoDBUsers(),
		},
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchMongoDBClusters(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	folderId := c.FolderId

	it := c.SDK.MDB().MongoDB().Cluster().ClusterIterator(ctx, &mongodb.ListClustersRequest{FolderId: folderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
