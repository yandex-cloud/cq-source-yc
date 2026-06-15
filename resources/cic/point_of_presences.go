package cic

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/cic/v1"
	cicsdk "github.com/yandex-cloud/go-sdk/services/cic/v1"
)

func PointOfPresences() *schema.Table {
	return &schema.Table{
		Name:        "yc_cic_point_of_presences",
		Description: `https://yandex.cloud/docs/interconnect/api-ref/grpc/PointOfPresence/list#yandex.cloud.cic.v1.PointOfPresence`,
		Resolver:    fetchPointOfPresences,
		Transform:   client.TransformWithStruct(&cic.PointOfPresence{}, client.PrimaryKeyIdTransformer),
	}
}

func fetchPointOfPresences(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	it := cicsdk.NewPointOfPresenceClient(c.SDKv2).Iterator(ctx, &cic.ListPointOfPresencesRequest{})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
