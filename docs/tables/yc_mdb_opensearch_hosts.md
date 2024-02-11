# Table: yc_mdb_opensearch_hosts

This table shows data for YC Managed Service for Opensearch Hosts.

https://cloud.yandex.ru/docs/managed-opensearch/api-ref/grpc/cluster_service#Host

The composite primary key for this table is (**name**, **cluster_id**).

## Relations

This table depends on [yc_mdb_opensearch_clusters](yc_mdb_opensearch_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|name (PK)|`utf8`|
|cluster_id (PK)|`utf8`|
|zone_id|`utf8`|
|resources|`json`|
|type|`utf8`|
|health|`utf8`|
|subnet_id|`utf8`|
|assign_public_ip|`bool`|
|system|`json`|
|node_group|`utf8`|
|roles|`list<item: int64, nullable>`|