# Table: yc_access_bindings_resourcemanager_folders

This table shows data for YC Access Bindings Resource Manager Folders.

https://cloud.yandex.ru/docs/resource-manager/api-ref/grpc/folder_service#AccessBinding

The composite primary key for this table is (**role_id**, **subject**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id|`utf8`|
|role_id (PK)|`utf8`|
|subject (PK)|`json`|