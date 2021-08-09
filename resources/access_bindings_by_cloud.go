package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
)

func AccessBindingsByCloud() *schema.Table {
	return &schema.Table{
		Name:        "yandex_access_bindings_by_cloud",
		Resolver:    fetchAccessBindingsByCloud,
		Multiplex:   client.IdentityMultiplex,
		IgnoreError: client.IgnoreErrorHandler,
		Columns: []schema.Column{
			{
				Name:     "cloud_id",
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

type RichAccessBindingByCloud struct {
	*access.AccessBinding
	CloudId string
}

func fetchAccessBindingsByCloud(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	resp, err := c.Services.ResourceManager.Cloud().ListAccessBindings(ctx, &access.ListAccessBindingsRequest{
		ResourceId: c.CloudId,
	})

	if err != nil {
		return err
	}

	for {
		for _, value := range resp.GetAccessBindings() {
			res <- RichAccessBindingByCloud{value, c.CloudId}
		}

		if resp.GetNextPageToken() == "" {
			break
		}

		resp, err = c.Services.ResourceManager.Cloud().ListAccessBindings(ctx, &access.ListAccessBindingsRequest{
			ResourceId: c.CloudId,
			PageToken:  resp.GetNextPageToken(),
		})

		if err != nil {
			return err
		}
	}

	return nil
}
