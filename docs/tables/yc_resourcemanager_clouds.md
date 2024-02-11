# Table: yc_resourcemanager_clouds

This table shows data for YC Resource Manager Clouds.

https://cloud.yandex.ru/docs/resource-manager/api-ref/grpc/cloud_service#Cloud1

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
|organization_id|`utf8`|
|labels|`json`|