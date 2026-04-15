package cloudregistry

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/cloudregistry/v1"
)

func IpPermissions() *schema.Table {
	return &schema.Table{
		Name:        "yc_cloudregistry_ip_permissions",
		Description: `https://yandex.cloud/ru/docs/cloud-registry/api-ref/grpc/Registry/listIpPermissions#yandex.cloud.cloudregistry.v1.IpPermission`,
		Resolver:    fetchIpPermissions,
		Transform:   client.TransformWithStruct(&cloudregistry.IpPermission{}),
		Columns: schema.ColumnList{
			client.ParentIdColumn,
		},
	}
}

func fetchIpPermissions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	registry, ok := parent.Item.(*cloudregistry.Registry)
	if !ok {
		return fmt.Errorf("parent in not type of *cloudregistry.Registry: %+v", registry)
	}

	it := c.SDK.CloudRegistry().Registry().RegistryIpPermissionsIterator(ctx, &cloudregistry.ListIpPermissionsRequest{RegistryId: registry.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
