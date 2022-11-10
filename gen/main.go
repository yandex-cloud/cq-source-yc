package main

import (
	"log"

	"github.com/yandex-cloud/cq-provider-yandex/gen/recipes"
)

func main() {
	for _, f := range []func() []*recipes.Resource{
		recipes.AccessBindings,
		recipes.CertificateManager,
		recipes.Compute,
		recipes.ContainerRegistry,
		recipes.IAM,
		recipes.K8s,
		recipes.KMS,
		recipes.OrganizationManager,
		recipes.ResourceManager,
		recipes.ApiGateway,
		recipes.Storage,
		recipes.VPC,
	} {
		for _, resource := range f() {
			if err := resource.Generate(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
