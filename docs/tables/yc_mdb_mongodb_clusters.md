# Table: yc_mdb_mongodb_clusters

This table shows data for YC Managed Service for Mongodb Clusters.

https://cloud.yandex.ru/docs/managed-mongodb/api-ref/grpc/cluster_service#Cluster1

The primary key for this table is **id**.

## Relations

The following tables depend on yc_mdb_mongodb_clusters:
  - [yc_mdb_mongodb_databases](yc_mdb_mongodb_databases.md)
  - [yc_mdb_mongodb_hosts](yc_mdb_mongodb_hosts.md)
  - [yc_mdb_mongodb_users](yc_mdb_mongodb_users.md)

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
|environment|`utf8`|
|monitoring|`json`|
|config|`json`|
|network_id|`utf8`|
|health|`utf8`|
|status|`utf8`|
|sharded|`bool`|
|maintenance_window|`json`|
|planned_operation|`json`|
|security_group_ids|`list<item: utf8, nullable>`|
|deletion_protection|`bool`|