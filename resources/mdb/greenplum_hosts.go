package mdb

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/greenplum/v1"
)

func GreenplumHosts() *schema.Table {
	return &schema.Table{
		Name:        "yc_mdb_greenplum_hosts",
		Description: `https://cloud.yandex.ru/docs/managed-greenplum/api-ref/grpc/cluster_service#Host`,
		Resolver:    fetchGreenplumHosts,
		Transform:   structNameClusterIdTransformer(&greenplum.Host{}),
	}
}

func fetchGreenplumHosts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	cluster, ok := parent.Item.(*greenplum.Cluster)
	if !ok {
		return fmt.Errorf("parent is not type of *greenplum.Cluster: %+v", cluster)
	}

	itMaster := c.SDK.MDB().Greenplum().Cluster().ClusterMasterHostsIterator(ctx, &greenplum.ListClusterHostsRequest{ClusterId: cluster.Id})
	for itMaster.Next() {
		res <- itMaster.Value()
	}
	if itMaster.Error() != nil {
		return itMaster.Error()
	}

	itSegment := c.SDK.MDB().Greenplum().Cluster().ClusterSegmentHostsIterator(ctx, &greenplum.ListClusterHostsRequest{ClusterId: cluster.Id})
	for itSegment.Next() {
		res <- itSegment.Value()
	}
	return itSegment.Error()
}
