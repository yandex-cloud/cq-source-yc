package plugin

import (
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var (
	Name    = "yc"
	Kind    = "source"
	Team    = "yandex-cloud"
	Version = "development"
)

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		Name,
		Version,
		NewClient,
		plugin.WithKind(Kind),
		plugin.WithTeam(Team),
	)
}
