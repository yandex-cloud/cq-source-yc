package recipies

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	storage_resource "github.com/yandex-cloud/cq-provider-yandex/resources/services/storage"
)

func Storage() []*Resource {
	return []*Resource{
		{
			Service:        "storage",
			SubService:     "buckets",
			Struct:         new(storage_resource.Bucket),
			SkipFields:     []string{id},
			ExtraColumns:   codegen.ColumnDefinitions{idCol},
			FieldsToUnwrap: []string{"Bucket"},
			Multiplex:      multiplexFolder,
		},
	}
}
