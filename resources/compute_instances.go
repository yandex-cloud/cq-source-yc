// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------

package resources

import (
	"context"

	"github.com/thoas/go-funk"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
)

func ComputeInstances() *schema.Table {
	return &schema.Table{
		Name:         "yandex_compute_instances",
		Resolver:     fetchComputeInstances,
		Multiplex:    client.FolderMultiplex,
		IgnoreError:  client.IgnoreErrorHandler,
		DeleteFilter: client.DeleteFolderFilter,
		Columns: []schema.Column{
			{
				Name:        "instance_id",
				Type:        schema.TypeString,
				Description: "",
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "folder_id",
				Type:        schema.TypeString,
				Description: "",
				Resolver:    client.ResolveFolderID,
			},
			{
				Name:        "created_at",
				Type:        schema.TypeTimestamp,
				Description: "",
				Resolver:    client.ResolveAsTime,
			},
			{
				Name:        "name",
				Type:        schema.TypeString,
				Description: "Name of the instance. 1-63 characters long.",
				Resolver:    schema.PathResolver("Name"),
			},
			{
				Name:        "description",
				Type:        schema.TypeString,
				Description: "Description of the instance. 0-256 characters long.",
				Resolver:    schema.PathResolver("Description"),
			},
			{
				Name:        "labels",
				Type:        schema.TypeJSON,
				Description: "",
				Resolver:    client.ResolveLabels,
			},
			{
				Name:        "zone_id",
				Type:        schema.TypeString,
				Description: "ID of the availability zone where the instance resides.",
				Resolver:    schema.PathResolver("ZoneId"),
			},
			{
				Name:        "platform_id",
				Type:        schema.TypeString,
				Description: "ID of the hardware platform configuration for the instance.",
				Resolver:    schema.PathResolver("PlatformId"),
			},
			{
				Name:        "resources_memory",
				Type:        schema.TypeBigInt,
				Description: "The amount of memory available to the instance, specified in bytes.",
				Resolver:    schema.PathResolver("Resources.Memory"),
			},
			{
				Name:        "resources_cores",
				Type:        schema.TypeBigInt,
				Description: "The number of cores available to the instance.",
				Resolver:    schema.PathResolver("Resources.Cores"),
			},
			{
				Name:        "resources_core_fraction",
				Type:        schema.TypeBigInt,
				Description: "Baseline level of CPU performance with the ability to burst performance above that baseline level.\n This field sets baseline performance for each core.",
				Resolver:    schema.PathResolver("Resources.CoreFraction"),
			},
			{
				Name:        "resources_gpus",
				Type:        schema.TypeBigInt,
				Description: "The number of GPUs available to the instance.",
				Resolver:    schema.PathResolver("Resources.Gpus"),
			},
			{
				Name:        "status",
				Type:        schema.TypeString,
				Description: "Status of the instance.",
				Resolver:    client.EnumPathResolver("Status"),
			},
			{
				Name:        "metadata",
				Type:        schema.TypeJSON,
				Description: "The metadata `key:value` pairs assigned to this instance. This includes custom metadata and predefined keys.\n\n For example, you may use the metadata in order to provide your public SSH key to the instance.\n For more information, see [Metadata](/docs/compute/concepts/vm-metadata).",
				Resolver:    schema.PathResolver("Metadata"),
			},
			{
				Name:        "boot_disk_mode",
				Type:        schema.TypeString,
				Description: "Access mode to the Disk resource.",
				Resolver:    client.EnumPathResolver("BootDisk.Mode"),
			},
			{
				Name:        "boot_disk_device_name",
				Type:        schema.TypeString,
				Description: "Serial number that is reflected into the /dev/disk/by-id/ tree\n of a Linux operating system running within the instance.\n\n This value can be used to reference the device for mounting, resizing, and so on, from within the instance.",
				Resolver:    schema.PathResolver("BootDisk.DeviceName"),
			},
			{
				Name:        "boot_disk_auto_delete",
				Type:        schema.TypeBool,
				Description: "Specifies whether the disk will be auto-deleted when the instance is deleted.",
				Resolver:    schema.PathResolver("BootDisk.AutoDelete"),
			},
			{
				Name:        "boot_disk_disk_id",
				Type:        schema.TypeString,
				Description: "ID of the disk that is attached to the instance.",
				Resolver:    schema.PathResolver("BootDisk.DiskId"),
			},
			{
				Name:        "fqdn",
				Type:        schema.TypeString,
				Description: "A domain name of the instance. FQDN is defined by the server\n in the format `<hostname>.<region_id>.internal` when the instance is created.\n If the hostname were not specified when the instance was created, FQDN would be `<id>.auto.internal`.",
				Resolver:    schema.PathResolver("Fqdn"),
			},
			{
				Name:        "scheduling_policy_preemptible",
				Type:        schema.TypeBool,
				Description: "True for short-lived compute instances. For more information, see [Preemptible VMs](/docs/compute/concepts/preemptible-vm).",
				Resolver:    schema.PathResolver("SchedulingPolicy.Preemptible"),
			},
			{
				Name:        "service_account_id",
				Type:        schema.TypeString,
				Description: "ID of the service account to use for [authentication inside the instance](/docs/compute/operations/vm-connect/auth-inside-vm).\n To get the service account ID, use a [yandex.cloud.iam.v1.ServiceAccountService.List] request.",
				Resolver:    schema.PathResolver("ServiceAccountId"),
			},
			{
				Name:        "network_settings_type",
				Type:        schema.TypeString,
				Description: "Network Type",
				Resolver:    client.EnumPathResolver("NetworkSettings.Type"),
			},
			{
				Name:        "placement_policy_placement_group_id",
				Type:        schema.TypeString,
				Description: "Placement group ID.",
				Resolver:    schema.PathResolver("PlacementPolicy.PlacementGroupId"),
			},
		},

		Relations: []*schema.Table{
			{
				Name:         "yandex_compute_instance_secondary_disks",
				Resolver:     fetchComputeInstanceSecondaryDisks,
				Multiplex:    client.IdentityMultiplex,
				IgnoreError:  client.IgnoreErrorHandler,
				DeleteFilter: client.DeleteFolderFilter,
				Columns: []schema.Column{
					{
						Name:        "mode",
						Type:        schema.TypeString,
						Description: "Access mode to the Disk resource.",
						Resolver:    client.EnumPathResolver("Mode"),
					},
					{
						Name:        "device_name",
						Type:        schema.TypeString,
						Description: "Serial number that is reflected into the /dev/disk/by-id/ tree\n of a Linux operating system running within the instance.\n\n This value can be used to reference the device for mounting, resizing, and so on, from within the instance.",
						Resolver:    schema.PathResolver("DeviceName"),
					},
					{
						Name:        "auto_delete",
						Type:        schema.TypeBool,
						Description: "Specifies whether the disk will be auto-deleted when the instance is deleted.",
						Resolver:    schema.PathResolver("AutoDelete"),
					},
					{
						Name:        "disk_id",
						Type:        schema.TypeString,
						Description: "ID of the disk that is attached to the instance.",
						Resolver:    schema.PathResolver("DiskId"),
					},
				},
			},
			{
				Name:         "yandex_compute_instance_network_interfaces",
				Resolver:     fetchComputeInstanceNetworkInterfaces,
				Multiplex:    client.IdentityMultiplex,
				IgnoreError:  client.IgnoreErrorHandler,
				DeleteFilter: client.DeleteFolderFilter,
				Columns: []schema.Column{
					{
						Name:        "index",
						Type:        schema.TypeString,
						Description: "The index of the network interface, generated by the server, 0,1,2... etc.\n Currently only one network interface is supported per instance.",
						Resolver:    schema.PathResolver("Index"),
					},
					{
						Name:        "mac_address",
						Type:        schema.TypeString,
						Description: "MAC address that is assigned to the network interface.",
						Resolver:    schema.PathResolver("MacAddress"),
					},
					{
						Name:        "subnet_id",
						Type:        schema.TypeString,
						Description: "ID of the subnet.",
						Resolver:    schema.PathResolver("SubnetId"),
					},
					{
						Name:        "security_group_ids",
						Type:        schema.TypeStringArray,
						Description: "ID's of security groups attached to the interface",
						Resolver:    schema.PathResolver("SecurityGroupIds"),
					},
				},
			},
			{
				Name:         "yandex_compute_instance_placement_policy_host_affinity_rules",
				Resolver:     fetchComputeInstancePlacementPolicyHostAffinityRules,
				Multiplex:    client.IdentityMultiplex,
				IgnoreError:  client.IgnoreErrorHandler,
				DeleteFilter: client.DeleteFolderFilter,
				Columns: []schema.Column{
					{
						Name:        "key",
						Type:        schema.TypeString,
						Description: "Affinity label or one of reserved values - 'yc.hostId', 'yc.hostGroupId'",
						Resolver:    schema.PathResolver("Key"),
					},
					{
						Name:        "op",
						Type:        schema.TypeString,
						Description: "Include or exclude action",
						Resolver:    client.EnumPathResolver("Op"),
					},
					{
						Name:        "values",
						Type:        schema.TypeStringArray,
						Description: "Affinity value or host ID or host group ID",
						Resolver:    schema.PathResolver("Values"),
					},
				},
			},
		},
	}
}

func fetchComputeInstances(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	locations := []string{c.FolderId}

	for _, f := range locations {
		req := &compute.ListInstancesRequest{FolderId: f}
		it := c.Services.Compute.Instance().InstanceIterator(ctx, req)
		for it.Next() {
			res <- it.Value()
		}
	}

	return nil
}

func fetchComputeInstanceSecondaryDisks(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	values := funk.Get(parent.Item, "SecondaryDisks")

	if funk.IsIteratee(values) {
		funk.ForEach(values, func(value interface{}) {
			res <- value
		})
	}

	return nil
}

func fetchComputeInstanceNetworkInterfaces(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	values := funk.Get(parent.Item, "NetworkInterfaces")

	if funk.IsIteratee(values) {
		funk.ForEach(values, func(value interface{}) {
			res <- value
		})
	}

	return nil
}

func fetchComputeInstancePlacementPolicyHostAffinityRules(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	values := funk.Get(parent.Item, "PlacementPolicy.HostAffinityRules")

	if funk.IsIteratee(values) {
		funk.ForEach(values, func(value interface{}) {
			res <- value
		})
	}

	return nil
}
