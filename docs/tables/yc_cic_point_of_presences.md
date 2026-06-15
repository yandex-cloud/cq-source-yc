# Table: yc_cic_point_of_presences

This table shows data for YC Interconnect Point Of Presences.

https://yandex.cloud/docs/interconnect/api-ref/grpc/PointOfPresence/list#yandex.cloud.cic.v1.PointOfPresence

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|location_address|`utf8`|
|connection_points|`list<item: utf8, nullable>`|