package gen

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
)

type tableBuilder struct {
	service  string
	resource *desc.MessageDescriptor

	absolutPath  []*desc.FieldDescriptor
	relativePath []*desc.FieldDescriptor

	multiplex      string
	messageDesc    *desc.MessageDescriptor
	defaultColumns map[string]*ColumnModel
	ignoredFields  map[string]struct{}
}

func (b *tableBuilder) WithMessageFromProto(messageName, pathToProto string, paths ...string) error {
	parser := protoparse.Parser{IncludeSourceCodeInfo: true, ImportPaths: paths}

	protoFiles, err := parser.ParseFiles(pathToProto)
	if err != nil {
		return err
	}

	protoFile := protoFiles[0]

	b.messageDesc = protoFile.FindMessage(protoFile.GetPackage() + "." + messageName)
	if b.messageDesc == nil {
		return fmt.Errorf("messageDesc %v not found", messageName)
	}

	b.resource = b.messageDesc
	return nil
}

func (b *tableBuilder) Build() (*TableModel, error) {
	if b.messageDesc == nil {
		return nil, fmt.Errorf("source of messageDesc wasn't specified")
	}

	expandedFields := b.expandFields(b.messageDesc.GetFields(), nil)
	forColumns, forRelations := b.filterFields(expandedFields)

	return &TableModel{
		Service:      b.service,
		Resource:     strcase.ToCamel(b.resource.GetName()),
		AbsolutPath:  fieldsToStrings(b.absolutPath),
		RelativePath: fieldsToStrings(b.relativePath),
		Multiplex:    b.multiplex,
		Columns:      b.generateColumns(forColumns),
		Relations:    b.generateRelations(forRelations),
	}, nil
}

func (b *tableBuilder) expandFields(fields []*desc.FieldDescriptor, path []*desc.FieldDescriptor) (expandedFields []expandedField) {
	for _, field := range fields {
		newExpandedField := expandedField{field, path}

		newPath := path
		newPath = append(newPath, field)

		switch {
		case b.containsIgnoredField(newExpandedField):
			continue
		case isExpandable(field) && !b.containsDefaultColumn(newExpandedField):
			expandedFields = append(expandedFields, b.expandFields(field.GetMessageType().GetFields(), newPath)...)
		default:
			expandedFields = append(expandedFields, newExpandedField)
		}
	}
	return
}

func (b *tableBuilder) filterFields(fields []expandedField) (forColumns []expandedField, forRelations []expandedField) {
	for _, field := range fields {
		if !isConvertableToRelation(field) || b.containsDefaultColumn(field) {
			forColumns = append(forColumns, field)
		} else {
			forRelations = append(forRelations, field)
		}
	}
	return
}

func (b *tableBuilder) containsDefaultColumn(field expandedField) bool {
	_, ok := b.defaultColumns[field.getAbsolutPath(b.absolutPath)]
	return ok
}

func (b *tableBuilder) containsIgnoredField(field expandedField) bool {
	_, ok := b.ignoredFields[field.getAbsolutPath(b.absolutPath)]
	return ok
}

func (b *tableBuilder) generateColumns(fields []expandedField) (columns []*ColumnModel) {
	for _, field := range fields {
		if col, defined := b.defaultColumns[field.getAbsolutPath(b.absolutPath)]; defined {
			columns = append(columns, col)
		} else {
			columns = append(columns, &ColumnModel{
				Name:        field.getColumnName(),
				Type:        field.getType(),
				Description: strings.TrimSpace(field.GetSourceInfo().GetLeadingComments()),
				Resolver:    fmt.Sprintf("%v(\"%v\")", field.getResolver(), field.getPath()),
			})
		}
	}
	return
}

func (b *tableBuilder) generateRelations(fields []expandedField) []*TableModel {
	tables := make([]*TableModel, 0, len(fields))

	for _, field := range fields {
		relativePath := field.path
		relativePath = append(relativePath, field.FieldDescriptor)

		absolutPath := b.absolutPath
		absolutPath = append(absolutPath, relativePath...)

		builder := tableBuilder{
			service:        b.service,
			resource:       b.resource,
			absolutPath:    absolutPath,
			relativePath:   relativePath,
			multiplex:      "client.IdentityMultiplex",
			messageDesc:    field.GetMessageType(),
			ignoredFields:  b.ignoredFields,
			defaultColumns: b.defaultColumns,
		}

		table, err := builder.Build()

		if err != nil {
			continue
		}

		tables = append(tables, table)
	}

	return tables
}
