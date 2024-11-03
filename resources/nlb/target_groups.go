package nlb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/loadbalancer/v1"
)

func TargetGroups() *schema.Table {
	return &schema.Table{
		Name:        "yc_nlb_target_groups",
		Description: `https://yandex.cloud/ru/docs/network-load-balancer/api-ref/grpc/TargetGroup/list#yandex.cloud.loadbalancer.v1.TargetGroup`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchTargetGroups,
		Transform:   client.TransformWithStruct(&loadbalancer.TargetGroup{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchTargetGroups(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.LoadBalancer().TargetGroup().TargetGroupIterator(ctx, &loadbalancer.ListTargetGroupsRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
