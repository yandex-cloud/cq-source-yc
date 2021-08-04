package integration_tests

import (
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/yandex-cloud/cq-provider-yandex/resources"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
)

func TestIntegrationVPCSubnets(t *testing.T) {
	yandexTestIntegrationHelper(t, resources.VPCSubnets(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "yandex_vpc_subnets",
			Filter: func(sq squirrel.SelectBuilder, _ *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": "cq-subnet-test-subnet",
					},
				},
			},
			//Relations: []*providertest.ResourceIntegrationVerification{
			//	{
			//		Name:           "yandex_vpc_networks",
			//		ForeignKeyName: "network_id",
			//		ExpectedValues: []providertest.ExpectedValue{
			//			{
			//				Count: 1,
			//			},
			//		},
			//	},
			//},
		}
	})
}
