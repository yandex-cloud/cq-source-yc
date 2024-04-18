package audittrails

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/audittrails/v1"
)

func Trails() *schema.Table {
	return &schema.Table{
		Name:        "yc_audittrails_trails",
		Description: `https://yandex.cloud/ru/docs/audit-trails/api-ref/grpc/trail_service#Trail1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchTrails,
		Transform:   client.TransformWithStruct(&audittrails.Trail{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchTrails(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.AuditTrails().Trail().TrailIterator(ctx, &audittrails.ListTrailsRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
