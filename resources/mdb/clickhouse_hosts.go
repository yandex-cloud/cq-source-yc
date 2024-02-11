package mdb

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/clickhouse/v1"
)

func ClickhouseHosts() *schema.Table {
	return &schema.Table{
		Name:        "yc_mdb_clickhouse_hosts",
		Description: `https://cloud.yandex.ru/docs/managed-clickhouse/api-ref/grpc/cluster_service#Host`,
		Resolver:    fetchClickhouseHosts,
		Transform:   structNameClusterIdTransformer(&clickhouse.Host{}),
	}
}

func fetchClickhouseHosts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	cluster, ok := parent.Item.(*clickhouse.Cluster)
	if !ok {
		return fmt.Errorf("parent is not type of *clickhouse.Cluster: %+v", cluster)
	}

	it := c.SDK.MDB().Clickhouse().Cluster().ClusterHostsIterator(ctx, &clickhouse.ListClusterHostsRequest{ClusterId: cluster.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
