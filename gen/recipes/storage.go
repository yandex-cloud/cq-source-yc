package recipes

import (
	storage_resource "github.com/yandex-cloud/cq-provider-yandex/resources/services/storage"
)

func Storage() []*Resource {
	return []*Resource{
		{
			Service:        "storage",
			SubService:     "buckets",
			Struct:         new(storage_resource.Bucket),
			SkipFields:     []string{id}, // Id is always ""
			FieldsToUnwrap: []string{"Bucket"},
			Multiplex:      multiplexFolder,
		},
	}
}
