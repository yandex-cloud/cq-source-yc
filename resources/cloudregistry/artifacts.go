package cloudregistry

import (
	"context"
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/cloudregistry/v1"
)

func Artifacts() *schema.Table {
	return &schema.Table{
		Name:        "yc_cloudregistry_artifacts",
		Description: `https://yandex.cloud/ru/docs/cloud-registry/api-ref/grpc/Registry/listArtifacts#yandex.cloud.cloudregistry.v1.Artifact`,
		Resolver:    fetchArtifacts,
		Transform:   client.TransformWithStruct(&cloudregistry.Artifact{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
			schema.Column{
				Name:     "registry_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}

func fetchArtifacts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	registry, ok := parent.Item.(*cloudregistry.Registry)
	if !ok {
		return fmt.Errorf("parent in not type of *cloudregistry.Registry: %+v", registry)
	}

	it := c.SDK.CloudRegistry().Registry().RegistryArtifactsIterator(ctx, &cloudregistry.ListArtifactsRequest{RegistryId: registry.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
