# Table: yc_compute_instances

This table shows data for YC Compute Instances.

https://cloud.yandex.ru/docs/compute/api-ref/grpc/instance_service#Instance

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|id (PK)|`utf8`|
|folder_id|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|description|`utf8`|
|labels|`json`|
|zone_id|`utf8`|
|platform_id|`utf8`|
|resources|`json`|
|status|`utf8`|
|metadata|`json`|
|metadata_options|`json`|
|boot_disk|`json`|
|secondary_disks|`json`|
|local_disks|`json`|
|filesystems|`json`|
|network_interfaces|`json`|
|serial_port_settings|`json`|
|gpu_settings|`json`|
|fqdn|`utf8`|
|scheduling_policy|`json`|
|service_account_id|`utf8`|
|network_settings|`json`|
|placement_policy|`json`|
|host_group_id|`utf8`|
|host_id|`utf8`|
|maintenance_policy|`utf8`|
|maintenance_grace_period|`json`|