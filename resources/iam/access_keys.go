package iam

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1/awscompatibility"
)

func AccessKeys() *schema.Table {
	return &schema.Table{
		Name:        "yc_iam_access_keys",
		Description: `https://cloud.yandex.ru/docs/iam/api-ref/grpc/access_key_service#AccessKey`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchAccessKeys,
		Transform:   client.TransformWithStruct(&awscompatibility.AccessKey{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchAccessKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	sa, ok := parent.Item.(*iam.ServiceAccount)
	if !ok {
		return fmt.Errorf("parent is not type of *iam.ServiceAccount: %+v", sa)
	}

	it := c.SDK.IAM().AWSCompatibility().AccessKey().AccessKeyIterator(ctx, &awscompatibility.ListAccessKeysRequest{ServiceAccountId: sa.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
