package cic

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/cic/v1"
	cicsdk "github.com/yandex-cloud/go-sdk/services/cic/v1"
)

func PrivateConnections() *schema.Table {
	return &schema.Table{
		Name:        "yc_cic_private_connections",
		Description: `https://yandex.cloud/ru/docs/interconnect/api-ref/grpc/PrivateConnection/list#yandex.cloud.cic.v1.PrivateConnection`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchPrivateConnections,
		Transform:   client.TransformWithStruct(&cic.PrivateConnection{}),
	}
}

func fetchPrivateConnections(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := cicsdk.NewPrivateConnectionClient(c.SDKv2).Iterator(ctx, &cic.ListPrivateConnectionsRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
