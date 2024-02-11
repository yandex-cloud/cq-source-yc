package mdb

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mysql/v1"
)

func MySQLDatabases() *schema.Table {
	return &schema.Table{
		Name:        "yc_mdb_mysql_databases",
		Description: `https://cloud.yandex.ru/docs/managed-mysql/api-ref/grpc/database_service#Database1`,
		Resolver:    fetchMySQLDatabases,
		Transform:   structNameClusterIdTransformer(&mysql.Database{}),
	}
}

func fetchMySQLDatabases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	cluster, ok := parent.Item.(*mysql.Cluster)
	if !ok {
		return fmt.Errorf("parent is not type of *mysql.Cluster: %+v", cluster)
	}

	it := c.SDK.MDB().MySQL().Database().DatabaseIterator(ctx, &mysql.ListDatabasesRequest{ClusterId: cluster.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
