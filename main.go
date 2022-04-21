package main

import (
	"github.com/cloudquery/cq-provider-sdk/serve"
	"github.com/yandex-cloud/cq-provider-yandex/resources/provider"
)

func main() {
	serve.Serve(&serve.Options{
		Name:     "yandex",
		Provider: provider.Provider(),
	})
}
