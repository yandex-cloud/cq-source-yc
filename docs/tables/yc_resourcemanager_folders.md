# Table: yc_resourcemanager_folders

This table shows data for YC Resource Manager Folders.

https://cloud.yandex.ru/docs/resource-manager/api-ref/grpc/folder_service#Folder1

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|id (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|description|`utf8`|
|labels|`json`|
|status|`utf8`|