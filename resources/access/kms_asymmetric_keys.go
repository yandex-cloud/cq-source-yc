package access

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	kms "github.com/yandex-cloud/go-genproto/yandex/cloud/kms/v1/asymmetricencryption"
)

func AsymmetricKeysAccessBindings() *schema.Table {
	return &schema.Table{
		Name:        "yc_access_bindings_kms_asymmetric_keys",
		Description: ``,
		Resolver:    fetchAsymmetricKeysAccessBindings,
		Transform:   Transform,
		Columns: schema.ColumnList{
			client.ParentIdColumn,
		},
	}
}

func fetchAsymmetricKeysAccessBindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	key, ok := parent.Item.(*kms.AsymmetricEncryptionKey)
	if !ok {
		return fmt.Errorf("parent is not type of *kms.SymmetricKey: %+v", key)
	}

	it := c.SDK.KMSAsymmetricEncryption().AsymmetricEncryptionKey().AsymmetricEncryptionKeyAccessBindingsIterator(ctx, &access.ListAccessBindingsRequest{ResourceId: key.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
