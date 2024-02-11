# Table: yc_organizationmanager_organizations

This table shows data for YC Cloud Organization Organizations.

https://cloud.yandex.ru/docs/organization/api-ref/grpc/organization_service#Organization1

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|description|`utf8`|
|title|`utf8`|
|labels|`json`|