package plugin

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/resources/services/access_bindings"
	apigateways "github.com/yandex-cloud/cq-provider-yandex/resources/services/api_gateways"
	"github.com/yandex-cloud/cq-provider-yandex/resources/services/certificatemanager"
	"github.com/yandex-cloud/cq-provider-yandex/resources/services/compute"
	"github.com/yandex-cloud/cq-provider-yandex/resources/services/containerregistry"
	"github.com/yandex-cloud/cq-provider-yandex/resources/services/iam"
	"github.com/yandex-cloud/cq-provider-yandex/resources/services/k8s"
	"github.com/yandex-cloud/cq-provider-yandex/resources/services/kms"
	"github.com/yandex-cloud/cq-provider-yandex/resources/services/organizationmanager"
	"github.com/yandex-cloud/cq-provider-yandex/resources/services/resourcemanager"
	"github.com/yandex-cloud/cq-provider-yandex/resources/services/storage"
	"github.com/yandex-cloud/cq-provider-yandex/resources/services/vpc"
)

func Tables() []*schema.Table {
	return []*schema.Table{
		access_bindings.ByCloud(),
		access_bindings.ByFolder(),
		access_bindings.ByOrganization(),
		certificatemanager.Certificates(),
		compute.Disks(),
		compute.Images(),
		compute.Instances(),
		containerregistry.Images(),
		containerregistry.Registries(),
		containerregistry.ScanResults(),
		iam.ServiceAccounts(),
		iam.UserAccountsByCloud(),
		iam.UserAccountsByFolder(),
		iam.UserAccountsByOrganization(),
		k8s.Clusters(),
		kms.SymmetricKeys(),
		organizationmanager.Federations(),
		organizationmanager.Organizations(),
		resourcemanager.Clouds(),
		resourcemanager.Folders(),
		apigateways.ApiGateways(),
		storage.Buckets(),
		vpc.Addresses(),
		vpc.Networks(),
		vpc.SecurityGroups(),
		vpc.Subnets(),
	}
}
