package main

import (
	"github.com/cloudquery/plugin-sdk/serve"
	"github.com/yandex-cloud/cq-provider-yandex/resources/plugin"
)

func main() {
	serve.Source(plugin.Plugin())
}
