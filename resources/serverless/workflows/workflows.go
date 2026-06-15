package workflows

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/workflows/v1"
)

func Workflows() *schema.Table {
	return &schema.Table{
		Name:        "yc_serverless_workflows",
		Description: `https://yandex.cloud/docs/serverless-integrations/workflows/api-ref/grpc/Workflow/list#yandex.cloud.serverless.workflows.v1.Workflow`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchWorkflows,
		Transform:   client.TransformWithStruct(&workflows.Workflow{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchWorkflows(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	it := c.SDK.Serverless().Workflow().Workflow().WorkflowIterator(ctx, &workflows.ListWorkflowsRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
