# Table: yc_organizationmanager_groups

This table shows data for YC Cloud Organization Groups.

https://yandex.cloud/ru/docs/organization/api-ref/grpc/Group/list#yandex.cloud.organizationmanager.v1.Group

The primary key for this table is **id**.

## Relations

The following tables depend on yc_organizationmanager_groups:
  - [yc_organizationmanager_group_members](yc_organizationmanager_group_members.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|organization_id|`utf8`|
|id (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|description|`utf8`|
|subject_container_id|`utf8`|
|external_id|`utf8`|
|labels|`json`|