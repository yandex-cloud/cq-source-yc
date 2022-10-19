package testingutils

import (
	"context"
	"testing"

	// "github.com/cloudquery/plugin-sdk/schema"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/cq-provider-yandex/resources/testingutils/servers"
)

func LocalTestProvider(t *testing.T, resourceMap map[string]*schema.Table) {
	back := context.Background()
	ctx, cancel := context.WithCancel(back)
	defer cancel()

	computeServ, err := servers.StartComputeServer(t, ctx)
	if err != nil {
		t.Fatal(err)
	}

	k8sServ, err := servers.StartK8SServer(t, ctx)
	if err != nil {
		t.Fatal(err)
	}

	vpcServ, err := servers.StartVpcServer(t, ctx)
	if err != nil {
		t.Fatal(err)
	}

	kmsServ, err := servers.StartKmsServer(t, ctx)
	if err != nil {
		t.Fatal(err)
	}

	organizationManagerServ, err := servers.StartOrganizationManagerServer(t, ctx)
	if err != nil {
		t.Fatal(err)
	}

	organizationManagerSAMLServ, err := servers.StartOrganizationManagerSAMLServer(t, ctx)
	if err != nil {
		t.Fatal(err)
	}

	certificateManagerServ, err := servers.StartCertificateManagerServer(t, ctx)
	if err != nil {
		t.Fatal(err)
	}

	iamServ, err := servers.StartIamServer(t, ctx)
	if err != nil {
		t.Fatal(err)
	}

	containerRegistryServ, err := servers.StartContainerRegistryServer(t, ctx)
	if err != nil {
		t.Fatal(err)
	}

	apiGatewayServ, err := servers.StartApiGatewayServer(t, ctx)
	if err != nil {
		t.Fatal(err)
	}

	resourceManagerServ, err := servers.StartResourceManagerServer(t, ctx)
	if err != nil {
		t.Fatal(err)
	}

	services := &client.Services{
		Compute:                 computeServ,
		K8S:                     k8sServ,
		VPC:                     vpcServ,
		KMS:                     kmsServ,
		OrganizationManager:     organizationManagerServ,
		OrganizationManagerSAML: organizationManagerSAMLServ,
		CertificateManager:      certificateManagerServ,
		IAM:                     iamServ,
		ContainerRegistry:       containerRegistryServ,
		ApiGateway:              apiGatewayServ,
		ResourceManager:         resourceManagerServ,
	}

	for _, table := range resourceMap {
		client.MockTestHelper(t, table, func() (*client.Services, error) { return services, nil }, client.TestOptions{})
	}
}
