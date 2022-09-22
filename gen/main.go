package main

import (
	"log"

	"github.com/yandex-cloud/cq-provider-yandex/gen/recipies"
)

func main() {
	for _, f := range []func() []*recipies.Resource{
		recipies.AccessBindings,
		recipies.Certificates,
	} {
		for _, resource := range f() {
			if err := resource.Generate(); err != nil {
				log.Fatal(err)
			}
		}
	}
}