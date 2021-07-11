package resources

import (
	"context"
	"github.com/yandex-cloud/cq-provider-yandex/tools"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
)

func ComputeInstances() *schema.Table {
	gen, err := tools.NewTableGenerator(
		"Compute",
		"Instance",
		tools.WithProtoFile("resources/proto/instance.proto"),
		tools.WithFetcher(fetchComputeInstances),
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

func fetchComputeInstances(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	locations := []string{c.FolderId}

	for _, f := range locations {
		req := &compute.ListInstancesRequest{FolderId: f}
		it := c.Services.Compute.Instance().InstanceIterator(ctx, req)
		for it.Next() {
			res <- it.Value()
		}
	}
	return nil
}
