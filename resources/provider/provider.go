package provider

import (
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/cq-provider-yandex/resources"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:      "yandex",
		Version:   "dev",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"AccessBindingsByCloud":            resources.AccessBindingsByCloud(),
			"AccessBindingsByFolder":           resources.AccessBindingsByFolder(),
			"AccessBindingsByOrganization":     resources.AccessBindingsByOrganization(),
			"CertificateManagerCertificates":   resources.CertificateManagerCertificates(),
			"ComputeDisks":                     resources.ComputeDisks(),
			"ComputeImages":                    resources.ComputeImages(),
			"ComputeInstances":                 resources.ComputeInstances(),
			"ContainerRegistryScanResults":     resources.ContainerRegistryScanResults(),
			"IAMServiceAccounts":               resources.IAMServiceAccounts(),
			"IAMUserAccountsByCloud":           resources.IAMUserAccountsByCloud(),
			"IAMUserAccountsByFolder":          resources.IAMUserAccountsByFolder(),
			"IAMUserAccountsByOrganization":    resources.IAMUserAccountsByOrganization(),
			"K8SClusters":                      resources.K8SClusters(),
			"KMSSymmetricKeys":                 resources.KMSSymmetricKeys(),
			"OrganizationManagerFederations":   resources.OrganizationManagerFederations(),
			"OrganizationManagerOrganizations": resources.OrganizationManagerOrganizations(),
			"ResourceManagerClouds":            resources.ResourceManagerClouds(),
			"ResourceManagerFolders":           resources.ResourceManagerFolders(),
			"ServerlessApiGateways":            resources.ServerlessApiGateways(),
			"StorageBuckets":                   resources.StorageBuckets(),
			"VPCAddresses":                     resources.VPCAddresses(),
			"VPCNetworks":                      resources.VPCNetworks(),
			"VPCSecurityGroups":                resources.VPCSecurityGroups(),
			"VPCSubnets":                       resources.VPCSubnets(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
