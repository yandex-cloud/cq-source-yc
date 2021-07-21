package gen

import (
	"strings"

	"github.com/jhump/protoreflect/desc"

	"github.com/jinzhu/inflection"

	"github.com/iancoleman/strcase"
)

type ResourceFileModel struct {
	Table     *TableModel
	Relations []*TableModel
}

type TableModel struct {
	Service string

	Resource *desc.MessageDescriptor

	AbsolutFieldPath []*desc.FieldDescriptor

	ParentMessage     *desc.MessageDescriptor
	RelativeFieldPath []*desc.FieldDescriptor

	Multiplex string

	Columns   []*ColumnModel
	Relations []*TableModel
}

func (t TableModel) ServiceSnake() string {
	return strcase.ToSnake(t.Service)
}

func (t TableModel) ServiceCamel() string {
	return strcase.ToCamel(t.Service)
}

func (t TableModel) ResourceSnake() string {
	return strcase.ToSnake(t.Resource.GetName())
}

func (t TableModel) ResourceCamel() string {
	return strcase.ToCamel(t.Resource.GetName())
}

func (t TableModel) ResourcesSnake() string {
	return strcase.ToSnake(inflection.Plural(t.Resource.GetName()))
}

func (t TableModel) ResourcesCamel() string {
	return strcase.ToCamel(inflection.Plural(t.Resource.GetName()))
}

func (t TableModel) AbsolutFieldPathSnake() string {
	return strcase.ToSnake(t.AbsolutFieldPathCamel())
}

func (t TableModel) AbsolutFieldPathCamel() string {
	if len(t.AbsolutFieldPath) == 0 {
		return ""
	}
	path := make([]string, 0, len(t.AbsolutFieldPath))
	for _, field := range t.AbsolutFieldPath[:len(t.AbsolutFieldPath)-1] {
		path = append(path, strcase.ToCamel(inflection.Singular(field.GetJSONName())))
	}
	path = append(path, strcase.ToCamel(inflection.Plural(t.AbsolutFieldPath[len(t.AbsolutFieldPath)-1].GetJSONName())))
	return strings.Join(path, "")
}

func (t TableModel) ParentMessageCamel() string {
	return asGoType(t.ParentMessage)
}

func (t TableModel) RelativeFieldPathCamelWithDot() string {
	path := make([]string, 0, len(t.RelativeFieldPath))
	for _, field := range t.RelativeFieldPath {
		path = append(path, strcase.ToCamel(field.GetJSONName()))
	}
	return strings.Join(path, ".")
}

func asGoType(message *desc.MessageDescriptor) string {
	return strings.ReplaceAll(strings.TrimPrefix(message.GetFullyQualifiedName(), message.GetFile().GetPackage()+"."), ".", "_")
}

type ColumnModel struct {
	Name        string
	Description string
	Type        string
	Resolver    string
}
