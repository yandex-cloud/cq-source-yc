package datasets

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/cq-source-yc/client/yc"
	dataset "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/dataset/v1"
)

func Datasets() *schema.Table {
	return &schema.Table{
		Name:        "yc_ai_datasets",
		Description: `https://yandex.cloud/docs/foundation-models/dataset/api-ref/grpc/Dataset/list#yandex.cloud.ai.dataset.v1.DatasetInfo`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchDatasets,
		Transform:   client.TransformWithStruct(&dataset.DatasetInfo{}, transformers.WithPrimaryKeys("DatasetId")),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchDatasets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	cc, err := c.GRPCConn(ctx, "yandex.cloud.ai.dataset.v1.DatasetService.List")
	if err != nil {
		return err
	}
	cl := dataset.NewDatasetServiceClient(cc)

	return yc.Paginate(ctx, res,
		func(t string) *dataset.ListDatasetsRequest {
			return &dataset.ListDatasetsRequest{FolderId: c.FolderId, PageToken: t}
		},
		cl.List,
		(*dataset.ListDatasetsResponse).GetDatasets,
		(*dataset.ListDatasetsResponse).GetNextPageToken,
	)
}
