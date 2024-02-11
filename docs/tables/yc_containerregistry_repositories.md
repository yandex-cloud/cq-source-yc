# Table: yc_containerregistry_repositories

This table shows data for YC Container Registry Repositories.

https://cloud.yandex.ru/docs/container-registry/api-ref/grpc/repository_service#Repository2

The primary key for this table is **id**.

## Relations

This table depends on [yc_containerregistry_registries](yc_containerregistry_registries.md).

The following tables depend on yc_containerregistry_repositories:
  - [yc_access_bindings_containerregistry_repositories](yc_access_bindings_containerregistry_repositories.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|registry_id|`utf8`|
|name|`utf8`|
|id (PK)|`utf8`|