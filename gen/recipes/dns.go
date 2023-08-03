package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/dns/v1"
)

func DNS() []*Resource {
	return []*Resource{
		{
			Service:      "dns",
			SubService:   "zones",
			Struct:       new(dns.DnsZone),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    multiplexFolder,
			Relations:    []string{"Records()"},
		},
		{
			Service:    "dns",
			SubService: "records",
			Struct:     new(dns.RecordSet),
			SkipFields: []string{"Name", "Type"},
			ExtraColumns: codegen.ColumnDefinitions{
				{
					Name:     "zone_id",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("id")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "name",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("Name")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "type",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("Type")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
	}
}
