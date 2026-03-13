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

const protobufOneofTag = "protobuf_oneof"

// protoTypeToArrow recursively resolves a reflect.Type to an Arrow DataType.
// It peels pointers and slices, then matches known protobuf types.
func protoTypeToArrow(t reflect.Type) (arrow.DataType, error) {
	switch t.Kind() {
	case reflect.Pointer:
		return protoTypeToArrow(t.Elem())
	case reflect.Slice:
		if t.Elem().Kind() == reflect.Uint8 {
			return nil, nil // []byte
		}
		elemType, err := protoTypeToArrow(t.Elem())
		if err != nil {
			return nil, err
		}
		// Not a protobuf type
		if elemType == nil {
			return nil, nil
		}
		if elemType == cqtypes.ExtensionTypes.JSON {
			return cqtypes.ExtensionTypes.JSON, nil
		}
		return arrow.ListOf(elemType), nil
	case reflect.Interface:
		return nil, nil // oneof containers — handled by tag check
	}

	switch reflect.New(t).Interface().(type) {
	case *timestamppb.Timestamp:
		return arrow.FixedWidthTypes.Timestamp_us, nil
	case protoreflect.Enum:
		return arrow.BinaryTypes.String, nil
	case *wrapperspb.DoubleValue:
		return arrow.PrimitiveTypes.Float64, nil
	case *wrapperspb.FloatValue:
		return arrow.PrimitiveTypes.Float32, nil
	case *wrapperspb.StringValue:
		return arrow.BinaryTypes.String, nil
	case *wrapperspb.Int64Value:
		return arrow.PrimitiveTypes.Int64, nil
	case *wrapperspb.Int32Value:
		return arrow.PrimitiveTypes.Int32, nil
	case *wrapperspb.UInt64Value:
		return arrow.PrimitiveTypes.Uint64, nil
	case *wrapperspb.UInt32Value:
		return arrow.PrimitiveTypes.Uint32, nil
	case *wrapperspb.BoolValue:
		return arrow.FixedWidthTypes.Boolean, nil
	case *wrapperspb.BytesValue:
		return arrow.BinaryTypes.Binary, nil
	case proto.Message:
		return cqtypes.ExtensionTypes.JSON, nil
	default:
		return nil, nil
	}
}

func TypeTransformer(field reflect.StructField) (arrow.DataType, error) {
	if _, ok := field.Tag.Lookup(protobufOneofTag); ok {
		return cqtypes.ExtensionTypes.JSON, nil
	}
	return protoTypeToArrow(field.Type)
}

// protoResolver recursively resolves a reflect.Type to a ColumnResolver.
// It peels pointers and slices, then matches known protobuf types.
func protoResolver(t reflect.Type, path string) schema.ColumnResolver {
	switch t.Kind() {
	case reflect.Pointer:
		return protoResolver(t.Elem(), path)
	case reflect.Slice:
		if t.Elem().Kind() == reflect.Uint8 {
			return nil // []byte
		}
		// Only return a slice resolver when protoTypeToArrow recognises the element type.
		if elemType, _ := protoTypeToArrow(t.Elem()); elemType != nil {
			return ResolveProtoSlice(path)
		}
		return nil
	case reflect.Interface:
		return nil
	}

	switch reflect.New(t).Interface().(type) {
	case *timestamppb.Timestamp:
		return ResolveProtoTimestamp(path)
	case protoreflect.Enum:
		return ResolveProtoEnum(path)
	case *wrapperspb.DoubleValue:
		return ResolveWrapperValue[float64, *wrapperspb.DoubleValue](path)
	case *wrapperspb.FloatValue:
		return ResolveWrapperValue[float32, *wrapperspb.FloatValue](path)
	case *wrapperspb.StringValue:
		return ResolveWrapperValue[string, *wrapperspb.StringValue](path)
	case *wrapperspb.Int64Value:
		return ResolveWrapperValue[int64, *wrapperspb.Int64Value](path)
	case *wrapperspb.Int32Value:
		return ResolveWrapperValue[int32, *wrapperspb.Int32Value](path)
	case *wrapperspb.UInt64Value:
		return ResolveWrapperValue[uint64, *wrapperspb.UInt64Value](path)
	case *wrapperspb.UInt32Value:
		return ResolveWrapperValue[uint32, *wrapperspb.UInt32Value](path)
	case *wrapperspb.BoolValue:
		return ResolveWrapperValue[bool, *wrapperspb.BoolValue](path)
	case *wrapperspb.BytesValue:
		return ResolveWrapperValue[[]byte, *wrapperspb.BytesValue](path)
	case proto.Message:
		return ResolveProtoMessage(path)
	default:
		return nil
	}
}

// ResolverTransformer returns a custom ColumnResolver for protobuf struct fields.
//
// Paths may contain dots when WithUnwrapStructFields is used (e.g. "CloudStatus.Id").
// All resolvers use funk.Get which supports dotted paths natively.
// ResolveOneofField navigates to the parent message for dotted paths before
// applying protoreflect-based oneof resolution.
func ResolverTransformer(field reflect.StructField, path string) schema.ColumnResolver {
	if oneofName, ok := field.Tag.Lookup(protobufOneofTag); ok {
		return ResolveOneofField(path, protoreflect.Name(oneofName))
	}
	return protoResolver(field.Type, path)
}

var options = []transformers.StructTransformerOption{
	transformers.WithTypeTransformer(TypeTransformer),
	transformers.WithResolverTransformer(ResolverTransformer),
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
