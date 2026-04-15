# Table: yc_cloudregistry_registries

This table shows data for YC Registry Registries.

https://yandex.cloud/ru/docs/cloud-registry/api-ref/grpc/Registry/list#yandex.cloud.cloudregistry.v1.Registry

The primary key for this table is **id**.

## Relations

The following tables depend on yc_cloudregistry_registries:
  - [yc_access_bindings_cloudregistry_registries](yc_access_bindings_cloudregistry_registries.md)
  - [yc_cloudregistry_ip_permissions](yc_cloudregistry_ip_permissions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|id (PK)|`utf8`|
|folder_id|`utf8`|
|name|`utf8`|
|description|`utf8`|
|kind|`utf8`|
|type|`utf8`|
|status|`utf8`|
|labels|`json`|
|properties|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|modified_at|`timestamp[us, tz=UTC]`|