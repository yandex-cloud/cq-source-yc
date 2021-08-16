package ycmodelbuilder

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/yandex-cloud/cq-provider-yandex/gen/util/ycmodel"
)

type TableBuilder struct {
	Service      string
	Resource     string
	AbsolutePath string
	RelativePath string
	Multiplex    string

	MessageDesc *desc.MessageDescriptor

	IgnoredFields map[string]struct{}
	Aliases       map[string]Alias

	Field  *expandedField
	Parent *TableBuilder
}

func (tb *TableBuilder) WithMessageFromProto(messageName, pathToProto string, paths ...string) error {
	parser := protoparse.Parser{IncludeSourceCodeInfo: true, ImportPaths: paths}

	protoFiles, err := parser.ParseFiles(pathToProto)
	if err != nil {
		return err
	}

	protoFile := protoFiles[0]

	tb.MessageDesc = protoFile.FindMessage(protoFile.GetPackage() + "." + messageName)
	if tb.MessageDesc == nil {
		return fmt.Errorf("MessageDesc %s not found", messageName)
	}

	tb.Resource = getCamelName(tb.MessageDesc)
	return nil
}

func (tb *TableBuilder) Build() (*ycmodel.Table, error) {
	if tb.MessageDesc == nil {
		return nil, fmt.Errorf("source of MessageDesc wasn't specified")
	}

	expandedFields := tb.expandFields(tb.MessageDesc.GetFields(), nil)
	forColumns, forRelations := tb.filterFields(expandedFields)

	relations, err := tb.generateRelations(forRelations)
	if err != nil {
		return nil, err
	}

	table := &ycmodel.Table{
		Service:      tb.Service,
		Resource:     tb.Resource,
		AbsolutePath: split(tb.AbsolutePath),
		RelativePath: split(tb.RelativePath),
		Multiplex:    tb.Multiplex,
		Columns:      tb.generateColumns(forColumns),
		Relations:    relations,
	}

	if alias, ok := tb.Aliases[tb.AbsolutePath]; ok {
		alias.ApplyToTable(table)
	}

	return table, nil
}

func (tb *TableBuilder) expandFields(fields []*desc.FieldDescriptor, path []string) (expandedFields []expandedField) {
	for _, field := range fields {
		newExpandedField := expandedField{field, path}

		newPath := path
		newPath = append(newPath, getCamelName(field))

		switch {
		case tb.containsIgnoredField(newExpandedField):
			continue
		case isExpandable(field) && !tb.containsAliases(newExpandedField):
			expandedFields = append(expandedFields, tb.expandFields(field.GetMessageType().GetFields(), newPath)...)
		default:
			expandedFields = append(expandedFields, newExpandedField)
		}
	}
	return
}

func (tb *TableBuilder) filterFields(fields []expandedField) (forColumns []expandedField, forRelations []expandedField) {
	for _, field := range fields {
		if !field.isConvertableToRelation() {
			forColumns = append(forColumns, field)
		} else {
			forRelations = append(forRelations, field)
		}
	}
	return
}

func (tb *TableBuilder) containsIgnoredField(field expandedField) bool {
	_, ok := tb.IgnoredFields[join(tb.AbsolutePath, field.getPath())]
	return ok
}

func (tb *TableBuilder) containsAliases(field expandedField) bool {
	_, ok := tb.Aliases[join(tb.AbsolutePath, field.getPath())]
	return ok
}

func (tb *TableBuilder) generateColumns(fields []expandedField) (columns []*ycmodel.Column) {
	columns = tb.appendIfRelation(columns)
	for _, field := range fields {
		column := &ycmodel.Column{
			Name:        field.getColumnName(),
			Type:        field.getType(),
			Description: strings.TrimSpace(field.GetSourceInfo().GetLeadingComments()),
			Resolver:    field.getResolver(),
		}

		if alias, ok := tb.Aliases[join(tb.AbsolutePath, field.getPath())]; ok {
			alias.ApplyToColumn(column)
		}

		columns = append(columns, column)
	}
	return
}

func (tb *TableBuilder) appendIfRelation(columns []*ycmodel.Column) []*ycmodel.Column {
	if tb.Parent != nil {
		var (
			parentName    string
			parentMsgDesc *desc.MessageDescriptor
		)

		if tb.Parent.Field == nil {
			parentName = strcase.ToSnake(tb.Resource)
			parentMsgDesc = tb.Parent.MessageDesc
		} else {
			parentName = tb.Parent.Field.getColumnName()
			parentMsgDesc = tb.Parent.Field.GetMessageType()
		}

		columns = append(columns, &ycmodel.Column{
			Name:        parentName + "_cq_id",
			Type:        "schema.TypeUUID",
			Description: fmt.Sprintf("cq_id of parent %s", parentName),
			Resolver:    "schema.ParentIdResolver",
		})

		if parentMsgDesc.FindFieldByName("id") != nil {
			columns = append(columns, &ycmodel.Column{
				Name:        parentName + "_id",
				Type:        "schema.TypeString",
				Description: fmt.Sprintf("id of parent %s", parentName),
				Resolver:    "schema.ParentResourceFieldResolver(\"id\")",
			})
		}
	}
	return columns
}

func (tb *TableBuilder) generateRelations(fields []expandedField) ([]*ycmodel.Table, error) {
	tables := make([]*ycmodel.Table, 0, len(fields))

	for _, field := range fields {
		builder := TableBuilder{
			Service:       tb.Service,
			Resource:      tb.Resource,
			AbsolutePath:  join(tb.AbsolutePath, field.getPath()),
			RelativePath:  field.getPath(),
			Multiplex:     "client.EmptyMultiplex",
			MessageDesc:   field.GetMessageType(),
			IgnoredFields: tb.IgnoredFields,
			Aliases:       tb.Aliases,
			Field:         &field,
			Parent:        tb,
		}

		table, err := builder.Build()

		if err != nil {
			return nil, err
		}

		tables = append(tables, table)
	}

	return tables, nil
}
