package k8s

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/k8s/v1"
)

func fetchClusters(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	req := &k8s.ListClustersRequest{FolderId: c.MultiplexedResourceId}
	it := c.Services.K8S.Cluster().ClusterIterator(ctx, req)
	for it.Next() {
		value := it.Value()
		cluster := Cluster{Cluster: value}
		if value.GetNetworkImplementation() != nil {
			cluster.NetworkImplementation = &k8s.Cluster_Cilium{Cilium: value.GetCilium()}
		}
		if value.GetInternetGateway() != nil {
			cluster.InternetGateway = &k8s.Cluster_GatewayIpv4Address{GatewayIpv4Address: cluster.GetGatewayIpv4Address()}
		}
		res <- cluster
	}

	return nil
}
