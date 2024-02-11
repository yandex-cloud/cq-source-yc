# Table: yc_dns_zones

This table shows data for YC DNS Zones.

https://cloud.yandex.ru/docs/dns/api-ref/grpc/dns_zone_service#DnsZone1

The primary key for this table is **id**.

## Relations

The following tables depend on yc_dns_zones:
  - [yc_dns_record_sets](yc_dns_record_sets.md)

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
|zone|`utf8`|
|private_visibility|`json`|
|public_visibility|`json`|