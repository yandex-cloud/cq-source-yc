package recipies

import (
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
)

func AccessBindings() []*Resource {
	unwrapSubject := []string{"Subject"}

	return []*Resource{
		{
			Service:        "access_bindings",
			SubService:     "by_organization",
			Struct:         new(access.AccessBinding),
			ExtraColumns:   resourceIDColumns,
			FieldsToUnwrap: unwrapSubject,
			Multiplex:      "client.MultiplexBy(client.Organizations)",
		},
		{
			Service:        "access_bindings",
			SubService:     "by_cloud",
			Struct:         new(access.AccessBinding),
			ExtraColumns:   resourceIDColumns,
			FieldsToUnwrap: unwrapSubject,
			Multiplex:      "client.MultiplexBy(client.Clouds)",
		},
		{
			Service:        "access_bindings",
			SubService:     "by_folder",
			Struct:         new(access.AccessBinding),
			ExtraColumns:   resourceIDColumns,
			FieldsToUnwrap: unwrapSubject,
			Multiplex:      "client.MultiplexBy(client.Folders)",
		},
	}
}
