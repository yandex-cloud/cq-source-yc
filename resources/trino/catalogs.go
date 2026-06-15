package trino

import (
	"context"
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/trino/v1"
)

func Catalogs() *schema.Table {
	return &schema.Table{
		Name:        "yc_trino_catalogs",
		Description: `https://yandex.cloud/docs/managed-trino/api-ref/grpc/Catalog/list#yandex.cloud.trino.v1.Catalog`,
		Resolver:    fetchCatalogs,
		Transform:   client.TransformWithStruct(&trino.Catalog{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
			schema.Column{
				Name:     "cluster_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}

func fetchCatalogs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	cluster, ok := parent.Item.(*trino.Cluster)
	if !ok {
		return fmt.Errorf("parent is not type of *trino.Cluster: %+v", cluster)
	}

	it := c.SDK.Trino().Catalog().CatalogIterator(ctx, &trino.ListCatalogsRequest{ClusterId: cluster.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
