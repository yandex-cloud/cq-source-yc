package resources

import (
	"context"
	"github.com/yandex-cloud/cq-provider-yandex/tools"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
)

func ComputeDisks() *schema.Table {
	gen, err := tools.NewTableGenerator(
		"yandex_compute_disks",
		"Compute",
		"Disk",
		"resources/proto/disk.proto",
		tools.GetCommonDefaultColumns("disk"),
		tools.IgnoredColumns{},
		fetchComputeDisks,
	)
	if err != nil {
		return nil
	}
	table, err := gen.Generate()
	if err != nil {
		return nil
	}
	return table
}

func fetchComputeDisks(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	locations := []string{c.FolderId}

	for _, f := range locations {
		req := &compute.ListDisksRequest{FolderId: f}
		it := c.Services.Compute.Disk().DiskIterator(ctx, req)
		for it.Next() {
			res <- it.Value()
		}
	}
	return nil
}
