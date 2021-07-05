package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/kms/v1"
)

func KmsKeyring() *schema.Table {
	return &schema.Table{
		Name:         "yandex_kms_symmetric_keys",
		Resolver:     fetchKmsSymmetricKeys,
		Multiplex:    client.FolderMultiplex,
		DeleteFilter: client.DeleteFolderFilter,
		IgnoreError:  client.IgnoreErrorHandler,
		//PostResourceResolver: client.AddGcpMetadata,
		Columns: []schema.Column{
			{
				Name:     "folder_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveFolderID,
			},
			{
				Name: "create_time",
				Type: schema.TypeString,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
		},
	}
}

func fetchKmsSymmetricKeys(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	// TODO: iterate over all  folders ???
	locations := []string{c.FolderId}

	for _, f := range locations {
		req := &kms.ListSymmetricKeysRequest{FolderId: f}
		it := c.Services.Kms.SymmetricKey().SymmetricKeyIterator(ctx, req)
		for it.Next() {
			res <- it.Value().GetId()
		}
	}
	return nil
}
