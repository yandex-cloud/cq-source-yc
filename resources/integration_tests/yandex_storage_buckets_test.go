package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/yandex-cloud/cq-provider-yandex/resources"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
)

func TestStorageBuckets(t *testing.T) {
	var tfTmpl = fmt.Sprintf(`
resource "yandex_storage_bucket" "foo" {
 bucket = "cq-s3-test-s3-%[1]s"
}
`, suffix)
	testIntegrationHelper(t, resources.StorageBuckets(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "yandex_storage_buckets",
			Filter: func(sq squirrel.SelectBuilder, _ *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"id": fmt.Sprintf("cq-s3-test-s3-%s", suffix),
					},
				},
			},
		}
	}, tfTmpl)
}
