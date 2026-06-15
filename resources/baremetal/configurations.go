package baremetal

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	baremetal "github.com/yandex-cloud/go-genproto/yandex/cloud/baremetal/v1alpha"
	baremetalsdk "github.com/yandex-cloud/go-sdk/services/baremetal/v1alpha"
)

// Configurations is a global catalog of available baremetal server
// configurations (no folder/cloud scope), so the table runs once.
func Configurations() *schema.Table {
	return &schema.Table{
		Name:        "yc_baremetal_configurations",
		Description: `https://yandex.cloud/docs/baremetal/api-ref/grpc/Configuration/list#yandex.cloud.baremetal.v1alpha.Configuration`,
		Resolver:    fetchConfigurations,
		Transform:   client.TransformWithStruct(&baremetal.Configuration{}, client.PrimaryKeyIdTransformer),
	}
}

func fetchConfigurations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	it := baremetalsdk.NewConfigurationClient(c.SDKv2).Iterator(ctx, &baremetal.ListConfigurationsRequest{})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
