package client

import (
	"context"

	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/go-sdk/gen/compute"
	"github.com/yandex-cloud/go-sdk/gen/iam"
	"github.com/yandex-cloud/go-sdk/gen/kms"
	"github.com/yandex-cloud/go-sdk/gen/resourcemanager"
	"github.com/yandex-cloud/go-sdk/gen/vpc"
)

type Services struct {
	KMS             *kms.KMS
	Compute         *compute.Compute
	VPC             *vpc.VPC
	IAM             *iam.IAM
	ResourceManager *resourcemanager.ResourceManager
}

func initServices(_ context.Context, sdk *ycsdk.SDK) (*Services, error) {
	return &Services{
		KMS:             sdk.KMS(),
		Compute:         sdk.Compute(),
		VPC:             sdk.VPC(),
		IAM:             sdk.IAM(),
		ResourceManager: sdk.ResourceManager(),
	}, nil
}
