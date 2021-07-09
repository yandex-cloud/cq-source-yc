package resources

import (
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:      "yandex",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"kms.keys":          KmsKeyring(),
			"compute.images":    ComputeImages(),
			"compute.instances": ComputeInstances(),
			"compute.disks":     ComputeDisks(),
			//"vpc.networks":         VpcNetworks(),
			//"vpc.subnets":          VpcSubnetworks(),
			"vpc.addresses": VpcAddresses(),
			//"iam.service_accounts": IamServiceAccounts(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
