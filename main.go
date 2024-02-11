package main

import (
	"context"
	"log"

	"github.com/cloudquery/plugin-sdk/v4/serve"
	"github.com/yandex-cloud/cq-source-yc/plugin"
)

func main() {
	p := serve.Plugin(plugin.Plugin())
	err := p.Serve(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
