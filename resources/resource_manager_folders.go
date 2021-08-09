package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
)

func ResourceManagerFolders() *schema.Table {
	return &schema.Table{
		Name:        "yandex_resource_manager_folders",
		Resolver:    fetchResourceManagerFolders,
		Multiplex:   client.IdentityMultiplex,
		IgnoreError: client.IgnoreErrorHandler,
		Columns: []schema.Column{
			{
				Name:        "id",
				Type:        schema.TypeString,
				Description: "ID of the folder.",
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "cloud_id",
				Type:        schema.TypeString,
				Description: "ID of the cloud that the folder belongs to.",
				Resolver:    schema.PathResolver("CloudId"),
			},
			{
				Name:        "created_at",
				Type:        schema.TypeTimestamp,
				Description: "",
				Resolver:    client.ResolveAsTime,
			},
			{
				Name:        "name",
				Type:        schema.TypeString,
				Description: "Name of the folder.\n The name is unique within the cloud. 3-63 characters long.",
				Resolver:    schema.PathResolver("Name"),
			},
			{
				Name:        "description",
				Type:        schema.TypeString,
				Description: "Description of the folder. 0-256 characters long.",
				Resolver:    schema.PathResolver("Description"),
			},
			{
				Name:        "labels",
				Type:        schema.TypeJSON,
				Description: "Resource labels as `key:value` pairs. Maximum of 64 per resource.",
				Resolver:    client.ResolveLabels,
			},
			{
				Name:        "status",
				Type:        schema.TypeString,
				Description: "Status of the folder.",
				Resolver:    client.EnumPathResolver("Status"),
			},
		},
	}

}

func fetchResourceManagerFolders(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	req := &resourcemanager.ListFoldersRequest{CloudId: c.CloudId}
	it := c.Services.ResourceManager.Folder().FolderIterator(ctx, req)
	for it.Next() {
		res <- it.Value()
	}

	return nil
}
