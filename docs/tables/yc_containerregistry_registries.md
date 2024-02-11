# Table: yc_containerregistry_registries

This table shows data for YC Container Registry Registries.

https://cloud.yandex.ru/docs/container-registry/api-ref/grpc/registry_service#Registry1

The primary key for this table is **id**.

## Relations

The following tables depend on yc_containerregistry_registries:
  - [yc_access_bindings_containerregistry_registries](yc_access_bindings_containerregistry_registries.md)
  - [yc_containerregistry_repositories](yc_containerregistry_repositories.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|id (PK)|`utf8`|
|folder_id|`utf8`|
|name|`utf8`|
|status|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|labels|`json`|