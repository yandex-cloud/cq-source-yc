package mdb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/redis/v1"
)

func RedisClusters() *schema.Table {
	return &schema.Table{
		Name:        "yc_mdb_redis_clusters",
		Description: `https://cloud.yandex.ru/docs/managed-mongodb/api-ref/grpc/cluster_service#Cluster1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchRedisClusters,
		Transform:   client.TransformWithStruct(&redis.Cluster{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
		Relations: schema.Tables{RedisHosts()},
	}
}

func fetchRedisClusters(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	folderId := c.FolderId

	it := c.SDK.MDB().Redis().Cluster().ClusterIterator(ctx, &redis.ListClustersRequest{FolderId: folderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
