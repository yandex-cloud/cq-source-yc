package util

import (
	"fmt"
	"os"
	"os/exec"
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

var templateFunctions = template.FuncMap{
	"flat":   ToFlat,
	"snake":  strcase.ToSnake,
	"camel":  strcase.ToCamel,
	"plural": inflection.Plural,
	"join":   func(sep string, elems []string) string { return strings.Join(elems, sep) },
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

	files, err := FilesInDir(dir.Path)
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

	err = file.Close()
	if err != nil {
		return err
	}

	err = exec.Command("goimports", "-w", out).Run()
	if err != nil {
		return fmt.Errorf("goimports -w finished with error: %s", err)
	}
	return nil
}

func FilesInDir(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func SilentExecute(dir TemplatesDir, data interface{}, out string) {
	if err := Execute(dir, data, out); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
}
