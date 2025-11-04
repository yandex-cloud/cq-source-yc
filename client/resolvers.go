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

		// Serialize using protojson with enums as strings
		marshaler := protojson.MarshalOptions{
			UseEnumNumbers:    false,
			EmitDefaultValues: true,
			UseProtoNames:     true,
		}

		jsonBytes, err := marshaler.Marshal(msg)
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
