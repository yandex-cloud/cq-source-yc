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

func TestCompute(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"ComputeDisks":     compute.Disks(),
		"ComputeImages":    compute.Images(),
		"ComputeInstances": compute.Instances(),
	}
	testingutils.LocalTestProvider(t, resourceMap)
}

func TestK8S(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"K8SCluster":    k8s.Clusters(),
		"K8SNodeGroups": k8s.NodeGroups(),
	}
	testingutils.LocalTestProvider(t, resourceMap)
}

func TestVPC(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"VPCAddresses":      vpc.Addresses(),
		"VPCNetworks":       vpc.Networks(),
		"VPCSecurityGroups": vpc.SecurityGroups(),
		"VPCSubnets":        vpc.Subnets(),
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
