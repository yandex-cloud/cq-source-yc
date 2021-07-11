package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/iancoleman/strcase"
	"github.com/jhump/protoreflect/desc"
	"github.com/thoas/go-funk"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

type DefaultColumns map[string]schema.Column

func getCommonYCDefaultCols(resourceName string) Option {
	return WithDefaultColumns(DefaultColumns{
		"Id": {
			Name:     strcase.ToSnake(resourceName) + "_id",
			Type:     schema.TypeString,
			Resolver: client.ResolveResourceId,
		},
		"FolderId": {
			Name:     "folder_id",
			Type:     schema.TypeString,
			Resolver: client.ResolveFolderID,
		},
		"CreatedAt": {
			Name:     "created_at",
			Type:     schema.TypeTimestamp,
			Resolver: client.ResolveAsTime,
		},
		"Labels": {
			Name:     "labels",
			Type:     schema.TypeJSON,
			Resolver: client.ResolveLabels,
		},
	})
}

type IgnoredColumns []string

type TableGenerator struct {
	tableName    string
	serviceName  string
	resourceName string
	protoFile    *desc.FileDescriptor
	defaultCols  DefaultColumns
	ignoreCols   map[string]bool
	fetcher      schema.TableResolver
}

type protoField struct {
	*desc.FieldDescriptor
	path []*desc.FieldDescriptor
}

func (pf protoField) getColName() string {
	var path []string
	for _, field := range pf.path {
		path = append(path, field.GetName())
	}
	path = append(path, pf.GetName())
	return strings.Join(path, "_")
}

func (pf protoField) getPath() string {
	var path []string
	for _, field := range pf.path {
		path = append(path, strings.Title(field.GetJSONName()))
	}
	path = append(path, strings.Title(pf.GetJSONName()))
	return strings.Join(path, ".")
}

func NewTableGenerator(serviceName string, resourceName string, opts ...Option) (*TableGenerator, error) {
	tg := &TableGenerator{
		tableName:    fmt.Sprintf("yandex_%v_%v", strcase.ToSnake(serviceName), strcase.ToSnake(resourceName)),
		serviceName:  serviceName,
		resourceName: resourceName,
	}

	for _, opt := range opts {
		opt.Apply(tg)
	}

	getCommonYCDefaultCols(resourceName).Apply(tg)

	return tg, nil
}

func (tg *TableGenerator) Generate() (*schema.Table, error) {
	if tg.protoFile == nil {
		return nil, fmt.Errorf("proto file wasn't parsed")
	}

	resourceQualifiedName := tg.protoFile.GetPackage() + "." + tg.resourceName
	mes := tg.protoFile.FindMessage(resourceQualifiedName)
	if mes == nil {
		return nil, fmt.Errorf("message with resourceName=%v not found in specified file", tg.resourceName)
	}

	expandedFields := tg.expandFields(mes.GetFields(), []*desc.FieldDescriptor{})
	forCols, _ := tg.filterFields(expandedFields)

	return &schema.Table{
		Name:         tg.tableName,
		Resolver:     tg.fetcher,
		Multiplex:    client.FolderMultiplex,
		IgnoreError:  client.IgnoreErrorHandler,
		DeleteFilter: client.DeleteFolderFilter,
		Columns:      tg.generateCols(forCols),
		Relations:    nil, // TODO
	}, nil
}

func (tg *TableGenerator) expandFields(fields []*desc.FieldDescriptor, path []*desc.FieldDescriptor) (expandedFields []protoField) {
	for _, field := range fields {
		newPath := append(path, field)
		newProtoField := protoField{field, path}
		switch {
		case tg.isIgnored(newProtoField.getPath()):
			continue
		case isExpandable(field) && !tg.hasDefaultValue(newProtoField.getPath()):
			expandedFields = append(expandedFields, tg.expandFields(field.GetMessageType().GetFields(), newPath)...)
		default:
			expandedFields = append(expandedFields, newProtoField)
		}
	}
	return
}

func (tg *TableGenerator) filterFields(fields []protoField) (forCols []protoField, forRelations []protoField) {
	for _, field := range fields {
		if !isConvertableToRelation(field) || tg.hasDefaultValue(field.getPath()) {
			forCols = append(forCols, field)
		} else {
			forRelations = append(forRelations, field)
		}
	}
	return
}

func (tg *TableGenerator) generateCols(fields []protoField) (cols []schema.Column) {
	for _, field := range fields {
		if col, def := tg.defaultCols[field.getPath()]; def {
			cols = append(cols, col)
		} else {
			cols = append(cols, schema.Column{
				Name:        field.getColName(),
				Type:        getType(field),
				Description: field.GetSourceInfo().GetLeadingComments(),
				Resolver:    PathResolver(field.getPath(), field.GetEnumType() != nil),
			})
		}
	}
	return
}

func PathResolver(path string, isEnum bool) schema.ColumnResolver {
	return func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		if isEnum {
			return r.Set(c.Name, funk.Get(r.Item, path, funk.WithAllowZero()).(fmt.Stringer).String())
		} else {
			return r.Set(c.Name, funk.Get(r.Item, path, funk.WithAllowZero()))
		}
	}
}

func (tg *TableGenerator) hasDefaultValue(path string) bool {
	_, ok := tg.defaultCols[path]
	return ok
}

func (tg *TableGenerator) isIgnored(path string) bool {
	_, ok := tg.ignoreCols[path]
	return ok
}

func isExpandable(field *desc.FieldDescriptor) bool {
	return !field.IsRepeated() && !field.IsMap() && field.GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE
}

func isConvertableToRelation(field protoField) bool {
	return field.IsRepeated() && !field.IsMap() && field.GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE
}

func getType(field protoField) schema.ValueType {
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
