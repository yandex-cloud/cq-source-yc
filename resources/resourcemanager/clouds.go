package resourcemanager

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
)

func Clouds() *schema.Table {
	return &schema.Table{
		Name:        "yc_resourcemanager_clouds",
		Description: `https://cloud.yandex.ru/docs/resource-manager/api-ref/grpc/cloud_service#Cloud1`,
		Multiplex:   client.CloudMultiplex,
		Resolver:    fetchClouds,
		Transform:   client.TransformWithStruct(&resourcemanager.Cloud{}, client.PrimaryKeyIdTransformer),
	}
}

func fetchClouds(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	cloud, err := c.SDK.ResourceManager().Cloud().Get(ctx, &resourcemanager.GetCloudRequest{CloudId: c.CloudId})
	if err != nil {
		return err
	}
	res <- cloud

	return nil
}
