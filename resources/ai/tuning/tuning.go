package tuning

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/cq-source-yc/client/yc"
	tuningpb "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/tuning/v1"
)

func Tunings() *schema.Table {
	return &schema.Table{
		Name:        "yc_ai_tuning_tasks",
		Description: `https://yandex.cloud/docs/foundation-models/tuning/api-ref/grpc/Tuning/list#yandex.cloud.ai.tuning.v1.TuningTask`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchTunings,
		Transform:   client.TransformWithStruct(&tuningpb.TuningTask{}, transformers.WithPrimaryKeys("TaskId")),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchTunings(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	cc, err := c.GRPCConn(ctx, "yandex.cloud.ai.tuning.v1.TuningService.List")
	if err != nil {
		return err
	}
	cl := tuningpb.NewTuningServiceClient(cc)

	return yc.Paginate(ctx, res,
		func(t string) *tuningpb.ListTuningsRequest {
			return &tuningpb.ListTuningsRequest{FolderId: c.FolderId, PageToken: t}
		},
		cl.List,
		(*tuningpb.ListTuningsResponse).GetTuningTasks,
		(*tuningpb.ListTuningsResponse).GetNextPageToken,
	)
}
