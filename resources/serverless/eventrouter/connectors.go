package eventrouter

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/eventrouter/v1"
)

func Connectors() *schema.Table {
	return &schema.Table{
		Name:        "yc_serverless_eventrouter_connectors",
		Description: `https://yandex.cloud/docs/serverless-integrations/eventrouter/api-ref/grpc/Connector/list#yandex.cloud.serverless.eventrouter.v1.Connector`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchConnectors,
		Transform:   client.TransformWithStruct(&eventrouter.Connector{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchConnectors(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	req := &eventrouter.ListConnectorsRequest{
		ContainerId: &eventrouter.ListConnectorsRequest_FolderId{FolderId: c.FolderId},
	}
	it := c.SDK.Serverless().Eventrouter().Connector().ConnectorIterator(ctx, req)
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
