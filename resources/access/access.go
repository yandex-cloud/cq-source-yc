package access

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
)

var (
	Transform = client.TransformWithStruct(&access.AccessBinding{}, transformers.WithPrimaryKeys("Subject", "RoleId"))
)

func NewTable(name string, fetcher schema.TableResolver) *schema.Table {
	return &schema.Table{
		Name:      "yc_access_bindings_" + name,
		Resolver:  fetcher,
		Transform: Transform,
		Columns: schema.ColumnList{
			client.ParentIdColumn,
		},
	}
}
