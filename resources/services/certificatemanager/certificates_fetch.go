package certificatemanager

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/certificatemanager/v1"
)

func fetchCertificates(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	req := &certificatemanager.ListCertificatesRequest{FolderId: c.MultiplexedResourceId}
	it := c.Services.CertificateManager.Certificate().CertificateIterator(ctx, req)
	for it.Next() {
		res <- it.Value()
	}

	return nil
}
