package gen

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
)

type tableBuilder struct {
	service      string
	resource     string
	absolutePath string
	relativePath string
	multiplex    string

	messageDesc *desc.MessageDescriptor

	ignoredFields map[string]struct{}
	aliases       map[string]Alias

	field  *expandedField
	parent *tableBuilder
}

func (tb *tableBuilder) WithMessageFromProto(messageName, pathToProto string, paths ...string) error {
	parser := protoparse.Parser{IncludeSourceCodeInfo: true, ImportPaths: paths}

	protoFiles, err := parser.ParseFiles(pathToProto)
	if err != nil {
		return err
	}

	protoFile := protoFiles[0]

	tb.messageDesc = protoFile.FindMessage(protoFile.GetPackage() + "." + messageName)
	if tb.messageDesc == nil {
		return fmt.Errorf("messageDesc %s not found", messageName)
	}

	tb.resource = getCamelName(tb.messageDesc)
	return nil
}

func (tb *tableBuilder) Build() (*TableModel, error) {
	if tb.messageDesc == nil {
		return nil, fmt.Errorf("source of messageDesc wasn't specified")
	}

	expandedFields := tb.expandFields(tb.messageDesc.GetFields(), nil)
	forColumns, forRelations := tb.filterFields(expandedFields)

	relations, err := tb.generateRelations(forRelations)
	if err != nil {
		return nil, err
	}

	table := &TableModel{
		Service:      tb.service,
		Resource:     tb.resource,
		AbsolutePath: split(tb.absolutePath),
		RelativePath: split(tb.relativePath),
		Multiplex:    tb.multiplex,
		Columns:      tb.generateColumns(forColumns),
		Relations:    relations,
	}

	if alias, ok := tb.aliases[tb.absolutePath]; ok {
		alias.ApplyToTable(table)
	}

	return table, nil
}

func (tb *tableBuilder) expandFields(fields []*desc.FieldDescriptor, path []string) (expandedFields []expandedField) {
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

func (tb *tableBuilder) filterFields(fields []expandedField) (forColumns []expandedField, forRelations []expandedField) {
	for _, field := range fields {
		if !field.isConvertableToRelation() {
			forColumns = append(forColumns, field)
		} else {
			forRelations = append(forRelations, field)
		}
	}
	return
}

func (tb *tableBuilder) containsIgnoredField(field expandedField) bool {
	_, ok := tb.ignoredFields[join(tb.absolutePath, field.getPath())]
	return ok
}

func (tb *tableBuilder) containsAliases(field expandedField) bool {
	_, ok := tb.aliases[join(tb.absolutePath, field.getPath())]
	return ok
}

func (tb *tableBuilder) generateColumns(fields []expandedField) (columns []*ColumnModel) {
	columns = tb.appendIfRelation(columns)
	for _, field := range fields {
		column := &ColumnModel{
			Name:        field.getColumnName(),
			Type:        field.getType(),
			Description: strings.TrimSpace(field.GetSourceInfo().GetLeadingComments()),
			Resolver:    field.getResolver(),
		}

		if alias, ok := tb.aliases[join(tb.absolutePath, field.getPath())]; ok {
			alias.ApplyToColumn(column)
		}

		columns = append(columns, column)
	}
	return
}

func (tb *tableBuilder) appendIfRelation(columns []*ColumnModel) []*ColumnModel {
	if tb.parent != nil {
		var (
			parentName    string
			parentMsgDesc *desc.MessageDescriptor
		)

		if tb.parent.field == nil {
			parentName = strcase.ToSnake(tb.resource)
			parentMsgDesc = tb.parent.messageDesc
		} else {
			parentName = tb.parent.field.getColumnName()
			parentMsgDesc = tb.parent.field.GetMessageType()
		}

		columns = append(columns, &ColumnModel{
			Name:        parentName + "_cq_id",
			Type:        "schema.TypeUUID",
			Description: fmt.Sprintf("cq_id of parent %s", parentName),
			Resolver:    "schema.ParentIdResolver",
		})

		if parentMsgDesc.FindFieldByName("id") != nil {
			columns = append(columns, &ColumnModel{
				Name:        parentName + "_id",
				Type:        "schema.TypeString",
				Description: fmt.Sprintf("id of parent %s", parentName),
				Resolver:    "schema.ParentResourceFieldResolver(\"id\")",
			})
		}
	}
	return columns
}

func (tb *tableBuilder) generateRelations(fields []expandedField) ([]*TableModel, error) {
	tables := make([]*TableModel, 0, len(fields))

	for _, field := range fields {
		builder := tableBuilder{
			service:       tb.service,
			resource:      tb.resource,
			absolutePath:  join(tb.absolutePath, field.getPath()),
			relativePath:  field.getPath(),
			multiplex:     "client.IdentityMultiplex",
			messageDesc:   field.GetMessageType(),
			ignoredFields: tb.ignoredFields,
			aliases:       tb.aliases,
			field:         &field,
			parent:        tb,
		}

		table, err := builder.Build()

		if err != nil {
			return nil, err
		}

		tables = append(tables, table)
	}

	return tables, nil
}
