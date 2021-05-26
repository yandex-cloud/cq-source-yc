package main

import (
	"github.com/GennadySpb/cq-provider-yandex/resources"
	"github.com/cloudquery/cq-provider-sdk/serve"
)

func main() {
	serve.Serve(&serve.Options{
		Name:     "yandex",
		Provider: resources.Provider(),
	})
}
