package eventrouter

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/eventrouter/v1"
)

func Buses() *schema.Table {
	return &schema.Table{
		Name:        "yc_serverless_eventrouter_buses",
		Description: `https://yandex.cloud/docs/serverless-integrations/eventrouter/api-ref/grpc/Bus/list#yandex.cloud.serverless.eventrouter.v1.Bus`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchBuses,
		Transform:   client.TransformWithStruct(&eventrouter.Bus{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchBuses(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	it := c.SDK.Serverless().Eventrouter().Bus().BusIterator(ctx, &eventrouter.ListBusesRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
