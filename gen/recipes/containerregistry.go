package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/containerregistry/v1"
)

func ContainerRegistry() []*Resource {
	var folderIDCol = codegen.ColumnDefinition{
		Name:        "folder_id",
		Type:        schema.TypeString,
		Resolver:    `client.ResolveMultiplexedResourceID`,
		Description: `Folder ID`,
	}

	return []*Resource{
		{
			Service:      "containerregistry",
			SubService:   "images",
			Struct:       new(containerregistry.Image),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol, folderIDCol},
			Multiplex:    multiplexFolder,
		},
		{
			Service:      "containerregistry",
			SubService:   "registries",
			Struct:       new(containerregistry.Registry),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    multiplexFolder,
		},
		{
			Service:      "containerregistry",
			SubService:   "scan_results",
			Struct:       new(containerregistry.ScanResult),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol, folderIDCol},
			Multiplex:    multiplexFolder,
		},
	}
}
