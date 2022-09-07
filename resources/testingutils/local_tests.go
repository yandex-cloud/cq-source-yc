package testingutils

import (
	"context"
	"testing"

	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/hashicorp/go-hclog"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/cq-provider-yandex/resources/testingutils/servers"
)

func LocalTestProvider(t *testing.T, resourceMap map[string]*schema.Table, verifiers map[string][]providertest.Verifier) {
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

	// TODO: s3 testing wit MINIO

	resource := providertest.ResourceTestCase{
		Provider: &provider.Provider{
			ResourceMap: resourceMap,
			Config: func() provider.Config {
				return &client.Config{
					FolderIDs: []string{"test-folder-id"},
				}
			},
			Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, diag.Diagnostics) {
				log := logging.New(&hclog.LoggerOptions{Level: hclog.Debug})
				folderIds := []string{"test-folder-id"}
				cloudIds := []string{"test-cloud-id"}
				organizationIds := []string{"test-organization-id"}
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
				c := client.NewYandexClient(log, folderIds, cloudIds, organizationIds, services, nil)
				return c, nil
			},
		},
		Verifiers: verifiers,
	}

	providertest.TestResource(t, resource)
}
