# Table: yc_cic_trunk_connections

This table shows data for YC Interconnect Trunk Connections.

https://yandex.cloud/ru/docs/interconnect/api-ref/grpc/TrunkConnection/list#yandex.cloud.cic.v1.TrunkConnection

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|description|`utf8`|
|folder_id|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|joint|`json`|
|point_of_presence_id|`utf8`|
|capacity|`utf8`|
|labels|`json`|
|status|`utf8`|
|deletion_protection|`bool`|