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
			"kms.keys":                        KMSSymmetricKeys(),
			"compute.addresses":               VPCAddresses(),
			"compute.images":                  ComputeImages(),
			"compute.instances":               ComputeInstances(),
			"compute.disks":                   ComputeDisks(),
			"vpc.networks":                    VPCNetworks(),
			"vpc.subnets":                     VPCSubnets(),
			"vpc.addresses":                   VPCAddresses(),
			"vpc.security_groups":             VPCSecurityGroups(),
			"iam.service_accounts":            IAMServiceAccounts(),
			"access_bindings_by_folder":       AccessBindingsByFolder(),
			"access_bindings_by_cloud":        AccessBindingsByCloud(),
			"resourcemanager.clouds":          ResourceManagerClouds(),
			"resourcemanager.folders":         ResourceManagerFolders(),
			"objects_storage":                 ObjectsStorages(),
			"k8s.clusters":                    K8SClusters(),
			"certificatemanager.certificates": CertificateManagerCertificates(),
			"serverless.apigateways":          ServerlessApiGateways(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
