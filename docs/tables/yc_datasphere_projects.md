# Table: yc_datasphere_projects

This table shows data for YC DataSphere Projects.

https://cloud.yandex.ru/docs/datasphere/api-ref/grpc/project_service#Project3

The primary key for this table is **id**.

## Relations

This table depends on [yc_datasphere_communities](yc_datasphere_communities.md).

The following tables depend on yc_datasphere_projects:
  - [yc_access_bindings_datasphere_projects](yc_access_bindings_datasphere_projects.md)

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
|settings|`json`|
|limits|`json`|
|community_id|`utf8`|