# Table: yc_mdb_postgresql_databases

This table shows data for YC Managed Service for Postgresql Databases.

https://cloud.yandex.ru/docs/managed-postgresql/api-ref/grpc/database_service#Database1

The composite primary key for this table is (**name**, **cluster_id**).

## Relations

This table depends on [yc_mdb_postgresql_clusters](yc_mdb_postgresql_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|name (PK)|`utf8`|
|cluster_id (PK)|`utf8`|
|owner|`utf8`|
|lc_collate|`utf8`|
|lc_ctype|`utf8`|
|extensions|`json`|
|template_db|`utf8`|
|deletion_protection|`json`|