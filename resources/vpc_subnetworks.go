package resources

import (
	"context"

	"github.com/yandex-cloud/cq-provider-yandex/tools"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func VpcSubnetworks() *schema.Table {
	gen, err := tools.NewTableGenerator(
		"Vpc",
		"Subnet",
		tools.WithProtoFile("yandex/cloud/vpc/v1/subnet.proto", "cloudapi"),
		tools.WithFetcher(fetchVpcSubnetworks),
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

func fetchVpcSubnetworks(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	locations := []string{c.FolderId}

	for _, f := range locations {
		req := &vpc.ListSubnetsRequest{FolderId: f}
		it := c.Services.Vpc.Subnet().SubnetIterator(ctx, req)
		for it.Next() {
			res <- it.Value()
		}
	}
	return nil
}
