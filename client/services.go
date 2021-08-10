package client

import (
	"context"

	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/go-sdk/gen/apigateway"
	"github.com/yandex-cloud/go-sdk/gen/certificatemanager"
	"github.com/yandex-cloud/go-sdk/gen/compute"
	"github.com/yandex-cloud/go-sdk/gen/iam"
	"github.com/yandex-cloud/go-sdk/gen/kms"
	k8s "github.com/yandex-cloud/go-sdk/gen/kubernetes"
	"github.com/yandex-cloud/go-sdk/gen/resourcemanager"
	"github.com/yandex-cloud/go-sdk/gen/vpc"
)

type Services struct {
	KMS                *kms.KMS
	Compute            *compute.Compute
	VPC                *vpc.VPC
	IAM                *iam.IAM
	ResourceManager    *resourcemanager.ResourceManager
	K8S                *k8s.Kubernetes
	CertificateManager *certificatemanager.CertificateManager
	ApiGateway         *apigateway.Apigateway
}

func initServices(_ context.Context, sdk *ycsdk.SDK) (*Services, error) {
	return &Services{
		KMS:                sdk.KMS(),
		Compute:            sdk.Compute(),
		VPC:                sdk.VPC(),
		IAM:                sdk.IAM(),
		ResourceManager:    sdk.ResourceManager(),
		K8S:                sdk.Kubernetes(),
		CertificateManager: sdk.Certificates(),
		ApiGateway:         sdk.Serverless().APIGateway(),
	}, nil
}
