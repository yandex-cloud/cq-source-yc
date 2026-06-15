# Table: yc_trino_catalogs

This table shows data for YC Trino Catalogs.

https://yandex.cloud/docs/managed-trino/api-ref/grpc/Catalog/list#yandex.cloud.trino.v1.Catalog

The primary key for this table is **id**.

## Relations

This table depends on [yc_trino_clusters](yc_trino_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|cluster_id|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|connector|`json`|
|description|`utf8`|
|labels|`json`|