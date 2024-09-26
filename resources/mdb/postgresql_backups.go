package mdb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/postgresql/v1"
)

func PostgreSQLBackups() *schema.Table {
	return &schema.Table{
		Name:        "yc_mdb_postgresql_backups",
		Description: `https://yandex.cloud/ru/docs/managed-postgresql/api-ref/grpc/backup_service#Backup`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchPostgreSQLBackups,
		Transform:   client.TransformWithStruct(&postgresql.Backup{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchPostgreSQLBackups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	folderId := c.FolderId

	it := c.SDK.MDB().PostgreSQL().Backup().BackupIterator(ctx, &postgresql.ListBackupsRequest{FolderId: folderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
