package mdb

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mysql/v1"
)

func MySQLUsers() *schema.Table {
	return &schema.Table{
		Name:        "yc_mdb_mysql_users",
		Description: `https://cloud.yandex.ru/docs/managed-mysql/api-ref/grpc/user_service#User1`,
		Resolver:    fetchMySQLUsers,
		Transform:   structNameClusterIdTransformer(&mysql.User{}),
	}
}

func fetchMySQLUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	cluster, ok := parent.Item.(*mysql.Cluster)
	if !ok {
		return fmt.Errorf("parent is not type of *mysql.Cluster: %+v", cluster)
	}

	it := c.SDK.MDB().MySQL().User().UserIterator(ctx, &mysql.ListUsersRequest{ClusterId: cluster.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
