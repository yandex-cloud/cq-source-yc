package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/lockbox/v1"
)

func Lockbox() []*Resource {
	return []*Resource{
		{
			Service:      "lockbox",
			SubService:   "secrets",
			Struct:       new(lockbox.Secret),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    multiplexFolder,
		},
	}
}
