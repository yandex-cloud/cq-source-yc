package gen

import (
	"strings"
	"text/template"

	"github.com/jinzhu/inflection"

	"github.com/iancoleman/strcase"
)

func togetherCase(s string) string {
	return strings.ToLower(strcase.ToCamel(s))
}

var generatorTemplateFunctions = template.FuncMap{
	"together": togetherCase,
	"snake":    strcase.ToSnake,
	"camel":    strcase.ToCamel,
	"plural":   inflection.Plural,

	"join": func(sep string, elems []string) string { return strings.Join(elems, sep) },
	"asFullFieldName": func(names []string) []string {
		if len(names) == 0 {
			return names
		}
		for i := 0; i < len(names)-1; i++ {
			names[i] = inflection.Singular(names[i])
		}
		names[len(names)-1] = inflection.Plural(names[len(names)-1])
		return names
	},

	"replaceSymmetricKey": func(resource string) string {
		if resource == "SymmetricKey" {
			return "Key"
		} else {
			return resource
		}
	},
}
