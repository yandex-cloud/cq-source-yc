package resources

import (
	"github.com/GennadySpb/cq-provider-yandex/client"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:      "yandex",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"kms.keys": KmsKeyring(),
			//"compute.addresses":    ComputeAddresses(),
			"compute.images": ComputeImages(),
			//"compute.instances":    ComputeInstances(),
			//"compute.networks":     ComputeNetworks(),
			//"compute.disks":        ComputeDisks(),
			//"compute.subnets":      ComputeSubnetworks(),
			//"iam.service_accounts": IamServiceAccounts(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
