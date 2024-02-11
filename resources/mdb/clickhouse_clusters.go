package mdb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/clickhouse/v1"
)

func ClickhouseClusters() *schema.Table {
	return &schema.Table{
		Name:        "yc_mdb_clickhouse_clusters",
		Description: `https://cloud.yandex.ru/docs/managed-clickhouse/api-ref/grpc/cluster_service#Cluster1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchClickhouseClusters,
		Transform:   client.TransformWithStruct(&clickhouse.Cluster{}, client.PrimaryKeyIdTransformer),
		Relations:   schema.Tables{ClickhouseDatabases(), ClickhouseUsers(), ClickhouseHosts()},
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchClickhouseClusters(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	folderId := c.FolderId

	it := c.SDK.MDB().Clickhouse().Cluster().ClusterIterator(ctx, &clickhouse.ListClustersRequest{FolderId: folderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
