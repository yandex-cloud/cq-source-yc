package access

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/functions/v1"
)

func ServerlessFunctionsAccessBindings() *schema.Table {
	return &schema.Table{
		Name:      "yc_access_bindings_serverless_functions",
		Resolver:  fetchServerlessFunctionsAccessBindings,
		Transform: Transform,
		Columns: schema.ColumnList{
			client.ParentIdColumn,
		},
	}
}

func fetchServerlessFunctionsAccessBindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	function, ok := parent.Item.(*functions.Function)
	if !ok {
		return fmt.Errorf("parent is not type of *functions.Function: %+v", function)
	}

	it := c.SDK.Serverless().Functions().Function().FunctionAccessBindingsIterator(ctx, &access.ListAccessBindingsRequest{ResourceId: function.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
