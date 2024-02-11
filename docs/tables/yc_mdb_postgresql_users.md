# Table: yc_mdb_postgresql_users

This table shows data for YC Managed Service for Postgresql Users.

https://cloud.yandex.ru/docs/managed-postgresql/api-ref/grpc/user_service#User1

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
|permissions|`json`|
|conn_limit|`int64`|
|settings|`json`|
|login|`json`|
|grants|`list<item: utf8, nullable>`|
|deletion_protection|`json`|
|user_password_encryption|`utf8`|