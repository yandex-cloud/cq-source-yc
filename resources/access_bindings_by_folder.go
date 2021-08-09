package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
)

func AccessBindingsByFolder() *schema.Table {
	return &schema.Table{
		Name:         "yandex_access_bindings_by_folder",
		Resolver:     fetchAccessBindingsByFolder,
		Multiplex:    client.FolderMultiplex,
		IgnoreError:  client.IgnoreErrorHandler,
		DeleteFilter: client.DeleteFolderFilter,
		Columns: []schema.Column{
			{
				Name:     "folder_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FolderOrCloudId"),
			},
			{
				Name:     "role_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RoleId"),
			},
			{
				Name:     "subject_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Subject.Id"),
			},
			{
				Name:     "subject_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Subject.Type"),
			},
		},
	}

}

type RichAccessBindingByFolder struct {
	*access.AccessBinding
	FolderId string
}

func fetchAccessBindingsByFolder(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	resp, err := c.Services.ResourceManager.Folder().ListAccessBindings(ctx, &access.ListAccessBindingsRequest{
		ResourceId: c.FolderId,
	})

	if err != nil {
		return err
	}

	for {
		for _, value := range resp.GetAccessBindings() {
			res <- RichAccessBindingByFolder{value, c.FolderId}
		}

		if resp.GetNextPageToken() == "" {
			break
		}

		resp, err = c.Services.ResourceManager.Folder().ListAccessBindings(ctx, &access.ListAccessBindingsRequest{
			ResourceId: c.FolderId,
			PageToken:  resp.GetNextPageToken(),
		})

		if err != nil {
			return err
		}
	}

	return nil
}
