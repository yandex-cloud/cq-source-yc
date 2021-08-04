package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/yandex-cloud/cq-provider-yandex/resources"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
)

func TestIntegrationComputeImages(t *testing.T) {
	var tfTmpl = `
resource "yandex_compute_image" "cq-image-test-image-%[1]s" {
  name         = "cq-image-test-image-%[1]s"
  source_image = "fd8vmcue7aajpmeo39kk"
}
`
	yandexTestIntegrationHelper(t, resources.ComputeImages(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "yandex_compute_images",
			Filter: func(sq squirrel.SelectBuilder, _ *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name": fmt.Sprintf("cq-image-test-image-%s", suffix),
				},
			}},
		}
	}, tfTmpl)
}
