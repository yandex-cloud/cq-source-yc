package cloudrouter

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/cloudrouter/v1"
	crsdk "github.com/yandex-cloud/go-sdk/services/cloudrouter/v1"
)

func RoutingInstances() *schema.Table {
	return &schema.Table{
		Name:        "yc_cloudrouter_routing_instances",
		Description: `https://yandex.cloud/ru/docs/cloud-router/api-ref/grpc/RoutingInstance/list#yandex.cloud.cloudrouter.v1.RoutingInstance`,
		Resolver:    nil,
		Transform:   client.TransformWithStruct(&cloudrouter.RoutingInstance{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchRoutingInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := crsdk.NewRoutingInstanceClient(c.SDKv2).Iterator(ctx, &cloudrouter.ListRoutingInstancesRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
