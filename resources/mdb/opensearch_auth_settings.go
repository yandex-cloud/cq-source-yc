package mdb

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/opensearch/v1"
)

func OpenSearchAuthSettings() *schema.Table {
	return &schema.Table{
		Name:        "yc_mdb_opensearch_auth_settings",
		Description: `https://cloud.yandex.ru/docs/managed-opensearch/api-ref/grpc/cluster_service#AuthSettings`,
		Resolver:    fetchOpenSearchAuthSettings,
		Transform:   client.TransformWithStruct(&opensearch.AuthSettings{}, transformers.WithUnwrapStructFields("Saml")),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
			clusterIdColumn,
		},
	}
}

func fetchOpenSearchAuthSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	cluster, ok := parent.Item.(*opensearch.Cluster)
	if !ok {
		return fmt.Errorf("parent is not type of *opensearch.Cluster: %+v", cluster)
	}

	resp, err := c.SDK.MDB().OpenSearch().Cluster().GetAuthSettings(ctx, &opensearch.GetAuthSettingsRequest{ClusterId: cluster.Id})
	if err != nil {
		return err
	}
	res <- resp

	return nil
}
