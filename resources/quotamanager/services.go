// Provides a dictionary table of quota-enabled services per resource container.
package quotamanager

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/quotamanager/v1"
)

type Service struct {
	quotamanager.Service
	ResourceType string
}

func QuotaServices() *schema.Table {
	return &schema.Table{
		Name:        "yc_quotamanager_quota_services",
		Description: `https://yandex.cloud/ru/docs/quota-manager/api-ref/grpc/QuotaLimit/listServices#yandex.cloud.quotamanager.v1.Service`,
		Resolver:    fetchQuotaServices,
		Transform:   client.TransformWithStruct(&Service{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithPrimaryKeys("Id", "ResourceType")),
		Relations: schema.Tables{
			QuotaLimits(),
		},
	}
}

var quotaEnabledResourceTypes = []client.ResourceType{client.ResourceTypeOrganization, client.ResourceTypeCloud}

func fetchQuotaServices(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	for _, rt := range quotaEnabledResourceTypes {
		rt := string(rt)
		it := c.SDK.QuotaManager().QuotaLimit().QuotaLimitServicesIterator(ctx, &quotamanager.ListServicesRequest{
			ResourceType: rt,
		})
		for it.Next() {
			res <- &Service{
				Service:      *it.Value(),
				ResourceType: rt,
			}
		}

		if err := it.Error(); err != nil {
			return err
		}
	}

	return nil
}
