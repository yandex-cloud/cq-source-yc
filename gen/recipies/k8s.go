package recipies

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	k8s_resource "github.com/yandex-cloud/cq-provider-yandex/resources/services/k8s"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/k8s/v1"
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
			Relations:      []string{"NodeGroups()"},
			Multiplex:      multiplexFolder,
		},
		{
			Service:      "k8s",
			SubService:   "node_groups",
			Struct:       new(k8s.NodeGroup),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    multiplexFolder,
		},
		{
			Service:      "k8s",
			SubService:   "nodes",
			Struct:       new(k8s.Node),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    "", // empty multiplex for relation
		},
	}
}
