package access

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/lockbox/v1"
)

func SecretsAccessBindings() *schema.Table {
	return &schema.Table{
		Name:        "yc_access_bindings_lockbox_secrets",
		Description: `https://cloud.yandex.ru/docs/lockbox/api-ref/grpc/secret_service#AccessBinding`,
		Resolver:    fetchSecretsAccessBindings,
		Transform:   Transform,
		Columns: schema.ColumnList{
			client.ParentIdColumn,
		},
	}
}

func fetchSecretsAccessBindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	secret, ok := parent.Item.(*lockbox.Secret)
	if !ok {
		return fmt.Errorf("parent is not type of *lockbox.Secret: %+v", secret)
	}

	it := c.SDK.LockboxSecret().Secret().SecretAccessBindingsIterator(ctx, &access.ListAccessBindingsRequest{ResourceId: secret.Id})
	for it.Next() {
		value := it.Value()
		res <- value
	}

	return it.Error()
}
