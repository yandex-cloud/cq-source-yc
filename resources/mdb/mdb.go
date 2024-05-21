package mdb

import (
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yandex-cloud/cq-source-yc/client"
)

var (
	primaryKeyNameClusterIdTransformer = transformers.WithPrimaryKeys("Name", "ClusterId")
	clusterIdResolver                  = schema.ParentColumnResolver("Id")
	clusterIdColumn                    = schema.Column{
		Name:       "cluster_id",
		Type:       arrow.BinaryTypes.String,
		Resolver:   clusterIdResolver,
		PrimaryKey: true,
	}
)

func structNameClusterIdTransformer(t any) schema.Transform {
	return client.TransformWithStruct(t, primaryKeyNameClusterIdTransformer)
}
