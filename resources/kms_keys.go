package resources

import (
	"context"
	"github.com/yandex-cloud/cq-provider-yandex/tools"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/kms/v1"
)

func KmsKeyring() *schema.Table {
	gen, err := tools.NewTableGenerator(
		"yandex_kms_symmetric_keys",
		"Kms",
		"SymmetricKey",
		"resources/proto/symmetric_key.proto",
		tools.GetCommonDefaultColumns("symmetric_key"),
		tools.IgnoredColumns{},
		fetchKmsSymmetricKeys,
	)
	if err != nil {
		return nil
	}
	table, err := gen.Generate()
	if err != nil {
		return nil
	}
	return table
}

func fetchKmsSymmetricKeys(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	locations := []string{c.FolderId}

	for _, f := range locations {
		req := &kms.ListSymmetricKeysRequest{FolderId: f}
		it := c.Services.Kms.SymmetricKey().SymmetricKeyIterator(ctx, req)
		for it.Next() {
			res <- it.Value()
		}
	}
	return nil
}
