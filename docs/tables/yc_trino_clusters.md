# Table: yc_trino_clusters

This table shows data for YC Trino Clusters.

https://yandex.cloud/docs/managed-trino/api-ref/grpc/Cluster/list#yandex.cloud.trino.v1.Cluster

The primary key for this table is **id**.

## Relations

The following tables depend on yc_trino_clusters:
  - [yc_trino_catalogs](yc_trino_catalogs.md)

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
|monitoring|`json`|
|trino|`json`|
|health|`utf8`|
|status|`utf8`|
|network|`json`|
|deletion_protection|`bool`|
|service_account_id|`utf8`|
|logging|`json`|
|coordinator_url|`utf8`|
|maintenance_window|`json`|
|planned_operation|`json`|