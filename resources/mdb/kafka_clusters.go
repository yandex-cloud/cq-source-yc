package mdb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/kafka/v1"
)

func KafkaClusters() *schema.Table {
	return &schema.Table{
		Name:        "yc_mdb_kafka_clusters",
		Description: `https://cloud.yandex.ru/docs/managed-kafka/api-ref/grpc/cluster_service#Cluster1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchKafkaClusters,
		Transform:   client.TransformWithStruct(&kafka.Cluster{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
		Relations: schema.Tables{KafkaTopics(), KafkaUsers()},
	}
}

func fetchKafkaClusters(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	folderId := c.FolderId

	it := c.SDK.MDB().Kafka().Cluster().ClusterIterator(ctx, &kafka.ListClustersRequest{FolderId: folderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
