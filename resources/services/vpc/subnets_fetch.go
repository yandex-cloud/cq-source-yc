package vpc

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
)

func fetchSubnets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.Services.VPC.Subnet().SubnetIterator(ctx, &vpc.ListSubnetsRequest{FolderId: c.MultiplexedResourceId})
	for it.Next() {
		res <- it.Value()
	}

	return nil
}
