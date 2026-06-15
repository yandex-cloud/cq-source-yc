package access

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
)

func CloudsAccessPolicyBindings() *schema.Table {
	return &schema.Table{
		Name:        "yc_access_policy_bindings_resourcemanager_clouds",
		Title:       "YC Access Policy Bindings for Clouds",
		Description: `https://yandex.cloud/docs/iam/concepts/access-control/#access-policies`,
		Multiplex:   client.CloudMultiplex,
		Resolver:    fetchCloudsAccessPolicyBindings,
		Transform:   AccessPolicyTransform,
		Columns: schema.ColumnList{
			client.MultiplexedResourceIdColumn,
		},
	}
}

func fetchCloudsAccessPolicyBindings(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	it := c.SDK.ResourceManager().Cloud().CloudAccessPolicyBindingsIterator(ctx, &access.ListAccessPolicyBindingsRequest{ResourceId: c.CloudId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
