package recipies

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
)

func AccessBindings() []*Resource {
	commonColumns := codegen.ColumnDefinitions{
		{
			Name:     "role_id",
			Type:     schema.TypeString,
			Resolver: `schema.PathResolver("RoleId")`,
			Options:  schema.ColumnCreationOptions{PrimaryKey: true},
		},
		{
			Name:     "subject_id",
			Type:     schema.TypeString,
			Resolver: `schema.PathResolver("Subject.Id")`,
			Options:  schema.ColumnCreationOptions{PrimaryKey: true},
		},
		{
			Name:     "subject_type",
			Type:     schema.TypeString,
			Resolver: `schema.PathResolver("Subject.Type")`,
		},
	}
	commonSkip := []string{"RoleId", "Subject"}
	return []*Resource{
		{
			Service:    "access_bindings",
			SubService: "by_organization",
			Struct:     new(access.AccessBinding),
			SkipFields: commonSkip,
			ExtraColumns: append(codegen.ColumnDefinitions{
				{
					Name:     "organization_id",
					Type:     schema.TypeString,
					Resolver: "client.ResolveMultiplexedResourceID",
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}, commonColumns...),
			Multiplex: multiplexOrg,
		},
		{
			Service:    "access_bindings",
			SubService: "by_cloud",
			Struct:     new(access.AccessBinding),
			SkipFields: commonSkip,
			ExtraColumns: append(codegen.ColumnDefinitions{
				{
					Name:     "cloud_id",
					Type:     schema.TypeString,
					Resolver: "client.ResolveMultiplexedResourceID",
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}, commonColumns...),
			Multiplex: multiplexCloud,
		},
		{
			Service:    "access_bindings",
			SubService: "by_folder",
			Struct:     new(access.AccessBinding),
			SkipFields: commonSkip,
			ExtraColumns: append(codegen.ColumnDefinitions{
				{
					Name:     "folder_id",
					Type:     schema.TypeString,
					Resolver: "client.ResolveMultiplexedResourceID",
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}, commonColumns...),
			Multiplex: multiplexFolder,
		},
	}
}
