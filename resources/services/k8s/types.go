package k8s

import (
	"github.com/yandex-cloud/go-genproto/yandex/cloud/k8s/v1"
)

type Cluster struct {
	*k8s.Cluster
	InternetGateway       *k8s.Cluster_GatewayIpv4Address
	NetworkImplementation *k8s.Cluster_Cilium
}
