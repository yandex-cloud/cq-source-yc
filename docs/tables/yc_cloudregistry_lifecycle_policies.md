# Table: yc_cloudregistry_lifecycle_policies

This table shows data for YC Registry Lifecycle Policies.

https://yandex.cloud/ru/docs/cloud-registry/api-ref/grpc/LifecyclePolicy/list#yandex.cloud.cloudregistry.v1.LifecyclePolicy

The primary key for this table is **id**.

## Relations

This table depends on [yc_cloudregistry_registries](yc_cloudregistry_registries.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|registry_id|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|description|`utf8`|
|rules|`json`|
|state|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|modified_at|`timestamp[us, tz=UTC]`|
|created_by|`utf8`|
|modified_by|`utf8`|