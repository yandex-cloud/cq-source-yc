// Code generated by codegen; DO NOT EDIT.

package k8s

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:      "yandex_k8s_clusters",
		Resolver:  fetchClusters,
		Multiplex: client.MultiplexBy(client.Folders),
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
				Name:     "folder_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FolderId"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("CreatedAt"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "status",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "health",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Health"),
			},
			{
				Name:     "network_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NetworkId"),
			},
			{
				Name:     "master",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Master"),
			},
			{
				Name:     "ip_allocation_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IpAllocationPolicy"),
			},
			{
				Name:     "service_account_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceAccountId"),
			},
			{
				Name:     "node_service_account_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NodeServiceAccountId"),
			},
			{
				Name:     "release_channel",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ReleaseChannel"),
			},
			{
				Name:     "network_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NetworkPolicy"),
			},
			{
				Name:     "kms_provider",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("KmsProvider"),
			},
			{
				Name:     "log_group_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LogGroupId"),
			},
			{
				Name:     "internet_gateway",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("InternetGateway"),
			},
			{
				Name:     "network_implementation",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NetworkImplementation"),
			},
		},

		Relations: []*schema.Table{
			NodeGroups(),
		},
	}
}