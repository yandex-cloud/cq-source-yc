package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"

	"github.com/yandex-cloud/cq-provider-yandex/gen/util"
)

func filterSchemas(decls []ast.Decl) (filtered []*ast.FuncDecl) {
	for _, decl := range decls {
		if isSchemaTable(decl) {
			filtered = append(filtered, decl.(*ast.FuncDecl))
		}
	}
	return
}

func isSchemaTable(decl ast.Decl) bool {
	funcDecl, ok := decl.(*ast.FuncDecl)
	if !ok {
		return false
	}

	if !(funcDecl.Type != nil && funcDecl.Type.Params != nil && funcDecl.Type.Params.List == nil &&
		funcDecl.Type.Results != nil && len(funcDecl.Type.Results.List) == 1) {
		return false
	}

	starExpr, ok := funcDecl.Type.Results.List[0].Type.(*ast.StarExpr)
	if !ok {
		return false
	}

	starExprX, ok := starExpr.X.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	starExprXXIdent, ok := starExprX.X.(*ast.Ident)
	if !ok {
		return false
	}

	return starExprXXIdent.Name == "schema" && starExprX.Sel.Name == "Table"
}

func main() {
	files, err := util.FilesInDir(util.ResourcesDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	resourceMap := map[string]struct{}{}

	for _, file := range files {
		if filepath.Ext(file) != ".go" {
			continue
		}

		parsedFile, err := parser.ParseFile(token.NewFileSet(), file, nil, 0)
		if err != nil {
			continue
		}

		funcDecls := filterSchemas(parsedFile.Decls)

		for _, decl := range funcDecls {
			resourceMap[decl.Name.Name] = struct{}{}
		}
	}

	util.SilentExecute(util.TemplatesDir{
		MainFile: "provider.go.tmpl",
		Path:     "templates",
	}, resourceMap, filepath.Join(util.ResourcesDir, "provider.go"))
}
