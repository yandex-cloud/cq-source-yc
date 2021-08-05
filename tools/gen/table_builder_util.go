package gen

import (
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"

	"github.com/iancoleman/strcase"
	"github.com/jhump/protoreflect/desc"
)

type OneOfer interface {
	GetOneOf() *desc.OneOfDescriptor
}

func getCamelName(d desc.Descriptor) string {
	if fd, ok := d.(OneOfer); ok && fd.GetOneOf() != nil {
		return strcase.ToCamel(fd.GetOneOf().GetName()) + "." + strcase.ToCamel(d.GetName())
	}
	return strcase.ToCamel(d.GetName())
}

func isExpandable(f *desc.FieldDescriptor) bool {
	return !f.IsRepeated() && !f.IsMap() && f.GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE
}

type expandedField struct {
	*desc.FieldDescriptor
	path []string
}

func (f expandedField) getPath() string {
	var path []string

	path = append(path, f.path...)

	path = append(path, getCamelName(f))

	return strings.Join(path, ".")
}

func (f expandedField) resolvePath(path []string) string {
	path = append(path, f.getPath())
	return strings.Join(path, ".")
}

func (f expandedField) isConvertableToRelation() bool {
	return f.IsRepeated() && !f.IsMap() && f.GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE
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

func join(paths ...string) string {
	var filteredPaths []string
	for _, path := range paths {
		if len(path) != 0 {
			filteredPaths = append(filteredPaths, path)
		}
	}
	return strings.Join(filteredPaths, ".")
}

func split(path string) []string {
	return strings.Split(path, ".")
}
