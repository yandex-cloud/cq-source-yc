package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"github.com/yandex-cloud/cq-provider-yandex/gen/util"
	"github.com/yandex-cloud/cq-provider-yandex/gen/util/modelfromproto"
)

func generate(resource, pathToProto string, opts ...modelfromproto.Option) {
	opts = append(opts, modelfromproto.WithProtoPaths("cloudapi", "cloudapi/third_party/googleapis"))

	resourceFileModel, err := modelfromproto.ResourceFileFromProto(resource, resource, pathToProto, opts...)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}

	out := filepath.Join(util.ResourcesDir, "serverless_"+strcase.ToSnake(inflection.Plural(resource))+".go")

	util.SilentExecute(util.TemplatesDir{
		MainFile: "serverless.go.tmpl",
		Path:     "templates",
	}, resourceFileModel, out)
}

func main() {
	generate("ApiGateway", "yandex/cloud/serverless/apigateway/v1/apigateway.proto")
}
