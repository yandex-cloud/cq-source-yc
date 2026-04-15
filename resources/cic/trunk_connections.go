package cic

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/cic/v1"
	cicsdk "github.com/yandex-cloud/go-sdk/services/cic/v1"
)

func TrunkConnections() *schema.Table {
	return &schema.Table{
		Name:        "yc_cic_trunk_connections",
		Description: `https://yandex.cloud/ru/docs/interconnect/api-ref/grpc/TrunkConnection/list#yandex.cloud.cic.v1.TrunkConnection`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchTrunkConnections,
		Transform:   client.TransformWithStruct(&cic.TrunkConnection{}),
	}
}

func fetchTrunkConnections(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := cicsdk.NewTrunkConnectionClient(c.SDKv2).Iterator(ctx, &cic.ListTrunkConnectionsRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
