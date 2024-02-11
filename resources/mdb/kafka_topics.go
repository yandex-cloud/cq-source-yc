package mdb

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/kafka/v1"
)

func KafkaTopics() *schema.Table {
	return &schema.Table{
		Name:        "yc_mdb_kafka_topics",
		Description: `https://cloud.yandex.ru/docs/managed-kafka/api-ref/grpc/topic_service#Topic1`,
		Resolver:    fetchKafkaTopics,
		Transform:   structNameClusterIdTransformer(&kafka.Topic{}),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchKafkaTopics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	cluster, ok := parent.Item.(*kafka.Cluster)
	if !ok {
		return fmt.Errorf("parent is not type of *kafka.Cluster: %+v", cluster)
	}

	it := c.SDK.MDB().Kafka().Topic().TopicIterator(ctx, &kafka.ListTopicsRequest{ClusterId: cluster.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
