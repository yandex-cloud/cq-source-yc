package recipies

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
)

func ResourceManager() []*Resource {
	return []*Resource{
		{
			Service:      "resourcemanager",
			SubService:   "clouds",
			Struct:       new(resourcemanager.Cloud),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    multiplexCloud,
			Relations:    []string{"Folders()"},
		},
		{
			Service:      "resourcemanager",
			SubService:   "folders",
			Struct:       new(resourcemanager.Folder),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    multiplexFolder,
		},
	}
}
