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
	absolutePath []string
	relativePath []string
	multiplex    string

	messageDesc *desc.MessageDescriptor

	defaultColumns map[string]*ColumnModel
	ignoredFields  map[string]struct{}
	Aliases        map[string]string
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

	return &TableModel{
		Service:      tb.service,
		Resource:     tb.resource,
		AbsolutePath: tb.absolutePath,
		RelativePath: tb.relativePath,
		Multiplex:    tb.multiplex,
		Columns:      tb.generateColumns(forColumns),
		Relations:    tb.generateRelations(forRelations),
		Alias:        tb.Aliases[strings.Join(tb.absolutePath, ".")],
	}, nil
}

func (tb *tableBuilder) expandFields(fields []*desc.FieldDescriptor, path []string) (expandedFields []expandedField) {
	for _, field := range fields {
		newExpandedField := expandedField{field, path}

		newPath := path
		newPath = append(newPath, getCamelName(field))

		switch {
		case tb.containsIgnoredField(newExpandedField):
			continue
		case isExpandable(field) && !tb.containsDefaultColumn(newExpandedField):
			expandedFields = append(expandedFields, tb.expandFields(field.GetMessageType().GetFields(), newPath)...)
		default:
			expandedFields = append(expandedFields, newExpandedField)
		}
	}
	return
}

func (tb *tableBuilder) filterFields(fields []expandedField) (forColumns []expandedField, forRelations []expandedField) {
	for _, field := range fields {
		if !field.isConvertableToRelation() || tb.containsDefaultColumn(field) {
			forColumns = append(forColumns, field)
		} else {
			forRelations = append(forRelations, field)
		}
	}
	return
}

func (tb *tableBuilder) containsDefaultColumn(field expandedField) bool {
	_, ok := tb.defaultColumns[field.resolvePath(tb.absolutePath)]
	return ok
}

func (tb *tableBuilder) containsIgnoredField(field expandedField) bool {
	_, ok := tb.ignoredFields[field.resolvePath(tb.absolutePath)]
	return ok
}

func (tb *tableBuilder) generateColumns(fields []expandedField) (columns []*ColumnModel) {
	for _, field := range fields {
		if col, defined := tb.defaultColumns[field.resolvePath(tb.absolutePath)]; defined {
			columns = append(columns, col)
		} else {
			var name string
			if alias, ok := tb.Aliases[field.resolvePath(tb.absolutePath)]; ok {
				name = alias
			} else {
				name = field.getColumnName()
			}
			columns = append(columns, &ColumnModel{
				Name:        name,
				Type:        field.getType(),
				Description: strings.TrimSpace(field.GetSourceInfo().GetLeadingComments()),
				Resolver:    fmt.Sprintf("%v(\"%v\")", field.getResolver(), field.getPath()),
			})
		}
	}
	return
}

func (tb *tableBuilder) generateRelations(fields []expandedField) []*TableModel {
	tables := make([]*TableModel, 0, len(fields))

	for _, field := range fields {
		relativePath := field.path
		relativePath = append(relativePath, getCamelName(field))

		absolutePath := tb.absolutePath
		absolutePath = append(absolutePath, relativePath...)

		builder := tableBuilder{
			service:        tb.service,
			resource:       tb.resource,
			absolutePath:   absolutePath,
			relativePath:   relativePath,
			multiplex:      "client.IdentityMultiplex",
			messageDesc:    field.GetMessageType(),
			ignoredFields:  tb.ignoredFields,
			defaultColumns: tb.defaultColumns,
			Aliases:        tb.Aliases,
		}

		table, err := builder.Build()

		if err != nil {
			continue
		}

		tables = append(tables, table)
	}

	return tables
}
