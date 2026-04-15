# Table: yc_mdb_postgresql_backups

This table shows data for YC Managed Service for Postgresql Backups.

https://yandex.cloud/ru/docs/managed-postgresql/api-ref/grpc/backup_service#Backup

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|id (PK)|`utf8`|
|folder_id|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|source_cluster_id|`utf8`|
|started_at|`timestamp[us, tz=UTC]`|
|size|`int64`|
|type|`utf8`|
|method|`utf8`|
|journal_size|`int64`|
|status|`utf8`|
|retention_policy_id|`utf8`|
|retention_policy_name|`utf8`|
|retain_until|`timestamp[us, tz=UTC]`|