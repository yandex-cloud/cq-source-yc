package client

import (
	"context"

	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/go-sdk/gen/apigateway"
	"github.com/yandex-cloud/go-sdk/gen/certificatemanager"
	"github.com/yandex-cloud/go-sdk/gen/compute"
	"github.com/yandex-cloud/go-sdk/gen/containerregistry"
	"github.com/yandex-cloud/go-sdk/gen/iam"
	"github.com/yandex-cloud/go-sdk/gen/kms"
	k8s "github.com/yandex-cloud/go-sdk/gen/kubernetes"
	"github.com/yandex-cloud/go-sdk/gen/organizationmanager"
	"github.com/yandex-cloud/go-sdk/gen/organizationmanager/saml"
	"github.com/yandex-cloud/go-sdk/gen/resourcemanager"
	"github.com/yandex-cloud/go-sdk/gen/vpc"
)

type Services struct {
	ApiGateway              *apigateway.Apigateway
	CertificateManager      *certificatemanager.CertificateManager
	Compute                 *compute.Compute
	ContainerRegistry       *containerregistry.ContainerRegistry
	IAM                     *iam.IAM
	K8S                     *k8s.Kubernetes
	KMS                     *kms.KMS
	OrganizationManager     *organizationmanager.OrganizationManager
	OrganizationManagerSAML *saml.OrganizationManagerSAML
	ResourceManager         *resourcemanager.ResourceManager
	VPC                     *vpc.VPC
}

func initServices(_ context.Context, sdk *ycsdk.SDK) (*Services, error) {
	return &Services{
		ApiGateway:              sdk.Serverless().APIGateway(),
		CertificateManager:      sdk.Certificates(),
		Compute:                 sdk.Compute(),
		ContainerRegistry:       sdk.ContainerRegistry(),
		IAM:                     sdk.IAM(),
		K8S:                     sdk.Kubernetes(),
		KMS:                     sdk.KMS(),
		OrganizationManager:     sdk.OrganizationManager(),
		OrganizationManagerSAML: sdk.OrganizationManagerSAML(),
		ResourceManager:         sdk.ResourceManager(),
		VPC:                     sdk.VPC(),
	}, nil
}
