# Table: yc_baremetal_configurations

This table shows data for YC Baremetal Configurations.

https://yandex.cloud/docs/baremetal/api-ref/grpc/Configuration/list#yandex.cloud.baremetal.v1alpha.Configuration

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|memory_gib|`int64`|
|cpu|`json`|
|disk_drives|`json`|
|network_capacity_gbps|`int64`|
|cpu_num|`int64`|
|network_interfaces|`json`|
|mounting_availability|`utf8`|