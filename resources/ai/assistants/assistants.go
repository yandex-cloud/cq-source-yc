package assistants

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/cq-source-yc/client/yc"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/ai/assistants/v1"
)

func Assistants() *schema.Table {
	return &schema.Table{
		Name:        "yc_ai_assistants",
		Description: `https://yandex.cloud/docs/foundation-models/assistants/api-ref/grpc/Assistant/list#yandex.cloud.ai.assistants.v1.Assistant`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchAssistants,
		Transform:   client.TransformWithStruct(&assistants.Assistant{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchAssistants(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	cc, err := c.GRPCConn(ctx, "yandex.cloud.ai.assistants.v1.AssistantService.List")
	if err != nil {
		return err
	}
	cl := assistants.NewAssistantServiceClient(cc)

	return yc.Paginate(ctx, res,
		func(t string) *assistants.ListAssistantsRequest {
			return &assistants.ListAssistantsRequest{FolderId: c.FolderId, PageToken: t}
		},
		cl.List,
		(*assistants.ListAssistantsResponse).GetAssistants,
		(*assistants.ListAssistantsResponse).GetNextPageToken,
	)
}
