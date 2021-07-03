package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
)

func VpcAddresses() *schema.Table {
	return &schema.Table{
		Name:         "yandex_vpc_public_addresses",
		Resolver:     fetchVpcAddresses,
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
				Name:     "address_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveResourceId,
			},
			{
				Name:     "address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Address.ExternalIpv4Address.Address"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveLabels,
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveAsTime,
			},
			{
				Name: "reserved",
				Type: schema.TypeBool,
			},
			{
				Name: "used",
				Type: schema.TypeBool,
			},
		},
	}
}

func fetchVpcAddresses(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	// TODO: iterate over all  folders ???
	locations := []string{c.FolderId}

	for _, f := range locations {
		req := &vpc.ListAddressesRequest{FolderId: f}
		it := c.Services.Vpc.Address().AddressIterator(ctx, req)
		for it.Next() {
			res <- it.Value()
		}
	}
	return nil
}
