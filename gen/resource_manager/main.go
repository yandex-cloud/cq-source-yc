package main

import (
	"path/filepath"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"github.com/yandex-cloud/cq-provider-yandex/gen/util"
)

func generate(resource string) {
	out := filepath.Join(util.ResourcesDir, "resourcemanager_"+strcase.ToSnake(inflection.Plural(resource))+".go")

	util.SilentExecute(util.TemplatesDir{
		MainFile: "resourcemanager_resource.go.tmpl",
		Path:     "templates",
	}, map[string]string{
		"resource": resource,
	}, out)
}

func main() {
	generate("Cloud")
	generate("Folder")
}
