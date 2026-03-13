package client

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/thoas/go-funk"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func ResolveProtoEnum(path string) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		data := funk.Get(resource.Item, path)
		if data == nil {
			return nil
		}
		enum, ok := data.(protoreflect.Enum)
		if !ok {
			return fmt.Errorf("unexpected type, wanted \"protoreflect.Enum\", have \"%T\"", data)
		}
		return resource.Set(c.Name, protoimpl.X.EnumStringOf(enum.Descriptor(), enum.Number()))
	}
}

func ResolveProtoTimestamp(path string) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		data := funk.Get(resource.Item, path)
		if data == nil {
			return nil
		}
		ts, ok := data.(*timestamppb.Timestamp)
		if !ok {
			return fmt.Errorf("unexpected type, wanted \"*timestamppb.Timestamp\", have \"%T\"", data)
		}
		return resource.Set(c.Name, ts.AsTime())
	}
}

// WrapperValue is a constraint for protobuf wrapper types that have a GetValue method
type WrapperValue[T any] interface {
	*wrapperspb.DoubleValue | *wrapperspb.FloatValue | *wrapperspb.StringValue |
		*wrapperspb.Int64Value | *wrapperspb.Int32Value | *wrapperspb.UInt64Value |
		*wrapperspb.UInt32Value | *wrapperspb.BoolValue | *wrapperspb.BytesValue
	GetValue() T
}

// ResolveWrapperValue is a generic resolver for protobuf wrapper types
func ResolveWrapperValue[T any, W WrapperValue[T]](path string) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		data := funk.Get(resource.Item, path)
		if data == nil {
			return nil
		}
		wrapper, ok := data.(W)
		if !ok {
			var zero W
			return fmt.Errorf("unexpected type, wanted \"%T\", have \"%T\"", zero, data)
		}
		return resource.Set(c.Name, wrapper.GetValue())
	}
}

var protomarshaler = protojson.MarshalOptions{
	UseEnumNumbers:    false,
	EmitUnpopulated:   false,
	EmitDefaultValues: false,

	// Important for snake_case
	UseProtoNames: true,
}

// ResolveOneofField resolves a protobuf oneof field using protoreflect.
//
// For flat paths (e.g. "NetworkImplementation"), it operates on resource.Item directly.
// For dotted paths (e.g. "Master.MasterType" from WithUnwrapStructFields),
// it navigates to the parent message via funk.Get before resolving the oneof.
//
// oneofName is the proto oneof name from the protobuf_oneof struct tag (e.g. "master_type").
// Produces {"active_field_name": <value>}.
func ResolveOneofField(path string, oneofName protoreflect.Name) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		var target any
		if i := strings.LastIndex(path, "."); i >= 0 {
			// Dotted path: navigate to the parent message that owns the oneof.
			target = funk.Get(resource.Item, path[:i])
		} else {
			target = resource.Item
		}
		if target == nil {
			return nil
		}

		msg, ok := target.(proto.Message)
		if !ok {
			return fmt.Errorf("expected proto.Message, got %T", target)
		}

		rv := msg.ProtoReflect()
		od := rv.Descriptor().Oneofs().ByName(oneofName)
		if od == nil {
			return fmt.Errorf("oneof %q not found in %s", oneofName, rv.Descriptor().FullName())
		}

		activeField := rv.WhichOneof(od)
		if activeField == nil {
			return nil
		}

		val := rv.Get(activeField)
		fieldName := string(activeField.Name())

		if activeField.Kind() == protoreflect.MessageKind || activeField.Kind() == protoreflect.GroupKind {
			jsonBytes, err := protomarshaler.Marshal(val.Message().Interface())
			if err != nil {
				return fmt.Errorf("failed to marshal oneof message: %w", err)
			}
			return resource.Set(c.Name, fmt.Appendf(nil, `{"%s":%s}`, fieldName, jsonBytes))
		}
		return resource.Set(c.Name, map[string]any{fieldName: val.Interface()})
	}
}

// ResolveProtoMessage creates a resolver that serializes protobuf messages using protojson
func ResolveProtoMessage(path string) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		data := funk.Get(resource.Item, path)
		if data == nil {
			return nil
		}

		// Check if data implements proto.Message
		msg, ok := data.(proto.Message)
		if !ok {
			return fmt.Errorf("unexpected type, wanted proto.Message, have %T", data)
		}

		jsonBytes, err := protomarshaler.Marshal(msg)
		if err != nil {
			return fmt.Errorf("failed to marshal proto message to JSON: %w", err)
		}

		return resource.Set(c.Name, jsonBytes)
	}
}

// ResolveProtoSlice resolves a slice of protobuf values by dispatching on the element type.
// It handles proto messages (→ JSON array), timestamps (→ []time.Time), enums (→ []string),
// and wrapper types (→ unwrapped slice) via runtime type inspection.
// TODO: remove code repetition
func ResolveProtoSlice(path string) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		data := funk.Get(resource.Item, path)
		if data == nil {
			return nil
		}
		val := reflect.ValueOf(data)
		if val.Kind() != reflect.Slice || val.Len() == 0 {
			return nil
		}

		first := val.Index(0).Interface()
		switch first.(type) {
		case *timestamppb.Timestamp:
			times := make([]time.Time, val.Len())
			for i := range val.Len() {
				times[i] = val.Index(i).Interface().(*timestamppb.Timestamp).AsTime()
			}
			return resource.Set(c.Name, times)

		case protoreflect.Enum:
			strs := make([]string, val.Len())
			for i := range val.Len() {
				e := val.Index(i).Interface().(protoreflect.Enum)
				strs[i] = protoimpl.X.EnumStringOf(e.Descriptor(), e.Number())
			}
			return resource.Set(c.Name, strs)

		case proto.Message:
			// Wrappers implement proto.Message but also have GetValue — unwrap them.
			if gv, ok := reflect.TypeOf(first).MethodByName("GetValue"); ok {
				out := reflect.MakeSlice(reflect.SliceOf(gv.Type.Out(0)), val.Len(), val.Len())
				for i := range val.Len() {
					out.Index(i).Set(reflect.ValueOf(val.Index(i).Interface()).MethodByName("GetValue").Call(nil)[0])
				}
				return resource.Set(c.Name, out.Interface())
			}
			// Regular proto messages — serialize as JSON array via protojson.
			var buf []byte
			buf = append(buf, '[')
			for i := range val.Len() {
				if i > 0 {
					buf = append(buf, ',')
				}
				jsonBytes, err := protomarshaler.Marshal(val.Index(i).Interface().(proto.Message))
				if err != nil {
					return fmt.Errorf("failed to marshal proto message at index %d: %w", i, err)
				}
				buf = append(buf, jsonBytes...)
			}
			buf = append(buf, ']')
			return resource.Set(c.Name, buf)

		default:
			return nil
		}
	}
}

var ParentIdColumn schema.Column = schema.Column{
	Name:       "id",
	Type:       arrow.BinaryTypes.String,
	Resolver:   schema.ParentColumnResolver("id"),
	PrimaryKey: true,
}

var MultiplexedResourceIdColumn schema.Column = schema.Column{
	Name: "id",
	Type: arrow.BinaryTypes.String,
	Resolver: func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		client := meta.(*Client)
		return resource.Set(c.Name, client.MultiplexedResourceId)
	},
	PrimaryKey: true,
}

func ResolveOrganization(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	client := meta.(*Client)
	return resource.Set(c.Name, client.OrganizationId)
}

func ResolveCloud(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	client := meta.(*Client)
	return resource.Set(c.Name, client.CloudId)
}

func ResolveFolder(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	client := meta.(*Client)
	return resource.Set(c.Name, client.FolderId)
}

var OrganiztionIdColumn schema.Column = schema.Column{
	Name:     "organization_id",
	Type:     arrow.BinaryTypes.String,
	Resolver: ResolveOrganization,
}

var CloudIdColumn schema.Column = schema.Column{
	Name:     "cloud_id",
	Type:     arrow.BinaryTypes.String,
	Resolver: ResolveCloud,
}

var FolderIdColumn schema.Column = schema.Column{
	Name:     "folder_id",
	Type:     arrow.BinaryTypes.String,
	Resolver: ResolveFolder,
}
