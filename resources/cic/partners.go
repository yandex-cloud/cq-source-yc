package cic

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/cic/v1"
	cicsdk "github.com/yandex-cloud/go-sdk/services/cic/v1"
)

func Partners() *schema.Table {
	return &schema.Table{
		Name:        "yc_cic_partners",
		Description: `https://yandex.cloud/docs/interconnect/api-ref/grpc/Partner/list#yandex.cloud.cic.v1.Partner`,
		Resolver:    fetchPartners,
		Transform:   client.TransformWithStruct(&cic.Partner{}, client.PrimaryKeyIdTransformer),
	}
}

func fetchPartners(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	it := cicsdk.NewPartnerClient(c.SDKv2).Iterator(ctx, &cic.ListPartnersRequest{})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
