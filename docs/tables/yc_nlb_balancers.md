# Table: yc_nlb_balancers

This table shows data for YC Network Load Balancer Balancers.

https://cloud.yandex.ru/docs/network-load-balancer/api-ref/grpc/network_load_balancer_service#NetworkLoadBalancer1

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
|status|`utf8`|
|type|`utf8`|
|session_affinity|`utf8`|
|listeners|`json`|
|attached_target_groups|`json`|
|deletion_protection|`bool`|