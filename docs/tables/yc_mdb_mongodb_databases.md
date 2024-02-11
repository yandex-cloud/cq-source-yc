# Table: yc_mdb_mongodb_databases

This table shows data for YC Managed Service for Mongodb Databases.

https://cloud.yandex.ru/docs/managed-mongodb/api-ref/grpc/database_service#Database1

The composite primary key for this table is (**name**, **cluster_id**).

## Relations

This table depends on [yc_mdb_mongodb_clusters](yc_mdb_mongodb_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|name (PK)|`utf8`|
|cluster_id (PK)|`utf8`|