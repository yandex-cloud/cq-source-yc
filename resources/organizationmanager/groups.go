package organizationmanager

import "github.com/cloudquery/plugin-sdk/v4/schema"

func Groups() *schema.Table {
	return &schema.Table{
		Name: "yc_organizationmanager_groups",
	}
}
