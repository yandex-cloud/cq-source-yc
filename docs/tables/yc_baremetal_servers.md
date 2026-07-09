# Table: yc_baremetal_servers

This table shows data for YC Baremetal Servers.

https://yandex.cloud/docs/baremetal/api-ref/grpc/Server/list#yandex.cloud.baremetal.v2.Server

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|id (PK)|`utf8`|
|folder_id|`utf8`|
|name|`utf8`|
|description|`utf8`|
|hardware_pool_id|`utf8`|
|state|`utf8`|
|os_settings|`json`|
|rental_period_start_time|`timestamp[us, tz=UTC]`|
|rental_period_id|`utf8`|
|next_rental_period_id|`utf8`|
|rental_period_end_time|`timestamp[us, tz=UTC]`|
|network_interfaces|`json`|
|prolongation_state|`utf8`|
|disks|`json`|
|configuration|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|annotations|`json`|