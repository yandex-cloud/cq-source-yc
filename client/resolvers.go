package client

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/thoas/go-funk"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		client.Logger.Debug().Str("MultiplexedResourceId", client.MultiplexedResourceId)

		return resource.Set(c.Name, client.MultiplexedResourceId)
	},
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
