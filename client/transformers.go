package client

import (
	"reflect"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	cqtypes "github.com/cloudquery/plugin-sdk/v4/types"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func typeTransformer(field reflect.StructField) (arrow.DataType, error) {
	switch reflect.New(field.Type).Elem().Interface().(type) {
	case *timestamppb.Timestamp,
		timestamppb.Timestamp:
		return arrow.FixedWidthTypes.Timestamp_us, nil
	case *protoreflect.Enum,
		protoreflect.Enum:
		return arrow.BinaryTypes.String, nil
	case *wrapperspb.DoubleValue,
		wrapperspb.DoubleValue:
		return arrow.PrimitiveTypes.Float64, nil
	case *wrapperspb.FloatValue,
		wrapperspb.FloatValue:
		return arrow.PrimitiveTypes.Float32, nil
	case *wrapperspb.StringValue,
		wrapperspb.StringValue:
		return arrow.BinaryTypes.String, nil
	case *wrapperspb.Int64Value,
		wrapperspb.Int64Value:
		return arrow.PrimitiveTypes.Int64, nil
	case *wrapperspb.Int32Value,
		wrapperspb.Int32Value:
		return arrow.PrimitiveTypes.Int32, nil
	case *wrapperspb.UInt64Value,
		wrapperspb.UInt64Value:
		return arrow.PrimitiveTypes.Uint64, nil
	case *wrapperspb.UInt32Value,
		wrapperspb.UInt32Value:
		return arrow.PrimitiveTypes.Uint32, nil
	case *wrapperspb.BoolValue,
		wrapperspb.BoolValue:
		return arrow.FixedWidthTypes.Boolean, nil
	case *wrapperspb.BytesValue,
		wrapperspb.BytesValue:
		return arrow.BinaryTypes.Binary, nil
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
	case *protoreflect.Enum,
		protoreflect.Enum:
		return ResolveProtoEnum(path)
	case *wrapperspb.DoubleValue,
		wrapperspb.DoubleValue:
		return ResolveWrapperValue[float64, *wrapperspb.DoubleValue](path)
	case *wrapperspb.FloatValue,
		wrapperspb.FloatValue:
		return ResolveWrapperValue[float32, *wrapperspb.FloatValue](path)
	case *wrapperspb.StringValue,
		wrapperspb.StringValue:
		return ResolveWrapperValue[string, *wrapperspb.StringValue](path)
	case *wrapperspb.Int64Value,
		wrapperspb.Int64Value:
		return ResolveWrapperValue[int64, *wrapperspb.Int64Value](path)
	case *wrapperspb.Int32Value,
		wrapperspb.Int32Value:
		return ResolveWrapperValue[int32, *wrapperspb.Int32Value](path)
	case *wrapperspb.UInt64Value,
		wrapperspb.UInt64Value:
		return ResolveWrapperValue[uint64, *wrapperspb.UInt64Value](path)
	case *wrapperspb.UInt32Value,
		wrapperspb.UInt32Value:
		return ResolveWrapperValue[uint32, *wrapperspb.UInt32Value](path)
	case *wrapperspb.BoolValue,
		wrapperspb.BoolValue:
		return ResolveWrapperValue[bool, *wrapperspb.BoolValue](path)
	case *wrapperspb.BytesValue,
		wrapperspb.BytesValue:
		return ResolveWrapperValue[[]byte, *wrapperspb.BytesValue](path)
	case *proto.Message,
		proto.Message:
		return ResolveProtoMessage(path)
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

func TransformColumnPrimaryKey(column schema.Column) schema.Column {
	col := column
	col.PrimaryKey = true
	return col
}
