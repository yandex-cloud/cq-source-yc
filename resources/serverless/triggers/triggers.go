package triggers

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/triggers/v1"
)

func Triggers() *schema.Table {
	return &schema.Table{
		Name:        "yc_serverless_triggers",
		Description: `https://yandex.cloud/docs/functions/triggers/api-ref/grpc/Trigger/list#yandex.cloud.serverless.triggers.v1.Trigger`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchTriggers,
		Transform:   client.TransformWithStruct(&triggers.Trigger{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchTriggers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	it := c.SDK.Serverless().Triggers().Trigger().TriggerIterator(ctx, &triggers.ListTriggersRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
