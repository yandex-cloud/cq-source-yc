package mdb

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/opensearch/v1"
)

func OpensearchHosts() *schema.Table {
	return &schema.Table{
		Name:        "yc_mdb_opensearch_hosts",
		Description: `https://cloud.yandex.ru/docs/managed-opensearch/api-ref/grpc/cluster_service#Host`,
		Resolver:    fetchOpensearchHosts,
		Transform:   structNameClusterIdTransformer(&opensearch.Host{}),
	}
}

func fetchOpensearchHosts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	cluster, ok := parent.Item.(*opensearch.Cluster)
	if !ok {
		return fmt.Errorf("parent is not type of *opensearch.Cluster: %+v", cluster)
	}

	it := c.SDK.MDB().OpenSearch().Cluster().ClusterHostsIterator(ctx, &opensearch.ListClusterHostsRequest{ClusterId: cluster.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
