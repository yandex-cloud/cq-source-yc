package access

import (
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
)

// AccessPolicyTransform builds columns from access.AccessPolicyBinding — the IAM
// "access policy" bindings (templated policies), distinct from plain role AccessBindings.
// https://yandex.cloud/docs/iam/concepts/access-control/#access-policies
var AccessPolicyTransform = client.TransformWithStruct(
	&access.AccessPolicyBinding{},
	transformers.WithPrimaryKeys("AccessPolicyTemplateId"),
)
