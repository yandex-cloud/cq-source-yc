package mdb

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/redis/v1"
)

func RedisHosts() *schema.Table {
	return &schema.Table{
		Name:        "yc_mdb_redis_hosts",
		Description: `https://cloud.yandex.ru/docs/managed-redis/api-ref/grpc/cluster_service#Host`,
		Resolver:    fetchRedisHosts,
		Transform:   structNameClusterIdTransformer(&redis.Host{}),
	}
}

func fetchRedisHosts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	cluster, ok := parent.Item.(*redis.Cluster)
	if !ok {
		return fmt.Errorf("parent is not type of *redis.Cluster: %+v", cluster)
	}

	it := c.SDK.MDB().Redis().Cluster().ClusterHostsIterator(ctx, &redis.ListClusterHostsRequest{ClusterId: cluster.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
