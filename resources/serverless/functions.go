package serverless

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/cq-source-yc/resources/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/functions/v1"
)

func FunctionsFunctions() *schema.Table {
	return &schema.Table{
		Name:        "yc_serverless_functions_functions",
		Description: ``,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchFunctionsFunctions,
		Transform:   client.TransformWithStruct(&functions.Function{}, client.PrimaryKeyIdTransformer),
		Relations:   schema.Tables{access.ServerlessFunctionsAccessBindings()},
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}
func fetchFunctionsFunctions(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	folderId := c.FolderId

	it := c.SDK.Serverless().Functions().Function().FunctionIterator(ctx, &functions.ListFunctionsRequest{FolderId: folderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
