package serverless

import (
	"path/filepath"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"github.com/yandex-cloud/cq-provider-yandex/tools/gen"
	ycmb "github.com/yandex-cloud/cq-provider-yandex/tools/gen/ycmodelbuilder"
)

var templatesDir = gen.TemplatesDir{
	MainFile: "resource.go.tmpl",
	Path:     "tools/gen/serverless/templates",
}

func Generate(resource, pathToProto, outDir string, opts ...ycmb.Option) error {
	resourceFileModel, err := ycmb.ResourceFileFromProto(resource, resource, pathToProto, opts...)
	if err != nil {
		return err
	}

	out := filepath.Join(outDir, "serverless_"+strcase.ToSnake(inflection.Plural(resource))+".go")

	return gen.Execute(templatesDir, resourceFileModel, out)
}
