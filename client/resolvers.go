package client

import (
	"context"
	"fmt"

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
// path is the Go struct field name of the oneof container (e.g. "NetworkImplementation").
// It is always a single, dot-free segment because the CQ SDK struct transformer only
// recurses into anonymous/embedded structs — proto message sub-fields are serialised as
// JSON blobs and never unwrapped, and oneof fields (interface types) are never structs,
// so the transformer never descends into them. See TestOneofPathsAreFlat.
//
// Previous implementation used reflect to unwrap the concrete oneof wrapper struct and
// parse the proto field name from struct tags. Using protoreflect + WhichOneof is cleaner:
// it works directly on the parent proto.Message (resource.Item) and obtains the active
// field descriptor — including its name — without any struct tag parsing.
//
// oneofName is the proto oneof name from the protobuf_oneof struct tag (e.g. "master_type").
// Produces {"active_field_name": <value>}.
func ResolveOneofField(path string, oneofName protoreflect.Name) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		msg, ok := resource.Item.(proto.Message)
		if !ok {
			return fmt.Errorf("expected proto.Message, got %T", resource.Item)
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
