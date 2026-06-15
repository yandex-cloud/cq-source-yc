package eventrouter

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/eventrouter/v1"
)

func Rules() *schema.Table {
	return &schema.Table{
		Name:        "yc_serverless_eventrouter_rules",
		Description: `https://yandex.cloud/docs/serverless-integrations/eventrouter/api-ref/grpc/Rule/list#yandex.cloud.serverless.eventrouter.v1.Rule`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchRules,
		Transform:   client.TransformWithStruct(&eventrouter.Rule{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchRules(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	req := &eventrouter.ListRulesRequest{
		ContainerId: &eventrouter.ListRulesRequest_FolderId{FolderId: c.FolderId},
	}
	it := c.SDK.Serverless().Eventrouter().Rule().RuleIterator(ctx, req)
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
