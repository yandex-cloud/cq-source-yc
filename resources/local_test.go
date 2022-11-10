package resources_test

import (
	"testing"

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
	"github.com/yandex-cloud/cq-provider-yandex/resources/services/vpc"
	"github.com/yandex-cloud/cq-provider-yandex/resources/testingutils"
)

func TestComputeDisks(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"ComputeDisks": compute.Disks(),
	}
	testingutils.LocalTestProvider(t, resourceMap)
}

func TestComputeImages(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"ComputeImages": compute.Images(),
	}
	testingutils.LocalTestProvider(t, resourceMap)
}

func TestComputeInstances(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"ComputeInstances": compute.Instances(),
	}
	testingutils.LocalTestProvider(t, resourceMap)
}

func TestK8S(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"K8SCluster": k8s.Clusters(),
	}
	testingutils.LocalTestProvider(t, resourceMap)
}

func TestVPCAdresses(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"VPCAddresses": vpc.Addresses(),
	}
	testingutils.LocalTestProvider(t, resourceMap)
}

func TestVPCNetworks(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"VPCNetworks": vpc.Networks(),
	}
	testingutils.LocalTestProvider(t, resourceMap)
}

func TestVPCSecurityGroups(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"VPCSecurityGroups": vpc.SecurityGroups(),
	}
	testingutils.LocalTestProvider(t, resourceMap)
}

func TestVPCSubnets(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"VPCSubnets": vpc.Subnets(),
	}
	testingutils.LocalTestProvider(t, resourceMap)
}

func TestKMS(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"KMSSymmetricKeys": kms.SymmetricKeys(),
	}
	testingutils.LocalTestProvider(t, resourceMap)
}

func TestOrganizationManagerOrganizations(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"OrganizationManagerOrganizations": organizationmanager.Organizations(),
	}
	testingutils.LocalTestProvider(t, resourceMap)
}

func TestOrganizationManagerFederations(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"OrganizationManagerFederations": organizationmanager.Federations(),
	}
	testingutils.LocalTestProvider(t, resourceMap)
}

func TestCertificateManager(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"CertificateManagerCertificates": certificatemanager.Certificates(),
	}
	testingutils.LocalTestProvider(t, resourceMap)
}

func TestIAM(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"IAMServiceAccounts": iam.ServiceAccounts(),
	}
	testingutils.LocalTestProvider(t, resourceMap)
}

func TestContainerRegistry(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"ContainerRegistryScanResults": containerregistry.ScanResults(),
	}
	testingutils.LocalTestProvider(t, resourceMap)
}

func TestServerlessApiGateway(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"ServerlessApiGateways": apigateways.ApiGateways(),
	}
	testingutils.LocalTestProvider(t, resourceMap)
}

func TestAccessBindingsByCloud(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"AccessBindingsByCloud": access_bindings.ByCloud(),
	}
	testingutils.LocalTestProvider(t, resourceMap)
}

func TestAccessBindingsByFolder(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"AccessBindingsByFolder": access_bindings.ByFolder(),
	}
	testingutils.LocalTestProvider(t, resourceMap)
}

func TestAccessBindingsByOrganization(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"AccessBindingsByOrganization": access_bindings.ByOrganization(),
	}
	testingutils.LocalTestProvider(t, resourceMap)
}

func TestIAMUserAccountsByCloud(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"IAMUserAccountsByCloud": iam.UserAccountsByCloud(),
	}
	testingutils.LocalTestProvider(t, resourceMap)
}

func TestIAMUserAccountsByFolder(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"IAMUserAccountsByFolder": iam.UserAccountsByFolder(),
	}
	testingutils.LocalTestProvider(t, resourceMap)
}

func TestIAMUserAccountsByOrganization(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"IAMUserAccountsByOrganization": iam.UserAccountsByOrganization(),
	}
	testingutils.LocalTestProvider(t, resourceMap)
}
