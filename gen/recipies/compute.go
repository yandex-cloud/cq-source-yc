package recipies

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
)

func Compute() []*Resource {
	return []*Resource{
		{
			Service:      "compute",
			SubService:   "disks",
			Struct:       new(compute.Disk),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    multiplexFolder,
		},
		{
			Service:      "compute",
			SubService:   "images",
			Struct:       new(compute.Image),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    multiplexFolder,
		},
		{
			Service:      "compute",
			SubService:   "instances",
			Struct:       new(compute.Instance),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    multiplexFolder,
		},
	}
}
