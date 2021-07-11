package resources

import (
	"context"
	"github.com/yandex-cloud/cq-provider-yandex/tools"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func VpcNetworks() *schema.Table {
	gen, err := tools.NewTableGenerator(
		"Vpc",
		"Network",
		tools.WithProtoFile("resources/proto/network.proto"),
		tools.WithFetcher(fetchVpcNetworks),
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

func fetchVpcNetworks(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	locations := []string{c.FolderId}

	for _, f := range locations {
		req := &vpc.ListNetworksRequest{FolderId: f}
		it := c.Services.Vpc.Network().NetworkIterator(ctx, req)
		for it.Next() {
			res <- it.Value()
		}
	}
	return nil
}
