# Table: yc_vpc_subnets

This table shows data for YC VPC Subnets.

https://cloud.yandex.ru/docs/vpc/api-ref/grpc/subnet_service#Subnet1

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
|zone_id|`utf8`|
|v4_cidr_blocks|`list<item: utf8, nullable>`|
|v6_cidr_blocks|`list<item: utf8, nullable>`|
|route_table_id|`utf8`|
|dhcp_options|`json`|