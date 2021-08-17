package integration_tests

//import (
//	"fmt"
//	"testing"
//
//	"github.com/Masterminds/squirrel"
//	"github.com/yandex-cloud/cq-provider-yandex/resources"
//
//	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
//)
//
//func TestIntegrationK8SClusters(t *testing.T) {
//	var tfTmpl = fmt.Sprintf(`
//resource "yandex_vpc_network" "cq-cluster-test-net-%[1]s" {
//  name = "cq-cluster-test-net-%[1]s"
//}
//
//resource "yandex_vpc_subnet" "cq-cluster-test-subnet-%[1]s" {
//  network_id     = yandex_vpc_network.cq-cluster-test-net-%[1]s.id
//  v4_cidr_blocks = ["10.2.0.0/16"]
//  name = "cq-cluster-test-subnet-%[1]s"
//}
//
//resource "yandex_kubernetes_cluster" "cq-cluster-test-cluster-%[1]s" {
//  name="cq-cluster-test-cq-cluster-%[1]s"
//
//  node_service_account_id="aje03ief32j856dgufa0"
//
//  service_account_id="aje03ief32j856dgufa0"
//
//  network_id=yandex_vpc_network.cq-cluster-test-net-%[1]s.id
//
//  master {
//    version = "1.19"
//    zonal {
//      zone      = "${yandex_vpc_subnet.cq-cluster-test-subnet-%[1]s.zone}"
//      subnet_id = "${yandex_vpc_subnet.cq-cluster-test-subnet-%[1]s.id}"
//    }
//  }
//}
//`, suffix)
//	testIntegrationHelper(t, resources.K8SClusters(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
//		return providertest.ResourceIntegrationVerification{
//			Name: "yandex_k8s_clusters",
//			Filter: func(sq squirrel.SelectBuilder, _ *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
//				return sq
//			},
//			ExpectedValues: []providertest.ExpectedValue{
//				{
//					Count: 1,
//					Data: map[string]interface{}{
//						"name": fmt.Sprintf("cq-cluster-test-cluster-%[1]s", suffix),
//					},
//				},
//			},
//		}
//	}, tfTmpl)
//}
