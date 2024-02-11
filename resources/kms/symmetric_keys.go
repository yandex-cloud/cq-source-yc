package kms

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	caccess "github.com/yandex-cloud/cq-source-yc/resources/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/kms/v1"
)

func SymmetricKeys() *schema.Table {
	return &schema.Table{
		Name:        "yc_kms_symmetric_keys",
		Description: ``,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchSymmetricKeys,
		Transform:   client.TransformWithStruct(&kms.SymmetricKey{}, client.PrimaryKeyIdTransformer),
		Relations:   schema.Tables{caccess.SymmetricKeysAccessBindings()},
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchSymmetricKeys(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.KMS().SymmetricKey().SymmetricKeyIterator(ctx, &kms.ListSymmetricKeysRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
