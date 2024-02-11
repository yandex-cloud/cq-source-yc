package access

import (
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
)

var (
	Transform = client.TransformWithStruct(&access.AccessBinding{}, transformers.WithPrimaryKeys("Subject", "RoleId"))
)
