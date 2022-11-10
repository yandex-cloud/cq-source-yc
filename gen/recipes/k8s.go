package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	k8s_resource "github.com/yandex-cloud/cq-provider-yandex/resources/services/k8s"
)

func K8s() []*Resource {
	return []*Resource{
		{
			Service:        "k8s",
			SubService:     "clusters",
			Struct:         new(k8s_resource.Cluster),
			SkipFields:     []string{id},
			ExtraColumns:   codegen.ColumnDefinitions{idCol},
			FieldsToUnwrap: []string{"Cluster"},
			Multiplex:      multiplexFolder,
		},
	}
}
