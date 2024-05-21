package client

import (
	"reflect"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	cqtypes "github.com/cloudquery/plugin-sdk/v4/types"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func typeTransformer(field reflect.StructField) (arrow.DataType, error) {
	switch reflect.New(field.Type).Elem().Interface().(type) {
	case *timestamppb.Timestamp,
		timestamppb.Timestamp:
		return arrow.FixedWidthTypes.Timestamp_us, nil
	case protoreflect.Enum:
		return arrow.BinaryTypes.String, nil
	case nil:
		return cqtypes.NewJSONType(), nil
	default:
		return nil, nil
	}
}

func resolverTransformer(field reflect.StructField, path string) schema.ColumnResolver {
	switch reflect.New(field.Type).Elem().Interface().(type) {
	case *timestamppb.Timestamp,
		timestamppb.Timestamp:
		return ResolveProtoTimestamp(path)
	case protoreflect.Enum:
		return ResolveProtoEnum(path)
	default:
		return nil
	}

}

var options = []transformers.StructTransformerOption{
	transformers.WithTypeTransformer(typeTransformer),
	transformers.WithResolverTransformer(resolverTransformer),
}

func SharedTransformers() []transformers.StructTransformerOption {
	return options
}

func TransformWithStruct(t any, opts ...transformers.StructTransformerOption) schema.Transform {
	return transformers.TransformWithStruct(t, append(options, opts...)...)
}

var PrimaryKeyIdTransformer transformers.StructTransformerOption = transformers.WithPrimaryKeys(("Id"))
