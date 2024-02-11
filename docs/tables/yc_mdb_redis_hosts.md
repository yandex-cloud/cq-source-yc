# Table: yc_mdb_redis_hosts

This table shows data for YC Managed Service for Redis Hosts.

https://cloud.yandex.ru/docs/managed-redis/api-ref/grpc/cluster_service#Host

The composite primary key for this table is (**name**, **cluster_id**).

## Relations

This table depends on [yc_mdb_redis_clusters](yc_mdb_redis_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|name (PK)|`utf8`|
|cluster_id (PK)|`utf8`|
|zone_id|`utf8`|
|subnet_id|`utf8`|
|resources|`json`|
|role|`utf8`|
|health|`utf8`|
|services|`json`|
|shard_name|`utf8`|
|replica_priority|`json`|
|assign_public_ip|`bool`|