package access

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/ydb/v1"
)

func YDBDatabasesAccessBindings() *schema.Table {
	return &schema.Table{
		Name:      "yc_access_bindings_ydb_databases",
		Resolver:  fetchYDBDatabasesAccessBindings,
		Transform: Transform,
		Columns: schema.ColumnList{
			client.ParentIdColumn,
		},
	}
}

func fetchYDBDatabasesAccessBindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	database, ok := parent.Item.(*ydb.Database)
	if !ok {
		return fmt.Errorf("parent is not type of *ydb.Database: %+v", database)
	}

	it := c.SDK.YDB().Database().DatabaseAccessBindingsIterator(ctx, &access.ListAccessBindingsRequest{
		ResourceId: database.Id,
	})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
