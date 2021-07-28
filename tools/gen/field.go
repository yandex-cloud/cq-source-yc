package gen

import (
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"

	"github.com/iancoleman/strcase"
	"github.com/jhump/protoreflect/desc"
)

type expandedField struct {
	*desc.FieldDescriptor
	path []*desc.FieldDescriptor
}

func (f expandedField) getPath() string {
	var path []string

	path = append(path, fieldsToStrings(f.path)...)

	// if f is a field within oneof
	if f.GetOneOf() != nil {
		path = append(path, strcase.ToCamel(f.GetOneOf().GetName()))
	}

	path = append(path, strcase.ToCamel(f.GetName()))

	return strings.Join(path, ".")
}

func (f expandedField) getAbsolutPath(relationAbsolutPath []*desc.FieldDescriptor) string {
	path := fieldsToStrings(relationAbsolutPath)
	path = append(path, f.getPath())
	return strings.Join(path, ".")
}

func (f expandedField) getColumnName() string {
	path := strings.Split(f.getPath(), ".")

	for i := range path {
		path[i] = strcase.ToSnake(path[i])
	}

	return strings.Join(path, "_")
}

func (f expandedField) getType() string {
	switch {
	case f.IsRepeated() && f.GetType() == descriptor.FieldDescriptorProto_TYPE_STRING:
		return "schema.TypeStringArray"
	case !f.IsRepeated() && f.GetType() == descriptor.FieldDescriptorProto_TYPE_STRING:
		return "schema.TypeString"
	case !f.IsRepeated() && f.GetType() == descriptor.FieldDescriptorProto_TYPE_INT64:
		return "schema.TypeBigInt"
	case !f.IsRepeated() && f.GetType() == descriptor.FieldDescriptorProto_TYPE_INT32:
		return "schema.TypeInt"
	case f.IsRepeated() && f.GetType() == descriptor.FieldDescriptorProto_TYPE_INT32:
		return "schema.TypeIntArray"
	case !f.IsRepeated() && f.GetType() == descriptor.FieldDescriptorProto_TYPE_BOOL:
		return "schema.TypeBool"
	case f.IsMap():
		return "schema.TypeJSON"
	default:
		return "schema.TypeString"
	}
}

func (f expandedField) getResolver() string {
	if f.GetType() != descriptor.FieldDescriptorProto_TYPE_ENUM {
		return "schema.PathResolver"
	} else {
		return "client.EnumPathResolver"
	}
}

func fieldsToStrings(fields []*desc.FieldDescriptor) []string {
	result := make([]string, 0, len(fields))

	for _, field := range fields {
		result = append(result, strcase.ToCamel(field.GetName()))
	}

	return result
}

func isExpandable(field *desc.FieldDescriptor) bool {
	return !field.IsRepeated() && !field.IsMap() && field.GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE
}

func isConvertableToRelation(field expandedField) bool {
	return field.IsRepeated() && !field.IsMap() && field.GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE
}
