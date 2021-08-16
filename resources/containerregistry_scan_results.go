package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/containerregistry/v1"
	"golang.org/x/sync/errgroup"
)

func ContainerRegistryScanResults() *schema.Table {
	return &schema.Table{
		Name:        "yandex_containerregistry_scan_results",
		Resolver:    fetchContainerRegistryScanResults,
		Multiplex:   client.MultiplexBy(client.Folders),
		IgnoreError: client.IgnoreErrorHandler,
		Columns: []schema.Column{
			{
				Name:            "id",
				Type:            schema.TypeString,
				Description:     "ID of the resource.",
				Resolver:        client.ResolveResourceId,
				CreationOptions: schema.ColumnCreationOptions{Nullable: false, Unique: true},
			},
			{
				Name:        "image_id",
				Type:        schema.TypeString,
				Description: "Output only. ID of the Image that the ScanResult belongs to.",
				Resolver:    schema.PathResolver("ImageId"),
			},
			{
				Name:        "scanned_at",
				Type:        schema.TypeTimestamp,
				Description: "Output only. The timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format when the scan been finished.",
				Resolver:    client.ResolveAsTime,
			},
			{
				Name:        "status",
				Type:        schema.TypeString,
				Description: "Output only. The status of the ScanResult.",
				Resolver:    client.EnumPathResolver("Status"),
			},
			{
				Name:        "vulnerabilities_critical",
				Type:        schema.TypeBigInt,
				Description: "Count of CRITICAL vulnerabilities.",
				Resolver:    schema.PathResolver("Vulnerabilities.Critical"),
			},
			{
				Name:        "vulnerabilities_high",
				Type:        schema.TypeBigInt,
				Description: "Count of HIGH vulnerabilities.",
				Resolver:    schema.PathResolver("Vulnerabilities.High"),
			},
			{
				Name:        "vulnerabilities_medium",
				Type:        schema.TypeBigInt,
				Description: "Count of MEDIUM vulnerabilities.",
				Resolver:    schema.PathResolver("Vulnerabilities.Medium"),
			},
			{
				Name:        "vulnerabilities_low",
				Type:        schema.TypeBigInt,
				Description: "Count of LOW vulnerabilities.",
				Resolver:    schema.PathResolver("Vulnerabilities.Low"),
			},
			{
				Name:        "vulnerabilities_negligible",
				Type:        schema.TypeBigInt,
				Description: "Count of NEGLIGIBLE vulnerabilities.",
				Resolver:    schema.PathResolver("Vulnerabilities.Negligible"),
			},
			{
				Name:        "vulnerabilities_undefined",
				Type:        schema.TypeBigInt,
				Description: "Count of other vulnerabilities.",
				Resolver:    schema.PathResolver("Vulnerabilities.Undefined"),
			},
		},
	}

}

func fetchContainerRegistryScanResults(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	g := errgroup.Group{}
	ch := make(chan interface{})

	g.Go(func() error {
		defer close(ch)
		req := &containerregistry.ListImagesRequest{FolderId: c.MultiplexedResourceId}
		it := c.Services.ContainerRegistry.Image().ImageIterator(ctx, req)
		for it.Next() {
			ch <- it.Value()
		}
		return nil
	})

	g.Go(func() error {
		for image := range ch {
			req := &containerregistry.ListScanResultsRequest{
				Id: &containerregistry.ListScanResultsRequest_ImageId{
					ImageId: image.(*containerregistry.Image).Id,
				},
			}
			it := c.Services.ContainerRegistry.Scanner().ScannerIterator(ctx, req)
			for it.Next() {
				res <- it.Value()
			}
		}
		return nil
	})

	return g.Wait()
}
