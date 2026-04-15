# Table: yc_cic_public_connections

This table shows data for YC Interconnect Public Connections.

https://yandex.cloud/ru/docs/interconnect/api-ref/grpc/PublicConnection/list#yandex.cloud.cic.v1.PublicConnection

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|id|`utf8`|
|name|`utf8`|
|description|`utf8`|
|folder_id|`utf8`|
|region_id|`utf8`|
|trunk_connection_id|`utf8`|
|vlan_id|`int64`|
|ipv4_peering|`json`|
|ipv4_allowed_service_types|`list<item: utf8, nullable>`|
|ipv4_peer_announced_prefixes|`list<item: utf8, nullable>`|
|labels|`json`|
|status|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|