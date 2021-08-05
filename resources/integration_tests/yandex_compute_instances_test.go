package integration_tests

import (
	"fmt"
	"testing"

	"github.com/yandex-cloud/cq-provider-yandex/resources"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
)

func TestIntegrationComputeInstances(t *testing.T) {
	var tfTmpl = fmt.Sprintf(`
resource "yandex_vpc_network" "cq-instance-test-net-%[1]s" {
  name = "cq-instance-test-net-%[1]s"
}

resource "yandex_vpc_subnet" "cq-instance-test-subnet-%[1]s" {
  name           = "cq-instance-test-subnet-%[1]s"
  network_id     = yandex_vpc_network.cq-instance-test-net-%[1]s.id
  v4_cidr_blocks = ["10.2.0.0/16"]
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
`, suffix)
	testIntegrationHelper(t, resources.ComputeInstances(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name:   "yandex_compute_instances",
			Filter: IdentityFilter,
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name": fmt.Sprintf("cq-instance-test-instance-%s", suffix),
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "yandex_compute_instance_network_interfaces",
					ForeignKeyName: "instance_id",
					Filter:         IdentityFilter,
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
						},
					},
				},
			},
		}
	}, tfTmpl)
}
