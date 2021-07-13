package resources

import (
	"context"

	"github.com/yandex-cloud/cq-provider-yandex/tools"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
)

func ComputeDisks() *schema.Table {
	table, err := tools.GenerateTable(
		tools.WithTableName("yandex_computes_disks"),
		tools.WithProtoFile("Disk", "yandex/cloud/compute/v1/disk.proto", "cloudapi"),
		tools.WithResolver(fetchComputeDisks),
	)
	if err != nil {
		return &schema.Table{}
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
