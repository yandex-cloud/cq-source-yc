package dns

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/dns/v1"
)

func fetchZones(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &dns.ListDnsZonesRequest{FolderId: c.MultiplexedResourceId}
	it := c.Services.DNS.DnsZone().DnsZoneIterator(ctx, req)
	for it.Next() {
		res <- it.Value()

	}

	return nil
}
