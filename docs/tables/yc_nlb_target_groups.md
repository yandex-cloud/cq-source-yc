# Table: yc_nlb_target_groups

This table shows data for YC Network Load Balancer Target Groups.

https://yandex.cloud/ru/docs/network-load-balancer/api-ref/grpc/TargetGroup/list#yandex.cloud.loadbalancer.v1.TargetGroup

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
|region_id|`utf8`|
|targets|`json`|