package mdb

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/clickhouse/v1"
)

func ClickhouseUsers() *schema.Table {
	return &schema.Table{
		Name:        "yc_mdb_clickhouse_users",
		Description: `https://cloud.yandex.ru/docs/managed-clickhouse/api-ref/grpc/user_service#User1`,
		Resolver:    fetchClickhouseUsers,
		Transform:   structNameClusterIdTransformer(&clickhouse.User{}),
	}
}

func fetchClickhouseUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	cluster, ok := parent.Item.(*clickhouse.Cluster)
	if !ok {
		return fmt.Errorf("parent is not type of *clickhouse.Cluster: %+v", cluster)
	}

	it := c.SDK.MDB().Clickhouse().User().UserIterator(ctx, &clickhouse.ListUsersRequest{ClusterId: cluster.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
