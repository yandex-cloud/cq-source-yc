package vpc

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
)

func Addresses() *schema.Table {
	return &schema.Table{
		Name:        "yc_vpc_addresses",
		Description: `https://cloud.yandex.ru/docs/vpc/api-ref/grpc/address_service#Address2`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchAddresses,
		Transform:   client.TransformWithStruct(&vpc.Address{}, client.PrimaryKeyIdTransformer),
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchAddresses(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	it := c.SDK.VPC().Address().AddressIterator(ctx, &vpc.ListAddressesRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
