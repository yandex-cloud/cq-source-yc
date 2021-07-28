package gen

import (
	"fmt"
	"os"
	"text/template"

	"github.com/jinzhu/inflection"

	"github.com/iancoleman/strcase"
)

func Generate(service, resource, pathToProto, outDir string, opts ...Option) error {
	co := NewCollapsedOptions(opts)

	tableBuilder := tableBuilder{
		service:        service,
		multiplex:      "client.FolderMultiplex",
		defaultColumns: co.defaultColumns,
		ignoredFields:  co.ignoredFields,
	}

	err := tableBuilder.WithMessageFromProto(resource, pathToProto, co.paths...)

	if err != nil {
		return err
	}

	tableBuilder.setDefaultYCColumns()

	tableModel, err := tableBuilder.Build()

	if err != nil {
		return err
	}

	file, err := os.Create(fmt.Sprintf("%v/%v_%v.go",
		outDir, strcase.ToSnake(tableModel.Service),
		strcase.ToSnake(inflection.Plural(tableModel.Resource))))

	if err != nil {
		return err
	}

	tmpl, err := template.New("resource.go.tmpl").Funcs(generatorTemplateFunctions).ParseFiles(
		"tools/gen/template/column.go.tmpl",
		"tools/gen/template/relation_resolver.go.tmpl",
		"tools/gen/template/resource_resolver.go.tmpl",
		"tools/gen/template/resource.go.tmpl",
		"tools/gen/template/relation.go.tmpl",
		"tools/gen/template/table.go.tmpl",
	)

	if err != nil {
		return err
	}

	resourceFileModel := ResourceFileModel{tableModel, expandRelations(tableModel)}

	err = tmpl.Execute(file, resourceFileModel)

	if err != nil {
		return err
	}

	return file.Close()
}

func (b tableBuilder) setDefaultYCColumns() {
	name := strcase.ToSnake(b.resource.GetName())

	b.defaultColumns["Id"] = &ColumnModel{
		Name:        name + "_id",
		Type:        "schema.TypeString",
		Description: fmt.Sprintf("ID of the %v.", name),
		Resolver:    "client.ResolveResourceId",
	}

	b.defaultColumns["FolderId"] = &ColumnModel{
		Name:        "folder_id",
		Type:        "schema.TypeString",
		Description: fmt.Sprintf("ID of the folder that the %v belongs to.", name),
		Resolver:    "client.ResolveFolderID",
	}

	b.defaultColumns["CreatedAt"] = &ColumnModel{
		Name:     "created_at",
		Type:     "schema.TypeTimestamp",
		Resolver: "client.ResolveAsTime",
	}

	b.defaultColumns["Labels"] = &ColumnModel{
		Name:        "labels",
		Type:        "schema.TypeJSON",
		Description: "Resource labels as `key:value` pairs. Maximum of 64 per resource.",
		Resolver:    "client.ResolveLabels",
	}
}

func expandRelations(table *TableModel) (tables []*TableModel) {
	for _, relation := range table.Relations {
		tables = append(tables, expandRelations(relation)...)
		tables = append(tables, relation)
	}
	return
}

func GenerateTests(service, resource, outDir string) error {
	file, err := os.Create(fmt.Sprintf("%v/%v_%v_test.go",
		outDir, strcase.ToSnake(service), strcase.ToSnake(inflection.Plural(resource))))

	if err != nil {
		return err
	}

	tmpl, err := template.New("resource_test.go.tmpl").Funcs(generatorTemplateFunctions).ParseFiles(
		"tools/gen/template/resource_test.go.tmpl",
	)

	if err != nil {
		return err
	}

	err = tmpl.Execute(file, ResourceTestFileModel{Resource: resource, Service: service})

	if err != nil {
		return err
	}

	return file.Close()
}
