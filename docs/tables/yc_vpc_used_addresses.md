# Table: yc_vpc_used_addresses

This table shows data for YC VPC Used Addresses.

https://yandex.cloud/ru/docs/vpc/api-ref/grpc/Subnet/listUsedAddresses#yandex.cloud.vpc.v1.UsedAddress

The composite primary key for this table is (**subnet_id**, **address**).

## Relations

This table depends on [yc_vpc_subnets](yc_vpc_subnets.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subnet_id (PK)|`utf8`|
|cloud_id|`utf8`|
|folder_id|`utf8`|
|address (PK)|`utf8`|
|ip_version|`utf8`|
|references|`json`|