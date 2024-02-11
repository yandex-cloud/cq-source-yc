package lockbox

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/cq-source-yc/resources/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/lockbox/v1"
)

func Secrets() *schema.Table {
	return &schema.Table{
		Name:        "yc_lockbox_secrets",
		Description: `https://cloud.yandex.ru/docs/lockbox/api-ref/grpc/secret_service#Secret1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchSecrets,
		Transform:   client.TransformWithStruct(&lockbox.Secret{}, client.PrimaryKeyIdTransformer),
		Relations:   schema.Tables{access.SecretsAccessBindings()},
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchSecrets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.LockboxSecret().Secret().SecretIterator(ctx, &lockbox.ListSecretsRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
