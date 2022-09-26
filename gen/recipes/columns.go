package recipes

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
)
