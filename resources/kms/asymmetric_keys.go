package kms

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	caccess "github.com/yandex-cloud/cq-source-yc/resources/access"
	kms "github.com/yandex-cloud/go-genproto/yandex/cloud/kms/v1/asymmetricencryption"
)

func AsymmetricKeys() *schema.Table {
	return &schema.Table{
		Name:        "yc_kms_asymmetric_keys",
		Description: `https://yandex.cloud/ru/docs/kms/api-ref/grpc/asymmetric_encryption_key_service#AsymmetricEncryptionKey2`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchAsymmetricKeys,
		Transform:   client.TransformWithStruct(&kms.AsymmetricEncryptionKey{}, client.PrimaryKeyIdTransformer),
		Relations:   schema.Tables{caccess.AsymmetricKeysAccessBindings()},
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchAsymmetricKeys(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.KMSAsymmetricEncryption().AsymmetricEncryptionKey().AsymmetricEncryptionKeyIterator(ctx, &kms.ListAsymmetricEncryptionKeysRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
