package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/yandex-cloud/cq-provider-yandex/resources"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
)

func TestIntegrationComputeDisks(t *testing.T) {
	var tfTmpl = `
resource "yandex_compute_disk" "cq-disk-test-disk-%[1]s" {
  name = "cq-disk-test-disk-%[1]s"
}
`
	suffix := acctest.RandString(10)
	yandexTestIntegrationHelper(t, resources.ComputeDisks(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "yandex_compute_disks",
			Filter: func(sq squirrel.SelectBuilder, _ *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name": fmt.Sprintf("cq-disk-test-disk-%s", suffix),
				},
			}},
		}
	}, tfTmpl, suffix)
}
