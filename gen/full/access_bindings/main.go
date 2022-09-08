package main

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/yandex-cloud/cq-provider-yandex/gen/util"
)

func generate(resource, manager string) {
	util.SilentExecute(util.TemplatesDir{
		MainFile: "access_bindings_by.go.tmpl",
		Path:     "templates",
	}, map[string]string{
		"resource": resource,
		"manager":  manager,
	}, fmt.Sprintf("%s/access_bindings_by_%s.go", util.ResourcesDir, strcase.ToSnake(resource)))

	util.SilentExecute(util.TemplatesDir{
		MainFile: "iam_user_accounts_by.go.tmpl",
		Path:     "templates",
	}, map[string]string{
		"resource": resource,
		"manager":  manager,
	}, fmt.Sprintf("%s/iam_user_accounts_by_%s.go", util.ResourcesDir, strcase.ToSnake(resource)))
}

func main() {
	generate("Organization", "OrganizationManager")
	generate("Cloud", "ResourceManager")
	generate("Folder", "ResourceManager")
}
