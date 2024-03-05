# Table: yc_access_bindings_ydb_databases

This table shows data for YC Access Bindings Managed Service for YDB Databases.

The composite primary key for this table is (**id**, **role_id**, **subject**).

## Relations

This table depends on [yc_ydb_databases](yc_ydb_databases.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|role_id (PK)|`utf8`|
|subject (PK)|`json`|