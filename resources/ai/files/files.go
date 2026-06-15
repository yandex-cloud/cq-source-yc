package files

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/cq-source-yc/client/yc"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/ai/files/v1"
)

func Files() *schema.Table {
	return &schema.Table{
		Name:        "yc_ai_files",
		Description: `https://yandex.cloud/docs/foundation-models/assistants/api-ref/grpc/Files/list#yandex.cloud.ai.files.v1.File`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchFiles,
		Transform:   client.TransformWithStruct(&files.File{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchFiles(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	cc, err := c.GRPCConn(ctx, "yandex.cloud.ai.files.v1.FileService.List")
	if err != nil {
		return err
	}
	cl := files.NewFileServiceClient(cc)

	return yc.Paginate(ctx, res,
		func(t string) *files.ListFilesRequest {
			return &files.ListFilesRequest{FolderId: c.FolderId, PageToken: t}
		},
		cl.List,
		(*files.ListFilesResponse).GetFiles,
		(*files.ListFilesResponse).GetNextPageToken,
	)
}
