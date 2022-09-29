package access_bindings

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	"google.golang.org/grpc"
)

type accessBindingsClient interface {
	ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error)
}

func fetchAccessBindings(provider func(*client.Client) accessBindingsClient) schema.TableResolver {
	return func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
		c := meta.(*client.Client)
		_client := provider(c)
		req := &access.ListAccessBindingsRequest{ResourceId: c.MultiplexedResourceId}
		for {
			resp, err := _client.ListAccessBindings(ctx, req)
			if err != nil {
				return err
			}
			for _, b := range resp.GetAccessBindings() {
				res <- b
			}

			if resp.GetNextPageToken() == "" {
				return nil
			}
			req.PageToken = resp.NextPageToken
		}
	}
}
