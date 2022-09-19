package recipies

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

var resourceIDColumns = codegen.ColumnDefinitions{
	{
		Name:     "resource_id",
		Type:     schema.TypeString,
		Resolver: "client.ResolveMultiplexedResourceID",
		Options:  schema.ColumnCreationOptions{PrimaryKey: true},
	},
}
