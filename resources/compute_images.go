package resources

import (
	"context"

	"github.com/GennadySpb/cq-provider-yandex/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
)

func ComputeImages() *schema.Table {
	return &schema.Table{
		Name:         "yandex_compute_images",
		Resolver:     fetchComputeImages,
		Multiplex:    client.FolderMultiplex,
		IgnoreError:  client.IgnoreErrorHandler,
		DeleteFilter: client.DeleteFolderFilter,
		Columns: []schema.Column{
			{
				Name:     "folder_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveFolderID,
			},
			{
				Name: "created_at",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "image_id",
				Type: schema.TypeString,
			},
			{
				Name: "image_type",
				Type: schema.TypeString,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: resolveComputeImageLabels,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeImages(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	// TODO: iterate over all  folders ???
	locations := []string{c.FolderId}

	for _, f := range locations {
		req := &compute.ListImagesRequest{FolderId: f}
		it := c.Services.Compute.Image().ImageIterator(ctx, req)
		for it.Next() {
			res <- it.Value()
		}
	}
	return nil
}

func resolveComputeImageLabels(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*compute.Image)
	labels := map[string]*string{}
	for k, v := range r.Labels {
		labels[k] = &v
	}
	return resource.Set("labels", labels)
}
