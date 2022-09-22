package recipies

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
)

func AccessBindings() []*Resource {
	unwrapSubject := []string{"Subject"}

	return []*Resource{
		{
			Service:    "access_bindings",
			SubService: "by_organization",
			Struct:     new(access.AccessBinding),
			ExtraColumns: codegen.ColumnDefinitions{
				{
					Name:     "organization_id",
					Type:     schema.TypeString,
					Resolver: "client.ResolveMultiplexedResourceID",
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
			FieldsToUnwrap: unwrapSubject,
			Multiplex:      multiplexOrg,
		},
		{
			Service:    "access_bindings",
			SubService: "by_cloud",
			Struct:     new(access.AccessBinding),
			ExtraColumns: codegen.ColumnDefinitions{
				{
					Name:     "cloud_id",
					Type:     schema.TypeString,
					Resolver: "client.ResolveMultiplexedResourceID",
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
			FieldsToUnwrap: unwrapSubject,
			Multiplex:      multiplexCloud,
		},
		{
			Service:    "access_bindings",
			SubService: "by_folder",
			Struct:     new(access.AccessBinding),
			ExtraColumns: codegen.ColumnDefinitions{
				{
					Name:     "folder_id",
					Type:     schema.TypeString,
					Resolver: "client.ResolveMultiplexedResourceID",
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
			FieldsToUnwrap: unwrapSubject,
			Multiplex:      multiplexFolder,
		},
	}
}
