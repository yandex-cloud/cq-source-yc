package plugin

import (
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

var (
	Version = "Development"
)

func Plugin() *plugins.SourcePlugin {
	allTables := Tables()
	// here you can append custom non-generated tables
	return plugins.NewSourcePlugin(
		"yandex",
		Version,
		allTables,
		client.Configure,
	)
}
