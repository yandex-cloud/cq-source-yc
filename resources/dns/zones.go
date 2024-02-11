package dns

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/dns/v1"
)

func Zones() *schema.Table {
	return &schema.Table{
		Name:        "yc_dns_zones",
		Description: `https://cloud.yandex.ru/docs/dns/api-ref/grpc/dns_zone_service#DnsZone1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchZones,
		Transform:   client.TransformWithStruct(&dns.DnsZone{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
		Relations: schema.Tables{RecordSets()},
	}
}

func fetchZones(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.DNS().DnsZone().DnsZoneIterator(ctx, &dns.ListDnsZonesRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
