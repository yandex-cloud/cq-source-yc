# Table: yc_baremetal_public_subnets

This table shows data for YC Baremetal Public Subnets.

https://yandex.cloud/docs/baremetal/api-ref/grpc/PublicSubnet/list#yandex.cloud.baremetal.v2.PublicSubnet

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
|hardware_pool_ids|`list<item: utf8, nullable>`|
|type|`utf8`|
|cidr_allocation_method|`json`|
|prefix_length|`int64`|
|cidr|`utf8`|
|dhcp_options|`json`|
|gateway_ip|`utf8`|
|public_prefix_pool_id|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|annotations|`json`|
|deletion_unlock_time|`timestamp[us, tz=UTC]`|