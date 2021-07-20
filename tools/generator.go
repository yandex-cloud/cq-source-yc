package main

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"os"
	"text/template"
)

func Generate(service, resource, pathToProto string, opts ...Option) error {
	co := NewCollapsedOptions(opts)

	tableBuilder := TableBuilder{
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

	file, err := os.Create(fmt.Sprintf("../resources/%v_%v.go", tableModel.ServiceSnake(), tableModel.ResourcesSnake()))

	if err != nil {
		return err
	}

	tmpl, err := template.New("resource.go.tmpl").ParseFiles(
		"template/column.go.tmpl",
		"template/relation_resolver.go.tmpl",
		"template/resource_resolver.go.tmpl",
		"template/resource.go.tmpl",
		"template/relation.go.tmpl",
		"template/table.go.tmpl",
	)

	if err != nil {
		return err
	}

	resourceFileModel := ResourceFileModel{tableModel, expandRelations(tableModel)}

	err = tmpl.Execute(file, resourceFileModel)

	if err != nil {
		return err
	}

	return nil
}

func (b TableBuilder) setDefaultYCColumns() {
	b.defaultColumns["Id"] = &ColumnModel{
		Name:     strcase.ToSnake(b.resource.GetName()) + "_id",
		Type:     "schema.TypeString",
		Resolver: "client.ResolveResourceId",
	}
	b.defaultColumns["FolderId"] = &ColumnModel{
		Name:     "folder_id",
		Type:     "schema.TypeString",
		Resolver: "client.ResolveFolderID",
	}
	b.defaultColumns["CreatedAt"] = &ColumnModel{
		Name:     "created_at",
		Type:     "schema.TypeTimestamp",
		Resolver: "client.ResolveAsTime",
	}
	b.defaultColumns["Labels"] = &ColumnModel{
		Name:     "labels",
		Type:     "schema.TypeJSON",
		Resolver: "client.ResolveLabels",
	}
}

func expandRelations(table *TableModel) (tables []*TableModel) {
	for _, relation := range table.Relations {
		tables = append(tables, expandRelations(relation)...)
		tables = append(tables, relation)
	}
	return
}
