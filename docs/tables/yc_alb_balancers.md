# Table: yc_alb_balancers

This table shows data for YC Application Load Balancer Balancers.

https://cloud.yandex.ru/docs/application-load-balancer/api-ref/grpc/load_balancer_service#LoadBalancer1

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|description|`utf8`|
|folder_id|`utf8`|
|labels|`json`|
|status|`utf8`|
|region_id|`utf8`|
|network_id|`utf8`|
|listeners|`json`|
|allocation_policy|`json`|
|log_group_id|`utf8`|
|security_group_ids|`list<item: utf8, nullable>`|
|created_at|`timestamp[us, tz=UTC]`|
|auto_scale_policy|`json`|
|log_options|`json`|