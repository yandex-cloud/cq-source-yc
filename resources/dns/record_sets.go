package dns

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/dns/v1"
)

func RecordSets() *schema.Table {
	return &schema.Table{
		Name:        "yc_dns_record_sets",
		Description: `https://cloud.yandex.ru/docs/dns/api-ref/grpc/dns_zone_service#RecordSet1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchRecordSets,
		Transform:   client.TransformWithStruct(&dns.RecordSet{}, transformers.WithPrimaryKeys("Name", "Type")),
		Columns: schema.ColumnList{
			{
				Name:       "zone_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("id"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchRecordSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	zone, ok := parent.Item.(*dns.DnsZone)
	if !ok {
		return fmt.Errorf("parent is not type of *dns.DnsZone: %+v", zone)
	}

	it := c.SDK.DNS().DnsZone().DnsZoneRecordSetsIterator(ctx, &dns.ListDnsZoneRecordSetsRequest{DnsZoneId: zone.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
