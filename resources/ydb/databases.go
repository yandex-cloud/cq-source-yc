package ydb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/ydb/v1"
)

func Databases() *schema.Table {
	return &schema.Table{
		Name:        "yc_ydb_databases",
		Description: `https://cloud.yandex.ru/docs/ydb/api-ref/grpc/database_service#Database1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchDatabases,
		Transform:   client.TransformWithStruct(&ydb.Database{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchDatabases(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.YDB().Database().DatabaseIterator(ctx, &ydb.ListDatabasesRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
