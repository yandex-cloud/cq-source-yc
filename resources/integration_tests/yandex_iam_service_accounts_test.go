package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/yandex-cloud/cq-provider-yandex/resources"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
)

func TestIntegrationIAMServiceAccount(t *testing.T) {
	var tfTmpl = fmt.Sprintf(`
resource "yandex_iam_service_account" "foo" {
  name = "cq-sa-test-sa-%[1]s"
}
`, suffix)
	testIntegrationHelper(t, resources.IAMServiceAccounts(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "yandex_iam_service_accounts",
			Filter: func(sq squirrel.SelectBuilder, _ *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name": fmt.Sprintf("cq-sa-test-sa-%s", suffix),
				},
			}},
		}
	}, tfTmpl)
}
