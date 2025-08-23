package quotamanager

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/quotamanager/v1"
)

type servicesCacher struct {
	responses sync.Map
}

func (sc *servicesCacher) Get(ctx context.Context, client *client.Client, resourceType client.ResourceType) ([]*quotamanager.Service, error) {
	if val, ok := sc.responses.Load(resourceType); ok {
		return val.([]*quotamanager.Service), nil
	}

	result, err := client.SDK.QuotaManager().QuotaLimit().ListServices(ctx, &quotamanager.ListServicesRequest{
		ResourceType: string(resourceType),
		// HACK: if services > 1000 this would work incorrectly
		PageSize: 1000,
	})
	if err != nil {
		return nil, err
	}

	sc.responses.Store(resourceType, result.Services)
	return result.Services, nil
}

var cacher = servicesCacher{}

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
		Multiplex: client.CombineMultiplex(client.OrganizationMultiplex, client.CloudMultiplex),
	}
}

func fetchQuotaLimits(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	if c.MultiplexedResourceId == "" {
		return fmt.Errorf("client must be multiplexed by resource hierarchy entity (multiplexed resource id is empty)")
	}
	if c.MultiplexedResourceType == "" {
		return fmt.Errorf("client must be multiplexed by resource hierarchy entity (multiplexed resource type is empty) ")
	}

	// HACK: artificial (resource, service) multiplex: (Organizations + Clouds) ✕ service (from API call).
	// Bad concurrency as CQ SDK limits apply per multiplexer entity – resource in our case,
	// so concurrency key would be (resource), hence if services count is big, it would be slow

	// Don't want to call rarely-changing response API per each org or cloud
	services, err := cacher.Get(ctx, c, c.MultiplexedResourceType)
	if err != nil {
		return fmt.Errorf("failed to fetch quota-enabled services for %s", c.MultiplexedResourceType)
	}

	resource := &quotamanager.Resource{
		Id:   c.MultiplexedResourceId,
		Type: string(c.MultiplexedResourceType),
	}

	var joinedErr error
	for _, service := range services {
		it := c.SDK.QuotaManager().QuotaLimit().QuotaLimitIterator(ctx, &quotamanager.ListQuotaLimitsRequest{
			Resource: resource,
			Service:  service.Id,
		})

		for it.Next() {
			res <- &QuotaLimit{
				Resource:   resource,
				QuotaLimit: it.Value(),
			}
		}

		if err := it.Error(); err != nil {
			// continue iterating because of artificial multiplex
			joinedErr = errors.Join(joinedErr, err)
		}
	}

	return joinedErr
}
