# Table: yc_organizationmanager_group_members

This table shows data for YC Cloud Organization Group Members.

https://yandex.cloud/ru/docs/organization/api-ref/grpc/Group/listMembers#yandex.cloud.organizationmanager.v1.GroupMember

The composite primary key for this table is (**id**, **subject_id**).

## Relations

This table depends on [yc_organizationmanager_groups](yc_organizationmanager_groups.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|subject_id (PK)|`utf8`|
|subject_type|`utf8`|