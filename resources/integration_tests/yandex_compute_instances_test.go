package integration_tests

import (
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/yandex-cloud/cq-provider-yandex/resources"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
)

func TestIntegrationComputeInstances(t *testing.T) {
	yandexTestIntegrationHelper(t, resources.ComputeInstances(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "yandex_compute_instances",
			Filter: func(sq squirrel.SelectBuilder, _ *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name": "cq-instance-test-instance",
				},
			}},
		}
	})
}
