package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/yandex-cloud/cq-provider-yandex/resources"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
)

func TestIntegrationKMSSymmetricKeys(t *testing.T) {
	var tfTmpl = `
resource "yandex_kms_symmetric_key" "cq-keys-test-keys-%[1]s" {
  name = "cq-keys-test-keys-%[1]s"
}
`
	yandexTestIntegrationHelper(t, resources.KMSSymmetricKeys(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "yandex_kms_symmetric_keys",
			Filter: func(sq squirrel.SelectBuilder, _ *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name": fmt.Sprintf("cq-keys-test-keys-%s", suffix),
				},
			}},
		}
	}, tfTmpl)
}
