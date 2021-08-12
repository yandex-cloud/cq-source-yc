package gen

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
)

type TemplatesDir struct {
	MainFile string
	Path     string
}

func ToTogether(s string) string {
	return strings.ToLower(strcase.ToCamel(s))
}

var templateFunctions = template.FuncMap{
	"together": ToTogether,
	"snake":    strcase.ToSnake,
	"camel":    strcase.ToCamel,
	"plural":   inflection.Plural,
	"join":     func(sep string, elems []string) string { return strings.Join(elems, sep) },
	"asFqn": func(names []string) []string {
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

func Execute(dir TemplatesDir, data interface{}, out string) error {
	file, err := os.Create(out)
	if err != nil {
		return err
	}

	files, err := filesInDir(dir.Path)
	if err != nil {
		return err
	}

	tmpl, err := template.New(dir.MainFile).Funcs(templateFunctions).ParseFiles(files...)
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}

	return file.Close()
}

func filesInDir(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
