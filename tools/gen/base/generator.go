package base

import (
	"path/filepath"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"github.com/yandex-cloud/cq-provider-yandex/tools/gen"
	ycmodel "github.com/yandex-cloud/cq-provider-yandex/tools/gen/ycmodelbuilder"
)

var templatesDir = gen.TemplatesDir{
	MainFile: "resource.go.tmpl",
	Path:     "tools/gen/base/templates",
}

func Generate(service, resource, pathToProto, outDir string, opts ...ycmodel.Option) error {
	resourceFileModel, err := ycmodel.ResourceFileFromProto(service, resource, pathToProto, opts...)
	if err != nil {
		return err
	}

	out := filepath.Join(outDir, gen.ToTogether(service)+"_"+strcase.ToSnake(inflection.Plural(resource))+".go")

	return gen.Execute(templatesDir, resourceFileModel, out)
}
