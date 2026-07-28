package baremetal

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/baremetal/v2"
	baremetalsdk "github.com/yandex-cloud/go-sdk/services/baremetal/v2"
)

// HardwarePools is a global catalog (no folder/cloud scope), so the table runs once.
func HardwarePools() *schema.Table {
	return &schema.Table{
		Name:        "yc_baremetal_hardware_pools",
		Description: `https://yandex.cloud/docs/baremetal/api-ref/grpc/HardwarePool/list#yandex.cloud.baremetal.v2.HardwarePool`,
		Multiplex:   client.GlobalMultiplex(client.ServiceBaremetal),
		Resolver:    fetchHardwarePools,
		Transform:   client.TransformWithStruct(&baremetal.HardwarePool{}, client.PrimaryKeyIdTransformers("HardwarePoolId")...),
	}
}

func fetchHardwarePools(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	it := baremetalsdk.NewHardwarePoolClient(c.SDKv2).HardwarePoolsIterator(ctx, &baremetal.ListHardwarePoolsRequest{})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
