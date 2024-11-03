package privatelink

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1/privatelink"
)

func PrivateEndpoints() *schema.Table {
	return &schema.Table{
		Name:        "yc_vpc_privatelink_private_endpoints",
		Title:       "VPC Private Endpoints",
		Description: `https://yandex.cloud/ru/docs/vpc/privatelink/api-ref/grpc/PrivateEndpoint/list#yandex.cloud.vpc.v1.privatelink.PrivateEndpoint`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchPrivateEndpoints,
		Transform:   client.TransformWithStruct(&privatelink.PrivateEndpoint{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchPrivateEndpoints(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.VPCPrivateLink().PrivateEndpoint().PrivateEndpointIterator(ctx, &privatelink.ListPrivateEndpointsRequest{
		Container: &privatelink.ListPrivateEndpointsRequest_FolderId{
			FolderId: c.FolderId,
		},
	})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
