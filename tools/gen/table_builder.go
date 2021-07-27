package gen

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
)

type expandedField struct {
	*desc.FieldDescriptor
	path []*desc.FieldDescriptor
}

func (f expandedField) getPathToResolve() string {
	path := f.getPath()
	for i := range path {
		path[i] = strcase.ToCamel(path[i])
	}
	return strings.Join(path, ".")
}

func (f expandedField) getColumnName() string {
	path := f.getPath()
	for i := range path {
		path[i] = strcase.ToSnake(path[i])
	}
	return strings.Join(path, "_")
}

func (f expandedField) getPath() []string {
	path := make([]string, 0, len(f.path)+1)
	for _, field := range f.path {
		path = append(path, strcase.ToCamel(field.GetJSONName()))
	}
	path = append(path, strcase.ToCamel(f.GetJSONName()))
	return path
}

type TableBuilder struct {
	service  string
	resource *desc.MessageDescriptor

	absolutFieldPath []*desc.FieldDescriptor

	parentMessage     *desc.MessageDescriptor
	relativeFieldPath []*desc.FieldDescriptor

	multiplex      string
	messageDesc    *desc.MessageDescriptor
	defaultColumns map[string]*ColumnModel
	ignoredFields  map[string]struct{}
}

func (b *TableBuilder) WithMessageFromProto(messageName, pathToProto string, paths ...string) error {
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

func (b *TableBuilder) Build() (*TableModel, error) {
	if b.messageDesc == nil {
		return nil, fmt.Errorf("source of messageDesc wasn't specified")
	}

	expandedFields := expandFields(b, b.messageDesc.GetFields(), nil)
	forColumns, forRelations := filterFields(b, expandedFields)

	return &TableModel{
		Service:           b.service,
		Resource:          b.resource,
		AbsolutFieldPath:  b.absolutFieldPath,
		ParentMessage:     b.parentMessage,
		RelativeFieldPath: b.relativeFieldPath,
		Multiplex:         b.multiplex,
		Columns:           generateColumns(b, forColumns),
		Relations:         generateRelations(b, forRelations),
	}, nil
}

// TODO: bug in oneof expansion
func expandFields(b *TableBuilder, fields []*desc.FieldDescriptor, path []*desc.FieldDescriptor) (expandedFields []expandedField) {
	for _, field := range fields {
		newPath := path
		newPath = append(newPath, field)
		newExpandedField := expandedField{field, path}
		switch {
		case b.containsIgnoredField(newExpandedField.getPathToResolve()):
			continue
		case isExpandable(field) && !b.containsDefaultColumn(newExpandedField.getPathToResolve()):
			expandedFields = append(expandedFields, expandFields(b, field.GetMessageType().GetFields(), newPath)...)
		default:
			expandedFields = append(expandedFields, newExpandedField)
		}
	}
	return
}

func filterFields(b *TableBuilder, fields []expandedField) (forColumns []expandedField, forRelations []expandedField) {
	for _, field := range fields {
		if !isConvertableToRelation(field) || b.containsDefaultColumn(field.getPathToResolve()) {
			forColumns = append(forColumns, field)
		} else {
			forRelations = append(forRelations, field)
		}
	}
	return
}

func (b *TableBuilder) containsDefaultColumn(path string) bool {
	absolutFieldPath := fieldsToStrings(b.absolutFieldPath)
	absolutFieldPath = append(absolutFieldPath, path)
	_, ok := b.defaultColumns[strings.Join(absolutFieldPath, ".")]
	return ok
}

func (b *TableBuilder) containsIgnoredField(path string) bool {
	absolutFieldPath := fieldsToStrings(b.absolutFieldPath)
	absolutFieldPath = append(absolutFieldPath, path)
	_, ok := b.ignoredFields[strings.Join(absolutFieldPath, ".")]
	return ok
}

func fieldsToStrings(fields []*desc.FieldDescriptor) []string {
	result := make([]string, 0, len(fields))
	for _, field := range fields {
		result = append(result, strcase.ToCamel(field.GetJSONName()))
	}
	return result
}

func isExpandable(field *desc.FieldDescriptor) bool {
	return !field.IsRepeated() && !field.IsMap() && field.GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE
}

func isConvertableToRelation(field expandedField) bool {
	return field.IsRepeated() && !field.IsMap() && field.GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE
}

func generateColumns(b *TableBuilder, fields []expandedField) (columns []*ColumnModel) {
	for _, field := range fields {
		if col, defined := b.defaultColumns[field.getPathToResolve()]; defined {
			columns = append(columns, col)
		} else {
			var resolver string
			if field.GetEnumType() == nil {
				resolver = "schema.PathResolver"
			} else {
				resolver = "client.EnumPathResolver"
			}
			columns = append(columns, &ColumnModel{
				Name:        field.getColumnName(),
				Type:        getType(field),
				Description: strings.TrimSpace(field.GetSourceInfo().GetLeadingComments()),
				Resolver:    fmt.Sprintf("%v(\"%v\")", resolver, field.getPathToResolve()),
			})
		}
	}
	return
}

func getType(field expandedField) string {
	switch {
	case field.IsRepeated() && field.GetType() == descriptor.FieldDescriptorProto_TYPE_STRING:
		return "schema.TypeStringArray"
	case !field.IsRepeated() && field.GetType() == descriptor.FieldDescriptorProto_TYPE_STRING:
		return "schema.TypeString"
	case !field.IsRepeated() && field.GetType() == descriptor.FieldDescriptorProto_TYPE_INT64:
		return "schema.TypeBigInt"
	case !field.IsRepeated() && field.GetType() == descriptor.FieldDescriptorProto_TYPE_INT32:
		return "schema.TypeInt"
	case field.IsRepeated() && field.GetType() == descriptor.FieldDescriptorProto_TYPE_INT32:
		return "schema.TypeIntArray"
	case !field.IsRepeated() && field.GetType() == descriptor.FieldDescriptorProto_TYPE_BOOL:
		return "schema.TypeBool"
	case field.IsMap():
		return "schema.TypeJSON"
	default:
		return "schema.TypeString"
	}
}

func generateRelations(b *TableBuilder, fields []expandedField) []*TableModel {
	tables := make([]*TableModel, 0, len(fields))
	for _, field := range fields {

		relativeFieldPath := field.path
		relativeFieldPath = append(relativeFieldPath, field.FieldDescriptor)

		absolutFieldPath := b.absolutFieldPath
		absolutFieldPath = append(absolutFieldPath, relativeFieldPath...)

		builder := TableBuilder{
			service:           b.service,
			resource:          b.resource,
			absolutFieldPath:  absolutFieldPath,
			parentMessage:     b.messageDesc,
			relativeFieldPath: relativeFieldPath,
			multiplex:         "client.IdentityMultiplex",
			messageDesc:       field.GetMessageType(),
			ignoredFields:     b.ignoredFields,
			defaultColumns:    b.defaultColumns,
		}

		table, err := builder.Build()

		if err != nil {
			continue
		}

		tables = append(tables, table)
	}
	return tables
}
