# Table: yandex_organizationmanager_organizations



The primary key for this table is **id**.

## Relations

The following tables depend on yandex_organizationmanager_organizations:
  - [yandex_organizationmanager_groups](yandex_organizationmanager_groups.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|created_at|Timestamp|
|name|String|
|description|String|
|title|String|
|labels|JSON|