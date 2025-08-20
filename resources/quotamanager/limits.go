package quotamanager

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/quotamanager/v1"
)

type QuotaLimit struct {
	Resource *quotamanager.Resource
	*quotamanager.QuotaLimit
}

func QuotaLimits() *schema.Table {
	return &schema.Table{
		Name:        "yc_quotamanager_quota_limits",
		Description: `https://yandex.cloud/ru/docs/quota-manager/api-ref/grpc/QuotaLimit/list#yandex.cloud.quotamanager.v1.QuotaLimit`,
		Resolver:    fetchQuotaLimits,
		Transform: client.TransformWithStruct(&QuotaLimit{},
			transformers.WithUnwrapStructFields("Resource", "QuotaLimit"),
			transformers.WithPrimaryKeys("Resource.Id", "QuotaId"),
		),
	}
}

func fetchQuotaLimits(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	service, ok := parent.Item.(*Service)
	if !ok {
		return fmt.Errorf("failed to extract service id from parent %T (%s)", parent.Item, parent.Item)
	}

	var joinedErr error
	// HACK: artificial child + multiplex: service (from parent) -> resource (from multiplex).
	// When using `schema.Tables/Relates`, the child table doesn't use multiplexing (feature), so we are doing it manually.
	// e.g. we want to resolve all `Service{Id: "vpc"}` quotas for `quotamanager.Resource{Id: "bg...", Type: "resource-manager.cloud"}`
	for _, cmeta := range client.CombineMultiplex(client.OrganizationMultiplex, client.CloudMultiplex)(meta) {
		c := cmeta.(*client.Client)

		resourceId := c.MultiplexedResourceId
		if resourceId == "" {
			return fmt.Errorf("client must be multiplexed by resource hierarchy entity (multiplexed resource id is empty)")
		}

		resourceType := c.MultiplexedResourceType
		if resourceType == "" {
			return fmt.Errorf("client must be multiplexed by resource hierarchy entity (multiplexed resource type is empty) ")
		}

		if slices.Index(quotaEnabledResourceTypes, resourceType) == -1 {
			return fmt.Errorf("%T (%s) is not a quota-enabled service", resourceType, resourceType)
		}

		it := c.SDK.QuotaManager().QuotaLimit().QuotaLimitIterator(ctx, &quotamanager.ListQuotaLimitsRequest{
			Resource: &quotamanager.Resource{
				Id:   resourceId,
				Type: string(resourceType),
			},
			Service: service.Id,
		})

		for it.Next() {
			ql := it.Value()

			res <- &QuotaLimit{
				Resource: &quotamanager.Resource{
					Id:   resourceId,
					Type: string(resourceType),
				},
				QuotaLimit: ql,
			}
		}

		if err := it.Error(); err != nil {
			// continue iterating because of artificial multiplex
			c.Logger.Err(err)
			joinedErr = errors.Join(joinedErr, err)
		}
	}

	return joinedErr
}
