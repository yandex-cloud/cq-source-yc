package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/cq-provider-yandex/tools"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
)

func IamServiceAccounts() *schema.Table {
	gen, err := tools.NewTableGenerator(
		"Iam",
		"ServiceAccount",
		tools.WithProtoFile("resources/proto/service_account.proto"),
		tools.WithFetcher(fetchIamServiceAccounts),
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

func fetchIamServiceAccounts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	locations := []string{c.FolderId}

	for _, f := range locations {
		req := &iam.ListServiceAccountsRequest{FolderId: f}
		it := c.Services.Iam.ServiceAccount().ServiceAccountIterator(ctx, req)
		for it.Next() {
			res <- it.Value()
		}
	}
	return nil
}
