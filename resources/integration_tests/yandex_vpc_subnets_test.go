package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/yandex-cloud/cq-provider-yandex/resources"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
)

func TestIntegrationVPCSubnets(t *testing.T) {
	var tfTmpl = `
resource "yandex_vpc_network" "cq-subnet-test-net-%[1]s" {
  name = "cq-subnet-test-net-%[1]s"
}

resource "yandex_vpc_subnet" "cq-subnet-test-subnet-%[1]s" {
  network_id     = yandex_vpc_network.cq-subnet-test-net-%[1]s.id
  v4_cidr_blocks = ["10.2.0.0/16"]
  name           = "cq-subnet-test-subnet-%[1]s"
}
`
	suffix := acctest.RandString(10)
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
						"name": fmt.Sprintf("cq-subnet-test-subnet-%s", suffix),
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
	}, tfTmpl, suffix)
}
