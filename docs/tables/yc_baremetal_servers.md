# Table: yc_baremetal_servers

This table shows data for YC Baremetal Servers.

https://yandex.cloud/docs/baremetal/api-ref/grpc/Server/list#yandex.cloud.baremetal.v1alpha.Server

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|configuration|`json`|
|id (PK)|`utf8`|
|folder_id|`utf8`|
|name|`utf8`|
|description|`utf8`|
|zone_id|`utf8`|
|hardware_pool_id|`utf8`|
|status|`utf8`|
|os_settings|`json`|
|network_interfaces|`json`|
|configuration_id|`utf8`|
|disks|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|labels|`json`|