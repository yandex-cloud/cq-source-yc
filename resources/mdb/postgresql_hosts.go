package mdb

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/postgresql/v1"
)

func PostgreSQLHosts() *schema.Table {
	return &schema.Table{
		Name:        "yc_mdb_postgresql_hosts",
		Description: `https://cloud.yandex.ru/docs/managed-postgresql/api-ref/grpc/cluster_service#Host`,
		Resolver:    fetchPostgreSQLHosts,
		Transform:   structNameClusterIdTransformer(&postgresql.Host{}),
	}
}

func fetchPostgreSQLHosts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	cluster, ok := parent.Item.(*postgresql.Cluster)
	if !ok {
		return fmt.Errorf("parent is not type of *postgresql.Cluster: %+v", cluster)
	}

	it := c.SDK.MDB().PostgreSQL().Cluster().ClusterHostsIterator(ctx, &postgresql.ListClusterHostsRequest{ClusterId: cluster.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
