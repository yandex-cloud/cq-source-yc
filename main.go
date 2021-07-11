package main

import (
	"github.com/cloudquery/cq-provider-sdk/serve"
	"github.com/yandex-cloud/cq-provider-yandex/resources"
)

func main() {
	serve.Serve(&serve.Options{
		Name:     "yandex",
		Provider: resources.Provider(),
	})
}
