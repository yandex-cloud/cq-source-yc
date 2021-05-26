package resources

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"

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
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: resolveComputeImageCreatedAt,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name:     "image_id",
				Type:     schema.TypeString,
				Resolver: resolveComputeImageID,
			},
			{
				Name: "family",
				Type: schema.TypeString,
			},
			{
				Name:     "os_type",
				Type:     schema.TypeString,
				Resolver: resolveComputeImageOsType,
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: resolveComputeImageLabels,
			},
			{
				Name: "product_ids",
				Type: schema.TypeStringArray,
			},
			{
				Name: "min_disk_size",
				Type: schema.TypeBigInt,
			},
			{
				Name: "storage_size",
				Type: schema.TypeBigInt,
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

func resolveComputeImageCreatedAt(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*compute.Image)
	ts, _ := getTime(r.GetCreatedAt())
	return resource.Set("created_at", ts)
}

func getTime(protots *timestamp.Timestamp) (*time.Time, error) {
	if protots == nil {
		return nil, nil
	}
	if !protots.IsValid() {
		return nil, fmt.Errorf("invalid proto timestamp")

	}

	ts := protots.AsTime()
	return &ts, nil
}

func resolveComputeImageID(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*compute.Image)
	return resource.Set("image_id", r.GetId())
}

func resolveComputeImageOsType(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*compute.Image)
	return resource.Set("os_type", r.GetOs().GetType().String())
}
