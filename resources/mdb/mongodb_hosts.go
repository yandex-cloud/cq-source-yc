package mdb

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mongodb/v1"
)

func MongoDBHosts() *schema.Table {
	return &schema.Table{
		Name:        "yc_mdb_mongodb_hosts",
		Description: `https://cloud.yandex.ru/docs/managed-mongodb/api-ref/grpc/cluster_service#Host`,
		Resolver:    fetchMongoDBHosts,
		Transform:   structNameClusterIdTransformer(&mongodb.Host{}),
	}
}

func fetchMongoDBHosts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	cluster, ok := parent.Item.(*mongodb.Cluster)
	if !ok {
		return fmt.Errorf("parent is not type of *mongodb.Cluster: %+v", cluster)
	}

	it := c.SDK.MDB().MongoDB().Cluster().ClusterHostsIterator(ctx, &mongodb.ListClusterHostsRequest{ClusterId: cluster.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
