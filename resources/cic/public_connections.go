package cic

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/cic/v1"
	cicsdk "github.com/yandex-cloud/go-sdk/services/cic/v1"
)

func PublicConnections() *schema.Table {
	return &schema.Table{
		Name:        "yc_cic_public_connections",
		Description: `https://yandex.cloud/ru/docs/interconnect/api-ref/grpc/PublicConnection/list#yandex.cloud.cic.v1.PublicConnection`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchPublicConnections,
		Transform:   client.TransformWithStruct(&cic.PublicConnection{}),
	}
}

func fetchPublicConnections(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := cicsdk.NewPublicConnectionClient(c.SDKv2).Iterator(ctx, &cic.ListPublicConnectionsRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
