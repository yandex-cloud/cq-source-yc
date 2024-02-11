# Table: yc_vpc_networks

This table shows data for YC VPC Networks.

https://cloud.yandex.ru/docs/vpc/api-ref/grpc/network_service#Network1

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
|default_security_group_id|`utf8`|