# Table: yc_mdb_clickhouse_users

This table shows data for YC Managed Service for Clickhouse Users.

https://cloud.yandex.ru/docs/managed-clickhouse/api-ref/grpc/user_service#User1

The composite primary key for this table is (**name**, **cluster_id**).

## Relations

This table depends on [yc_mdb_clickhouse_clusters](yc_mdb_clickhouse_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|name (PK)|`utf8`|
|cluster_id (PK)|`utf8`|
|permissions|`json`|
|settings|`json`|
|quotas|`json`|