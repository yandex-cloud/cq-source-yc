# Table: yandex_organizationmanager_groups


The primary key for this table is **id**.

## Relations
This table depends on [`yandex_organizationmanager_organizations`](yandex_organizationmanager_organizations.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|id (PK)|String|
|organization_id|String|
|created_at|Timestamp|
|name|String|
|description|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|