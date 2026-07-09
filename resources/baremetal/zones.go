package baremetal

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/baremetal/v2"
	baremetalsdk "github.com/yandex-cloud/go-sdk/services/baremetal/v2"
)

// Zones is a global catalog of availability zones (no folder/cloud scope), so the
// table runs once.
func Zones() *schema.Table {
	return &schema.Table{
		Name:        "yc_baremetal_zones",
		Description: `https://yandex.cloud/docs/baremetal/api-ref/grpc/Zone/list#yandex.cloud.baremetal.v2.Zone`,
		Resolver:    fetchZones,
		Transform:   client.TransformWithStruct(&baremetal.Zone{}, client.PrimaryKeyIdTransformers("ZoneId")...),
	}
}

func fetchZones(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	it := baremetalsdk.NewZoneClient(c.SDKv2).ZonesIterator(ctx, &baremetal.ListZonesRequest{})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
