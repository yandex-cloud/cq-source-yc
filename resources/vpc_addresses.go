package resources

import (
	"context"

	"github.com/yandex-cloud/cq-provider-yandex/tools"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
)

func VpcAddresses() *schema.Table {
	gen, err := tools.NewTableGenerator(
		"Vpc",
		"Address",
		tools.WithProtoFile("yandex/cloud/vpc/v1/address.proto", "cloudapi"),
		tools.WithIgnoredColumns("Type", "IpVersion"),
		tools.WithFetcher(fetchVpcAddresses),
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

func fetchVpcAddresses(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	locations := []string{c.FolderId}

	for _, f := range locations {
		req := &vpc.ListAddressesRequest{FolderId: f}
		it := c.Services.Vpc.Address().AddressIterator(ctx, req)
		for it.Next() {
			res <- it.Value()
		}
	}
	return nil
}
