package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	"google.golang.org/grpc"
)

func AccessBindings() *schema.Table {
	return &schema.Table{
		Name:        "yandex_access_bindings",
		Resolver:    fetchAccessBindings,
		Multiplex:   client.FolderAndCloudMultiplex,
		IgnoreError: client.IgnoreErrorHandler,
		Columns: []schema.Column{
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceId"),
			},
			{
				Name:     "resource_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceType"),
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

type accessBindingsLister interface {
	ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error)
}

type RichAccessBinding struct {
	*access.AccessBinding
	ResourceId   string
	ResourceType string
}

func fetchAccessBindings(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	var (
		lister       accessBindingsLister
		resourceId   string
		resourceType string
	)
	switch {
	case c.FolderId != "":
		lister = c.Services.ResourceManager.Folder()
		resourceId = c.FolderId
		resourceType = "folder"
	case c.CloudId != "":
		lister = c.Services.ResourceManager.Cloud()
		resourceId = c.CloudId
		resourceType = "cloud"
	}

	resp, err := lister.ListAccessBindings(ctx, &access.ListAccessBindingsRequest{
		ResourceId: resourceId,
	})

	if err != nil {
		return err
	}

	for {
		for _, value := range resp.GetAccessBindings() {
			res <- RichAccessBinding{
				AccessBinding: value,
				ResourceId:    resourceId,
				ResourceType:  resourceType,
			}
		}

		if resp.GetNextPageToken() == "" {
			break
		}

		resp, err = lister.ListAccessBindings(ctx, &access.ListAccessBindingsRequest{
			ResourceId: resourceId,
			PageToken:  resp.GetNextPageToken(),
		})

		if err != nil {
			return err
		}
	}

	return nil
}
