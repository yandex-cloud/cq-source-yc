package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc"
	"github.com/thoas/go-funk"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

type collapsedOptions struct {
	tableName      string
	message        *desc.MessageDescriptor
	resolver       schema.TableResolver
	defaultColumns map[string]schema.Column
	ignoredFields  map[string]struct{}
}

func (co *collapsedOptions) containsDefaultColumn(path string) bool {
	_, ok := co.defaultColumns[path]
	return ok
}

func (co *collapsedOptions) containsIgnoredField(path string) bool {
	_, ok := co.ignoredFields[path]
	return ok
}

type expandedField struct {
	*desc.FieldDescriptor
	path []*desc.FieldDescriptor
}

func (ef *expandedField) getPathToResolve() string {
	path := make([]string, 0, len(ef.path)+1)
	for _, field := range ef.path {
		path = append(path, strings.Title(field.GetJSONName()))
	}
	path = append(path, strings.Title(ef.GetJSONName()))
	return strings.Join(path, ".")
}

func (ef expandedField) getColumnName() string {
	path := make([]string, 0, len(ef.path)+1)
	for _, field := range ef.path {
		path = append(path, field.GetName())
	}
	path = append(path, ef.GetName())
	return strings.Join(path, "_")
}

func GenerateTable(opts ...Option) (*schema.Table, error) {
	co := &collapsedOptions{}

	for _, opt := range opts {
		err := opt.Apply(co)
		if err != nil {
			return nil, err
		}
	}

	if co.message == nil {
		return nil, fmt.Errorf("source of message descriptor wasn't specified")
	}

	expandedFields := expandFields(co, co.message.GetFields(), nil)
	forColumns, forRelations := filterFields(co, expandedFields)

	return &schema.Table{
		Name:         co.tableName,
		Resolver:     co.resolver,
		Multiplex:    client.FolderMultiplex,
		IgnoreError:  client.IgnoreErrorHandler,
		DeleteFilter: client.DeleteFolderFilter,
		Columns:      generateColumns(co, forColumns),
		Relations:    generateRelations(co, forRelations),
	}, nil
}

func expandFields(co *collapsedOptions, fields []*desc.FieldDescriptor, path []*desc.FieldDescriptor) (expandedFields []expandedField) {
	for _, field := range fields {
		newPath := path
		newPath = append(newPath, field)
		newExpandedField := expandedField{field, path}
		switch {
		case co.containsIgnoredField(newExpandedField.getPathToResolve()):
			continue
		case isExpandable(field) && !co.containsDefaultColumn(newExpandedField.getPathToResolve()):
			expandedFields = append(expandedFields, expandFields(co, field.GetMessageType().GetFields(), newPath)...)
		default:
			expandedFields = append(expandedFields, newExpandedField)
		}
	}
	return
}

func filterFields(co *collapsedOptions, fields []expandedField) (forColumns []expandedField, forRelations []expandedField) {
	for _, field := range fields {
		if !isConvertableToRelation(field) || co.containsDefaultColumn(field.getPathToResolve()) {
			forColumns = append(forColumns, field)
		} else {
			forRelations = append(forRelations, field)
		}
	}
	return
}

func isExpandable(field *desc.FieldDescriptor) bool {
	return !field.IsRepeated() && !field.IsMap() && field.GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE
}

func isConvertableToRelation(field expandedField) bool {
	return field.IsRepeated() && !field.IsMap() && field.GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE
}

func generateColumns(co *collapsedOptions, fields []expandedField) (columns []schema.Column) {
	for _, field := range fields {
		if col, def := co.defaultColumns[field.getPathToResolve()]; def {
			columns = append(columns, col)
		} else {
			columns = append(columns, schema.Column{
				Name:        field.getColumnName(),
				Type:        getType(field),
				Description: field.GetSourceInfo().GetLeadingComments(),
				Resolver:    getPathResolver(field.getPathToResolve(), field.GetEnumType() != nil),
			})
		}
	}
	return
}

func getPathResolver(path string, isEnum bool) schema.ColumnResolver {
	return func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		if isEnum {
			return r.Set(c.Name, funk.Get(r.Item, path, funk.WithAllowZero()).(fmt.Stringer).String())
		} else {
			return r.Set(c.Name, funk.Get(r.Item, path, funk.WithAllowZero()))
		}
	}
}

func getType(field expandedField) schema.ValueType {
	switch {
	case field.IsRepeated() && field.GetType() == descriptor.FieldDescriptorProto_TYPE_STRING:
		return schema.TypeStringArray
	case !field.IsRepeated() && field.GetType() == descriptor.FieldDescriptorProto_TYPE_STRING:
		return schema.TypeString
	case !field.IsRepeated() && field.GetType() == descriptor.FieldDescriptorProto_TYPE_INT64:
		return schema.TypeBigInt
	case !field.IsRepeated() && field.GetType() == descriptor.FieldDescriptorProto_TYPE_INT32:
		return schema.TypeInt
	case field.IsRepeated() && field.GetType() == descriptor.FieldDescriptorProto_TYPE_INT32:
		return schema.TypeIntArray
	case !field.IsRepeated() && field.GetType() == descriptor.FieldDescriptorProto_TYPE_BOOL:
		return schema.TypeBool
	case field.IsMap():
		return schema.TypeJSON
	default:
		return schema.TypeString
	}
}

func generateRelations(co *collapsedOptions, fields []expandedField) (tables []*schema.Table) {
	for _, field := range fields {
		table, err := GenerateTable(
			WithTableName(co.tableName+"_"+field.getColumnName()),
			WithMessage(field.GetMessageType()),
			WithResolver(getRelationResolver(field.getPathToResolve())),
		)
		if err != nil {
			continue
		}
		tables = append(tables, table)
	}
	return
}

func getRelationResolver(relativeFieldName string) schema.TableResolver {
	return func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
		// TODO
		return nil
	}
}
