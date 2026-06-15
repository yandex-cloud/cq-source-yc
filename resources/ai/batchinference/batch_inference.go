package batchinference

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/cq-source-yc/client/yc"
	batchinferencepb "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/batch_inference/v1"
)

func BatchInferences() *schema.Table {
	return &schema.Table{
		Name:        "yc_ai_batch_inference_tasks",
		Description: `https://yandex.cloud/docs/foundation-models/batch/api-ref/grpc/BatchInference/list#yandex.cloud.ai.batch_inference.v1.BatchInferenceTask`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchBatchInferences,
		Transform:   client.TransformWithStruct(&batchinferencepb.BatchInferenceTask{}, transformers.WithPrimaryKeys("TaskId")),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchBatchInferences(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	cc, err := c.GRPCConn(ctx, "yandex.cloud.ai.batch_inference.v1.BatchInferenceService.List")
	if err != nil {
		return err
	}
	cl := batchinferencepb.NewBatchInferenceServiceClient(cc)

	return yc.Paginate(ctx, res,
		func(t string) *batchinferencepb.ListBatchInferencesRequest {
			return &batchinferencepb.ListBatchInferencesRequest{FolderId: c.FolderId, PageToken: t}
		},
		cl.List,
		(*batchinferencepb.ListBatchInferencesResponse).GetTasks,
		(*batchinferencepb.ListBatchInferencesResponse).GetNextPageToken,
	)
}
