package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"github.com/yandex-cloud/cq-provider-yandex/gen/util"
	ycmodel "github.com/yandex-cloud/cq-provider-yandex/gen/util/ycmodelbuilder"
)

func generate(resource, pathToProto string, opts ...ycmodel.Option) {
	opts = append(opts, ycmodel.WithProtoPaths("cloudapi", "cloudapi/third_party/googleapis"))

	resourceFileModel, err := ycmodel.ResourceFileFromProto(resource, resource, pathToProto, opts...)
	if err != nil {
		fmt.Fprint(os.Stderr)
		return
	}

	out := filepath.Join(util.ResourcesDir, "serverless_"+strcase.ToSnake(inflection.Plural(resource))+".go")

	util.SilentExecute(util.TemplatesDir{
		MainFile: "serverless_resource.go.tmpl",
		Path:     "templates",
	}, resourceFileModel, out)
}

func main() {
	generate("ApiGateway", "yandex/cloud/serverless/apigateway/v1/apigateway.proto")
}
