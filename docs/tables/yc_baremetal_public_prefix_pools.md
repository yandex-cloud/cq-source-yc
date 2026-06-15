# Table: yc_baremetal_public_prefix_pools

This table shows data for YC Baremetal Public Prefix Pools.

https://yandex.cloud/docs/baremetal/api-ref/grpc/PublicPrefixPool/list#yandex.cloud.baremetal.v1alpha.PublicPrefixPool

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
|cidr|`utf8`|
|min_available_prefix|`int64`|
|created_at|`timestamp[us, tz=UTC]`|
|labels|`json`|