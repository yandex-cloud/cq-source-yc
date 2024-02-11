# Table: yc_mdb_postgresql_hosts

This table shows data for YC Managed Service for Postgresql Hosts.

https://cloud.yandex.ru/docs/managed-postgresql/api-ref/grpc/cluster_service#Host

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
|zone_id|`utf8`|
|resources|`json`|
|role|`utf8`|
|health|`utf8`|
|services|`json`|
|subnet_id|`utf8`|
|replication_source|`utf8`|
|priority|`json`|
|config|`json`|
|assign_public_ip|`bool`|
|replica_type|`utf8`|