package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/yandex-cloud/cq-provider-yandex/resources"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
)

func TestIntegrationComputeInstances(t *testing.T) {
	var tfTmpl = `
resource "yandex_vpc_network" "cq-instance-test-net-%[1]s" {
  name = "cq-instance-test-net-%[1]s"
}

resource "yandex_vpc_subnet" "cq-instance-test-subnet-%[1]s" {
  network_id     = yandex_vpc_network.cq-instance-test-net-%[1]s.id
  v4_cidr_blocks = ["10.2.0.0/16"]
  name           = "cq-instance-test-subnet-%[1]s"
}

resource "yandex_compute_instance" "cq-instance-test-instance-%[1]s" {
  name = "cq-instance-test-instance-%[1]s"
  boot_disk {
    initialize_params {
      image_id = "fd8vmcue7aajpmeo39kk"
    }
  }
  network_interface {
    subnet_id = yandex_vpc_subnet.cq-instance-test-subnet-%[1]s.id
  }
  resources {
    cores = 2
    memory = 4
  }
}
`
	suffix := acctest.RandString(10)
	yandexTestIntegrationHelper(t, resources.ComputeInstances(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "yandex_compute_instances",
			Filter: func(sq squirrel.SelectBuilder, _ *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name": fmt.Sprintf("cq-instance-test-instance-%s", suffix),
				},
			}},
		}
	}, tfTmpl, suffix)
}
