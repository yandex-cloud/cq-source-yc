package gen

import (
	"fmt"
	"strings"

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
		return fmt.Errorf("messageDesc %v not found", messageName)
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

	table := &TableModel{
		Service:      tb.service,
		Resource:     tb.resource,
		AbsolutePath: split(tb.absolutePath),
		RelativePath: split(tb.relativePath),
		Multiplex:    tb.multiplex,
		Columns:      tb.generateColumns(forColumns),
		Relations:    tb.generateRelations(forRelations),
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
	for _, field := range fields {
		column := &ColumnModel{
			Name:        field.getColumnName(),
			Type:        field.getType(),
			Description: strings.TrimSpace(field.GetSourceInfo().GetLeadingComments()),
			Resolver:    fmt.Sprintf("%v(\"%v\")", field.getResolver(), field.getPath()),
		}

		if alias, ok := tb.aliases[join(tb.absolutePath, field.getPath())]; ok {
			alias.ApplyToColumn(column)
		}

		columns = append(columns, column)
	}
	return
}

func (tb *tableBuilder) generateRelations(fields []expandedField) []*TableModel {
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
		}

		table, err := builder.Build()

		if err != nil {
			continue
		}

		tables = append(tables, table)
	}

	return tables
}
