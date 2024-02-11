package iam

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
)

func ApiKeys() *schema.Table {
	return &schema.Table{
		Name:        "yc_iam_api_keys",
		Description: `https://cloud.yandex.ru/docs/iam/api-ref/grpc/api_key_service#ApiKey`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchApiKeys,
		Transform:   client.TransformWithStruct(&iam.ApiKey{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchApiKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	sa, ok := parent.Item.(*iam.ServiceAccount)
	if !ok {
		return fmt.Errorf("parent is not type of *iam.ServiceAccount: %+v", sa)
	}

	it := c.SDK.IAM().ApiKey().ApiKeyIterator(ctx, &iam.ListApiKeysRequest{ServiceAccountId: sa.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
