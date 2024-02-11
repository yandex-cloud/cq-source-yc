package nlb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/loadbalancer/v1"
)

func Balancers() *schema.Table {
	return &schema.Table{
		Name:        "yc_nlb_balancers",
		Description: `https://cloud.yandex.ru/docs/network-load-balancer/api-ref/grpc/network_load_balancer_service#NetworkLoadBalancer1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchBalancers,
		Transform:   client.TransformWithStruct(&loadbalancer.NetworkLoadBalancer{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchBalancers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	folderId := c.FolderId

	it := c.SDK.LoadBalancer().NetworkLoadBalancer().NetworkLoadBalancerIterator(ctx, &loadbalancer.ListNetworkLoadBalancersRequest{FolderId: folderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
