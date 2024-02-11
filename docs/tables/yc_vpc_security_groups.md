# Table: yc_vpc_security_groups

This table shows data for YC VPC Security Groups.

https://cloud.yandex.ru/docs/vpc/api-ref/grpc/security_group_service#SecurityGroup1

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
|network_id|`utf8`|
|status|`utf8`|
|rules|`json`|
|default_for_network|`bool`|