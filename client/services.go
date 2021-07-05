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
	Kms             *kms.KMS
	Compute         *compute.Compute
	Vpc             *vpc.VPC
	Iam             *iam.IAM
	ResourceManager *resourcemanager.ResourceManager
}

func initServices(ctx context.Context, sdk *ycsdk.SDK) (*Services, error) {
	return &Services{
		Kms:             sdk.KMS(),
		Compute:         sdk.Compute(),
		Vpc:             sdk.VPC(),
		Iam:             sdk.IAM(),
		ResourceManager: sdk.ResourceManager(),
	}, nil
}
