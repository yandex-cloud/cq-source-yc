package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	storage_resource "github.com/yandex-cloud/cq-provider-yandex/resources/services/storage"
)

func Storage() []*Resource {
	return []*Resource{
		{
			Service:        "storage",
			SubService:     "buckets",
			Struct:         new(storage_resource.Bucket),
			SkipFields:     []string{id, "Name", "FolderId"}, // Id is always ""
			FieldsToUnwrap: []string{"Bucket"},
			Multiplex:      multiplexFolder,
			ExtraColumns: codegen.ColumnDefinitions{
				{
					Name:     "name",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("Name")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "folder_id",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("FolderId")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
	}
}
