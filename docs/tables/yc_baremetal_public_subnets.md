# Table: yc_baremetal_public_subnets

This table shows data for YC Baremetal Public Subnets.

https://yandex.cloud/docs/baremetal/api-ref/grpc/PublicSubnet/list#yandex.cloud.baremetal.v1alpha.PublicSubnet

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
|zone_id|`utf8`|
|hardware_pool_ids|`list<item: utf8, nullable>`|
|type|`utf8`|
|prefix_length|`int64`|
|cidr|`utf8`|
|dhcp_options|`json`|
|gateway_ip|`utf8`|
|public_prefix_pool_id|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|labels|`json`|
|deletion_unlocked_at|`timestamp[us, tz=UTC]`|