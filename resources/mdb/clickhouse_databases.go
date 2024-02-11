package mdb

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/clickhouse/v1"
)

func ClickhouseDatabases() *schema.Table {
	return &schema.Table{
		Name:        "yc_mdb_clickhouse_databases",
		Description: `https://cloud.yandex.ru/docs/managed-clickhouse/api-ref/grpc/database_service#Database1`,
		Resolver:    fetchClickhouseDatabases,
		Transform:   structNameClusterIdTransformer(&clickhouse.Database{}),
	}
}

func fetchClickhouseDatabases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	cluster, ok := parent.Item.(*clickhouse.Cluster)
	if !ok {
		return fmt.Errorf("parent is not type of *clickhouse.Cluster: %+v", cluster)
	}

	it := c.SDK.MDB().Clickhouse().Database().DatabaseIterator(ctx, &clickhouse.ListDatabasesRequest{ClusterId: cluster.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
