package recipies

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

const (
	id = "Id"
)

var (
	idCol = codegen.ColumnDefinition{
		Name:        "id",
		Type:        schema.TypeString,
		Resolver:    `schema.PathResolver("Id")`,
		Description: `Resource ID`,
		Options:     schema.ColumnCreationOptions{PrimaryKey: true},
	}

	folderIDCol = codegen.ColumnDefinition{
		Name:        "folder_id",
		Type:        schema.TypeString,
		Resolver:    `client.ResolveMultiplexedResourceID`,
		Description: `Folder ID`,
	}
)
