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
	lockbox "github.com/yandex-cloud/go-sdk/gen/lockboxsecret"
	"github.com/yandex-cloud/go-sdk/gen/organizationmanager"
	"github.com/yandex-cloud/go-sdk/gen/organizationmanager/saml"
	"github.com/yandex-cloud/go-sdk/gen/resourcemanager"
	"github.com/yandex-cloud/go-sdk/gen/storage-api"
	"github.com/yandex-cloud/go-sdk/gen/vpc"
)

type (
	Services struct {
		CertificateManager      *certificatemanager.CertificateManager
		Compute                 *compute.Compute
		ContainerRegistry       *containerregistry.ContainerRegistry
		IAM                     *iam.IAM
		K8S                     *k8s.Kubernetes
		KMS                     *kms.KMS
		OrganizationManager     *organizationmanager.OrganizationManager
		OrganizationManagerSAML *saml.OrganizationManagerSAML
		ResourceManager         *resourcemanager.ResourceManager
		ApiGateway              *apigateway.Apigateway
		Storage                 *storage.StorageAPI
		VPC                     *vpc.VPC
		LockboxSecret           *lockbox.LockboxSecret
	}
)

func initServices(_ context.Context, sdk *ycsdk.SDK) (*Services, error) {
	return &Services{
		CertificateManager:      sdk.Certificates(),
		Compute:                 sdk.Compute(),
		ContainerRegistry:       sdk.ContainerRegistry(),
		IAM:                     sdk.IAM(),
		K8S:                     sdk.Kubernetes(),
		KMS:                     sdk.KMS(),
		OrganizationManager:     sdk.OrganizationManager(),
		OrganizationManagerSAML: sdk.OrganizationManagerSAML(),
		ResourceManager:         sdk.ResourceManager(),
		ApiGateway:              sdk.Serverless().APIGateway(),
		Storage:                 sdk.StorageAPI(),
		VPC:                     sdk.VPC(),
		LockboxSecret:           sdk.LockboxSecret(),
	}, nil
}
