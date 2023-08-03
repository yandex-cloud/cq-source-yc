package dns

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/dns/v1"
)

func fetchRecords(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &dns.ListDnsZoneRecordSetsRequest{DnsZoneId: parent.Item.(*dns.DnsZone).Id}
	it := c.Services.DNS.DnsZone().DnsZoneRecordSetsIterator(ctx, req)

	for it.Next() {
		res <- it.Value()
	}

	return nil
}
