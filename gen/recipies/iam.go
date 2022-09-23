package recipies

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	iam_resource "github.com/yandex-cloud/cq-provider-yandex/resources/services/iam"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
)

func IAM() []*Resource {
	return []*Resource{
		{
			Service:      "iam",
			SubService:   "service_accounts",
			Struct:       new(iam.ServiceAccount),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    multiplexFolder,
		},
		{
			Service:    "iam",
			SubService: "user_accounts_by_organization",
			Struct:     new(iam_resource.UserAccount),
			SkipFields: []string{id},
			ExtraColumns: codegen.ColumnDefinitions{
				idCol,
				{
					Name:     "organization_id",
					Type:     schema.TypeString,
					Resolver: "client.ResolveMultiplexedResourceID",
				},
			},
			Multiplex:      multiplexOrg,
			FieldsToUnwrap: []string{"UserAccount"},
		},
		{
			Service:    "iam",
			SubService: "user_accounts_by_cloud",
			Struct:     new(iam_resource.UserAccount),
			SkipFields: []string{id},
			ExtraColumns: codegen.ColumnDefinitions{
				idCol,
				{
					Name:     "cloud_id",
					Type:     schema.TypeString,
					Resolver: "client.ResolveMultiplexedResourceID",
				},
			},
			Multiplex:      multiplexCloud,
			FieldsToUnwrap: []string{"UserAccount"},
		},
		{
			Service:    "iam",
			SubService: "user_accounts_by_folder",
			Struct:     new(iam_resource.UserAccount),
			SkipFields: []string{id},
			ExtraColumns: codegen.ColumnDefinitions{
				idCol,
				{
					Name:     "folder_id",
					Type:     schema.TypeString,
					Resolver: "client.ResolveMultiplexedResourceID",
				},
			},
			Multiplex:      multiplexFolder,
			FieldsToUnwrap: []string{"UserAccount"},
		},
	}
}
