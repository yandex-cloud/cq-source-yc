package resources_test

import (
	"testing"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/yandex-cloud/cq-provider-yandex/resources"
	"github.com/yandex-cloud/cq-provider-yandex/resources/testingutils"
)

func TestCompute(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"ComputeDisks":     resources.ComputeDisks(),
		"ComputeImages":    resources.ComputeImages(),
		"ComputeInstances": resources.ComputeInstances(),
	}
	verifiers := map[string][]providertest.Verifier{
		"ComputeDisks": {
			providertest.VerifyNoEmptyColumnsExcept("yandex_compute_disks", "source_source_image_id", "source_source_snapshot_id"),
			providertest.VerifyAtMostOneOf("yandex_compute_disks", "source_source_image_id", "source_source_snapshot_id"),
			providertest.VerifyAtLeastOneRow(),
		},
		"ComputeImages":    {providertest.VerifyAtLeastOneRow()},
		"ComputeInstances": {providertest.VerifyAtLeastOneRow()},
	}
	testingutils.LocalTestProvider(t, resourceMap, verifiers)
}

func TestK8S(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"K8SCluster": resources.K8SClusters(),
	}
	verifiers := map[string][]providertest.Verifier{
		"K8SCluster": {providertest.VerifyAtLeastOneRow()},
	}
	testingutils.LocalTestProvider(t, resourceMap, verifiers)
}

func TestVPC(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"VPCAddresses":      resources.VPCAddresses(),
		"VPCNetworks":       resources.VPCNetworks(),
		"VPCSecurityGroups": resources.VPCSecurityGroups(),
		"VPCSubnets":        resources.VPCSubnets(),
	}
	verifiers := map[string][]providertest.Verifier{
		"VPCAddresses":      {providertest.VerifyAtLeastOneRow()},
		"VPCNetworks":       {providertest.VerifyAtLeastOneRow()},
		"VPCSecurityGroups": {providertest.VerifyAtLeastOneRow()},
		"VPCSubnets":        {providertest.VerifyAtLeastOneRow()},
	}
	testingutils.LocalTestProvider(t, resourceMap, verifiers)
}

func TestKMS(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"KMSSymmetricKeys": resources.KMSSymmetricKeys(),
	}
	verifiers := map[string][]providertest.Verifier{
		"KMSSymmetricKeys": {providertest.VerifyAtLeastOneRow()},
	}
	testingutils.LocalTestProvider(t, resourceMap, verifiers)
}

func TestOrganizationManagerOrganizations(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"OrganizationManagerOrganizations": resources.OrganizationManagerOrganizations(),
	}
	verifiers := map[string][]providertest.Verifier{
		"OrganizationManagerOrganizations": {providertest.VerifyAtLeastOneRow()},
	}
	testingutils.LocalTestProvider(t, resourceMap, verifiers)
}

func TestOrganizationManagerFederations(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"OrganizationManagerFederations": resources.OrganizationManagerFederations(),
	}
	verifiers := map[string][]providertest.Verifier{
		"OrganizationManagerFederations": {providertest.VerifyAtLeastOneRow()},
	}
	testingutils.LocalTestProvider(t, resourceMap, verifiers)
}

func TestCertificateManager(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"CertificateManagerCertificates": resources.CertificateManagerCertificates(),
	}
	verifiers := map[string][]providertest.Verifier{
		"CertificateManagerCertificates": {providertest.VerifyAtLeastOneRow()},
	}
	testingutils.LocalTestProvider(t, resourceMap, verifiers)
}

func TestIAM(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"IAMServiceAccounts": resources.IAMServiceAccounts(),
	}
	verifiers := map[string][]providertest.Verifier{
		"IAMServiceAccounts": {providertest.VerifyAtLeastOneRow()},
	}
	testingutils.LocalTestProvider(t, resourceMap, verifiers)
}

func TestContainerRegistry(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"ContainerRegistryScanResults": resources.ContainerRegistryScanResults(),
	}
	verifiers := map[string][]providertest.Verifier{
		"ContainerRegistryScanResults": {providertest.VerifyAtLeastOneRow()},
	}
	testingutils.LocalTestProvider(t, resourceMap, verifiers)
}

func TestServerlessApiGateway(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"ServerlessApiGateways": resources.ServerlessApiGateways(),
	}
	verifiers := map[string][]providertest.Verifier{
		"ServerlessApiGateways": {providertest.VerifyAtLeastOneRow()},
	}
	testingutils.LocalTestProvider(t, resourceMap, verifiers)
}

func TestAccessBindingsByCloud(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"AccessBindingsByCloud": resources.AccessBindingsByCloud(),
	}
	verifiers := map[string][]providertest.Verifier{
		"AccessBindingsByCloud": {providertest.VerifyAtLeastOneRow()},
	}
	testingutils.LocalTestProvider(t, resourceMap, verifiers)
}

func TestAccessBindingsByFolder(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"AccessBindingsByFolder": resources.AccessBindingsByFolder(),
	}
	verifiers := map[string][]providertest.Verifier{
		"AccessBindingsByFolder": {providertest.VerifyAtLeastOneRow()},
	}
	testingutils.LocalTestProvider(t, resourceMap, verifiers)
}

func TestAccessBindingsByOrganization(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"AccessBindingsByOrganization": resources.AccessBindingsByOrganization(),
	}
	verifiers := map[string][]providertest.Verifier{
		"AccessBindingsByOrganization": {providertest.VerifyAtLeastOneRow()},
	}
	testingutils.LocalTestProvider(t, resourceMap, verifiers)
}

func TestIAMUserAccountsByCloud(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"IAMUserAccountsByCloud": resources.IAMUserAccountsByCloud(),
	}
	verifiers := map[string][]providertest.Verifier{
		"IAMUserAccountsByCloud": {providertest.VerifyAtLeastOneRow()},
	}
	testingutils.LocalTestProvider(t, resourceMap, verifiers)
}

func TestIAMUserAccountsByFolder(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"IAMUserAccountsByFolder": resources.IAMUserAccountsByFolder(),
	}
	verifiers := map[string][]providertest.Verifier{
		"IAMUserAccountsByFolder": {providertest.VerifyAtLeastOneRow()},
	}
	testingutils.LocalTestProvider(t, resourceMap, verifiers)
}

func TestIAMUserAccountsByOrganization(t *testing.T) {
	resourceMap := map[string]*schema.Table{
		"IAMUserAccountsByOrganization": resources.IAMUserAccountsByOrganization(),
	}
	verifiers := map[string][]providertest.Verifier{
		"IAMUserAccountsByOrganization": {providertest.VerifyAtLeastOneRow()},
	}
	testingutils.LocalTestProvider(t, resourceMap, verifiers)
}
