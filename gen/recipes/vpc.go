package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
)

func VPC() []*Resource {
	return []*Resource{
		{
			Service:      "vpc",
			SubService:   "addresses",
			Struct:       new(vpc.Address),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    multiplexFolder,
		},
		{
			Service:      "vpc",
			SubService:   "networks",
			Struct:       new(vpc.Network),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    multiplexFolder,
		},
		{
			Service:      "vpc",
			SubService:   "security_groups",
			Struct:       new(vpc.SecurityGroup),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    multiplexFolder,
		},
		{
			Service:      "vpc",
			SubService:   "subnets",
			Struct:       new(vpc.Subnet),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    multiplexFolder,
		},
	}
}
