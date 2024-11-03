# Table: yc_mdb_greenplum_clusters

This table shows data for YC Managed Service for Greenplum Clusters.

https://cloud.yandex.ru/docs/managed-greenplum/api-ref/grpc/cluster_service#Cluster1

The primary key for this table is **id**.

## Relations

The following tables depend on yc_mdb_greenplum_clusters:
  - [yc_mdb_greenplum_hosts](yc_mdb_greenplum_hosts.md)

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
|config|`json`|
|description|`utf8`|
|labels|`json`|
|environment|`utf8`|
|monitoring|`json`|
|master_config|`json`|
|segment_config|`json`|
|master_host_count|`int64`|
|segment_host_count|`int64`|
|segment_in_host|`int64`|
|network_id|`utf8`|
|health|`utf8`|
|status|`utf8`|
|maintenance_window|`json`|
|planned_operation|`json`|
|security_group_ids|`list<item: utf8, nullable>`|
|user_name|`utf8`|
|deletion_protection|`bool`|
|host_group_ids|`list<item: utf8, nullable>`|
|cluster_config|`json`|
|cloud_storage|`json`|
|master_host_group_ids|`list<item: utf8, nullable>`|
|segment_host_group_ids|`list<item: utf8, nullable>`|