# Table: yc_baremetal_private_subnets

This table shows data for YC Baremetal Private Subnets.

https://yandex.cloud/docs/baremetal/api-ref/grpc/PrivateSubnet/list#yandex.cloud.baremetal.v1alpha.PrivateSubnet

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
|status|`utf8`|
|zone_id|`utf8`|
|hardware_pool_id|`utf8`|
|vrf_options|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|labels|`json`|