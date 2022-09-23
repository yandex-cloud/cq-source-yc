package main

import (
	"log"

	"github.com/yandex-cloud/cq-provider-yandex/gen/recipies"
)

func main() {
	for _, f := range []func() []*recipies.Resource{
		recipies.AccessBindings,
		recipies.CertificateManager,
		recipies.Compute,
		recipies.ContainerRegistry,
		recipies.IAM,
		recipies.K8s,
		recipies.KMS,
		recipies.OrganizationManager,
		recipies.ResourceManager,
		recipies.Serverless,
		recipies.Storage,
		recipies.VPC,
	} {
		for _, resource := range f() {
			if err := resource.Generate(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
