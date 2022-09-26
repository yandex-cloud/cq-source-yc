package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1/saml"
)

func OrganizationManager() []*Resource {
	return []*Resource{
		{
			Service:      "organizationmanager",
			SubService:   "federations",
			Struct:       new(saml.Federation),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    multiplexOrg,
		},
		{
			Service:      "organizationmanager",
			SubService:   "organizations",
			Struct:       new(organizationmanager.Organization),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Relations:    []string{"Groups()"},
			Multiplex:    multiplexOrg,
		},
		{
			Service:      "organizationmanager",
			SubService:   "groups",
			Struct:       new(organizationmanager.Group),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    "", // Relation for organization
		},
	}
}
