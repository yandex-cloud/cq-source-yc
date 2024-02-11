# Table: yc_datasphere_communities

This table shows data for YC DataSphere Communities.

https://cloud.yandex.ru/docs/datasphere/api-ref/grpc/community_service#Community3

The primary key for this table is **id**.

## Relations

The following tables depend on yc_datasphere_communities:
  - [yc_access_bindings_datasphere_communities](yc_access_bindings_datasphere_communities.md)
  - [yc_datasphere_projects](yc_datasphere_projects.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|description|`utf8`|
|labels|`json`|
|created_by_id|`utf8`|
|organization_id|`utf8`|