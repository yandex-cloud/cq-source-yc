package mdb

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/kafka/v1"
)

func KafkaUsers() *schema.Table {
	return &schema.Table{
		Name:        "yc_mdb_kafka_users",
		Description: `https://cloud.yandex.ru/docs/managed-kafka/api-ref/grpc/user_service#User1`,
		Resolver:    fetchKafkaUsers,
		Transform:   structNameClusterIdTransformer(&kafka.User{}),
	}
}

func fetchKafkaUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	cluster, ok := parent.Item.(*kafka.Cluster)
	if !ok {
		return fmt.Errorf("parent is not type of *kafka.Cluster: %+v", cluster)
	}

	it := c.SDK.MDB().Kafka().User().UserIterator(ctx, &kafka.ListUsersRequest{ClusterId: cluster.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
