package datatransfer

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/datatransfer/v1"
)

func Transfers() *schema.Table {
	return &schema.Table{
		Name:        "yc_datatransfer_transfers",
		Description: `https://yandex.cloud/ru/docs/data-transfer/api-ref/grpc/transfer_service#Transfer`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchTransfers,
		Transform:   client.TransformWithStruct(&datatransfer.Transfer{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchTransfers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.DataTransfer().Transfer().TransferIterator(ctx, &datatransfer.ListTransfersRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
