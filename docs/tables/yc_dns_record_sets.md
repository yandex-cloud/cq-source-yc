# Table: yc_dns_record_sets

This table shows data for YC DNS Record Sets.

https://cloud.yandex.ru/docs/dns/api-ref/grpc/dns_zone_service#RecordSet1

The composite primary key for this table is (**zone_id**, **name**, **type**).

## Relations

This table depends on [yc_dns_zones](yc_dns_zones.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|zone_id (PK)|`utf8`|
|name (PK)|`utf8`|
|type (PK)|`utf8`|
|ttl|`int64`|
|data|`list<item: utf8, nullable>`|