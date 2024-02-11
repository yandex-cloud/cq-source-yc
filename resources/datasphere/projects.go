package datasphere

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/cq-source-yc/resources/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/datasphere/v2"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:        "yc_datasphere_projects",
		Description: `https://cloud.yandex.ru/docs/datasphere/api-ref/grpc/project_service#Project3`,
		Multiplex:   nil,
		Resolver:    fetchProjects,
		Transform:   client.TransformWithStruct(&datasphere.Project{}, client.PrimaryKeyIdTransformer),
		Relations:   schema.Tables{access.DatasphereProjectsBindings()},
	}
}

func fetchProjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	community, ok := parent.Item.(*datasphere.Community)
	if !ok {
		return fmt.Errorf("parent is not type of *datasphere.Community: %+v", community)
	}

	it := c.SDK.Datasphere().Project().ProjectIterator(ctx, &datasphere.ListProjectsRequest{CommunityId: community.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
