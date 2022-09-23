// Code generated by codegen; DO NOT EDIT.

package k8s

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func Nodes() *schema.Table {
	return &schema.Table{
		Name:     "yandex_k8s_nodes",
		Resolver: fetchNodes,
		Columns: []schema.Column{
			{
				Name:        "id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Id"),
				Description: `Resource ID`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "status",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "spec",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec"),
			},
			{
				Name:     "cloud_status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CloudStatus"),
			},
			{
				Name:     "kubernetes_status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("KubernetesStatus"),
			},
		},
	}
}