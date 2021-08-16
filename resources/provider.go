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
			"AccessBindingsByCloud":            AccessBindingsByCloud(),
			"AccessBindingsByFolder":           AccessBindingsByFolder(),
			"AccessBindingsByOrganization":     AccessBindingsByOrganization(),
			"CertificateManagerCertificates":   CertificateManagerCertificates(),
			"ComputeDisks":                     ComputeDisks(),
			"ComputeImages":                    ComputeImages(),
			"ComputeInstances":                 ComputeInstances(),
			"IAMServiceAccounts":               IAMServiceAccounts(),
			"IAMUserAccountsByClouds":          IAMUserAccountsByClouds(),
			"IAMUserAccountsByFolders":         IAMUserAccountsByFolders(),
			"IAMUserAccountsByOrganizations":   IAMUserAccountsByOrganizations(),
			"K8SClusters":                      K8SClusters(),
			"KMSSymmetricKeys":                 KMSSymmetricKeys(),
			"StorageBuckets":                   StorageBuckets(),
			"OrganizationManagerFederations":   OrganizationManagerFederations(),
			"OrganizationManagerOrganizations": OrganizationManagerOrganizations(),
			"ResourceManagerClouds":            ResourceManagerClouds(),
			"ResourceManagerFolders":           ResourceManagerFolders(),
			"ServerlessApiGateways":            ServerlessApiGateways(),
			"VPCAddresses":                     VPCAddresses(),
			"VPCNetworks":                      VPCNetworks(),
			"VPCSecurityGroups":                VPCSecurityGroups(),
			"VPCSubnets":                       VPCSubnets(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
