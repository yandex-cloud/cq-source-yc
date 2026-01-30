package vpc

import (
	"context"
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
)

func SubnetUsedAddresses() *schema.Table {
	return &schema.Table{
		Name:        "yc_vpc_used_addresses",
		Description: `https://yandex.cloud/ru/docs/vpc/api-ref/grpc/Subnet/listUsedAddresses#yandex.cloud.vpc.v1.UsedAddress`,
		Resolver:    fetchSubnetUsedAddresses,
		Transform:   client.TransformWithStruct(&vpc.UsedAddress{}, transformers.WithPrimaryKeys("Address")),
		Columns: schema.ColumnList{
			schema.Column{
				Name:       "subnet_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("id"),
				PrimaryKey: true,
			},
			client.CloudIdColumn,
			client.FolderIdColumn,
		},
	}
}

func fetchSubnetUsedAddresses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	subnet, ok := parent.Item.(*vpc.Subnet)
	if !ok {
		return fmt.Errorf("parent is not type of *vpc.Subnet: %+v", subnet)
	}

	it := c.SDK.VPC().Subnet().SubnetUsedAddressesIterator(ctx, &vpc.ListUsedAddressesRequest{SubnetId: subnet.Id})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}
