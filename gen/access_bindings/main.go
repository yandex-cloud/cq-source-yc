package main

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/yandex-cloud/cq-provider-yandex/gen/util"
)

func generate(resource, manager string) {
	util.SilentExecute(util.TemplatesDir{
		MainFile: "access_bindings_by_resource.go.tmpl",
		Path:     "templates",
	}, map[string]string{
		"resource": resource,
		"manager":  manager,
	}, fmt.Sprintf("resources1/access_bindings_by_%s.go", strcase.ToSnake(resource)))

	util.SilentExecute(util.TemplatesDir{
		MainFile: "iam_user_accounts_by_resource.go.tmpl",
		Path:     "templates",
	}, map[string]string{
		"resource": resource,
		"manager":  manager,
	}, fmt.Sprintf("resources1/iam_user_accounts_by_%s.go", strcase.ToSnake(resource)))
}

func main() {
	generate("Organization", "OrganizationManager")
	generate("Cloud", "ResourceManager")
	generate("Folder", "ResourceManager")
}
