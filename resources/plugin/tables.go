package plugin

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/resources/services/access_bindings"
	"github.com/yandex-cloud/cq-provider-yandex/resources/services/certificatemanager"
	"github.com/yandex-cloud/cq-provider-yandex/resources/services/compute"
	"github.com/yandex-cloud/cq-provider-yandex/resources/services/containerregistry"
	"github.com/yandex-cloud/cq-provider-yandex/resources/services/iam"
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
	}
}
