package alb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/apploadbalancer/v1"
)

func Balancers() *schema.Table {
	return &schema.Table{
		Name:        "yc_alb_balancers",
		Description: `https://cloud.yandex.ru/docs/application-load-balancer/api-ref/grpc/load_balancer_service#LoadBalancer1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchBalancers,
		Transform:   client.TransformWithStruct(&apploadbalancer.LoadBalancer{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchBalancers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	folderId := c.FolderId

	it := c.SDK.ApplicationLoadBalancer().LoadBalancer().LoadBalancerIterator(ctx, &apploadbalancer.ListLoadBalancersRequest{FolderId: folderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
