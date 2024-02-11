package access

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
)

func ServiceAccountsAccessBindings() *schema.Table {
	return &schema.Table{
		Name:        "yc_access_iam_service_accounts",
		Description: `https://cloud.yandex.ru/docs/iam/api-ref/grpc/service_account_service#AccessBinding`,
		Resolver:    fetchServiceAccountsAccessBindings,
		Transform:   Transform,
		Columns: schema.ColumnList{
			client.ParentIdColumn,
		},
	}
}

func fetchServiceAccountsAccessBindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	sa, ok := parent.Item.(*iam.ServiceAccount)
	if !ok {
		return fmt.Errorf("parent is not type of *iam.ServiceAccount: %+v", sa)
	}

	it := c.SDK.IAM().ServiceAccount().ServiceAccountAccessBindingsIterator(ctx, &access.ListAccessBindingsRequest{ResourceId: sa.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
