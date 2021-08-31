package resources_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/cq-provider-yandex/resources"
)

func TestStorageBuckets(t *testing.T) {
	s := session.Must(session.NewSession())
	resource := providertest.ResourceTestData{
		Table:  resources.StorageBuckets(),
		Config: client.Config{},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			c := client.NewYandexClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), nil, nil, nil, nil, s3.New(s, aws.NewConfig().
				WithEndpoint("http://cq_provider_yandex_s3:9000").
				WithRegion("us-east-1").
				WithCredentials(credentials.NewStaticCredentials("user", "12345678", ""))))
			return c, nil
		},
		Verifiers: []providertest.Verifier{
			providertest.VerifyAtLeastOneRow("yandex_storage_buckets"),
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
}
