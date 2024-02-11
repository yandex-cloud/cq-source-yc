package iam

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/cq-source-yc/resources/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
)

func ServiceAccounts() *schema.Table {
	return &schema.Table{
		Name:        "yc_iam_service_accounts",
		Description: `https://cloud.yandex.ru/docs/iam/api-ref/grpc/service_account_service#ServiceAccount1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchServiceAccounts,
		Transform:   client.TransformWithStruct(&iam.ServiceAccount{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
		Relations: schema.Tables{
			AccessKeys(),
			ApiKeys(),
			Keys(),
			access.ServiceAccountsAccessBindings(),
		},
	}
}

func fetchServiceAccounts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.IAM().ServiceAccount().ServiceAccountIterator(ctx, &iam.ListServiceAccountsRequest{FolderId: c.FolderId})
	for it.Next() {
		value := it.Value()
		res <- value
	}

	return it.Error()
}
