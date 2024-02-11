# Table: yc_lockbox_secrets

This table shows data for YC Lockbox Secrets.

https://cloud.yandex.ru/docs/lockbox/api-ref/grpc/secret_service#Secret1

The primary key for this table is **id**.

## Relations

The following tables depend on yc_lockbox_secrets:
  - [yc_access_bindings_lockbox_secrets](yc_access_bindings_lockbox_secrets.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|id (PK)|`utf8`|
|folder_id|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|description|`utf8`|
|labels|`json`|
|kms_key_id|`utf8`|
|status|`utf8`|
|current_version|`json`|
|deletion_protection|`bool`|