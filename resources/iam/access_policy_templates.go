package iam

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
)

// AccessPolicyTemplates lists the IAM access policy templates catalog. The List
// RPC is global (no folder/cloud/org scope), so the table has no Multiplex and
// runs once. Each template is returned as an access.AccessPolicy.
func AccessPolicyTemplates() *schema.Table {
	return &schema.Table{
		Name:        "yc_iam_access_policy_templates",
		Description: `https://yandex.cloud/docs/iam/concepts/access-control/#access-policies`,
		Resolver:    fetchAccessPolicyTemplates,
		Transform:   client.TransformWithStruct(&access.AccessPolicy{}, client.PrimaryKeyIdTransformer),
	}
}

func fetchAccessPolicyTemplates(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	it := c.SDK.IAM().AccessPolicyTemplate().AccessPolicyTemplateIterator(ctx, &iam.ListAccessPolicyTemplatesRequest{})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
