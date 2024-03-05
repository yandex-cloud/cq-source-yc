package main

import (
	"context"
	"log"
	"os"

	"github.com/cloudquery/plugin-sdk/v4/serve"
	"github.com/yandex-cloud/cq-source-yc/plugin"
)

func main() {
	options := []serve.PluginOption{}
	if dsn := os.Getenv("SENTRY_DSN"); dsn != "" {
		options = append(options, serve.WithPluginSentryDSN(dsn))
	}

	p := serve.Plugin(plugin.Plugin(), options...)

	err := p.Serve(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
