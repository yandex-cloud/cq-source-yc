# Table: yc_cloudregistry_artifacts

This table shows data for YC Registry Artifacts.

https://yandex.cloud/ru/docs/cloud-registry/api-ref/grpc/Registry/listArtifacts#yandex.cloud.cloudregistry.v1.Artifact

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
|path|`utf8`|
|name|`utf8`|
|kind|`utf8`|
|status|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|created_by|`utf8`|
|modified_at|`timestamp[us, tz=UTC]`|
|modified_by|`utf8`|
|properties|`json`|
|content|`json`|