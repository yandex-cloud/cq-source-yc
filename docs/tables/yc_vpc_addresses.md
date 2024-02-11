# Table: yc_vpc_addresses

This table shows data for YC VPC Addresses.

https://cloud.yandex.ru/docs/vpc/api-ref/grpc/address_service#Address2

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
|address|`json`|
|reserved|`bool`|
|used|`bool`|
|type|`utf8`|
|ip_version|`utf8`|
|deletion_protection|`bool`|