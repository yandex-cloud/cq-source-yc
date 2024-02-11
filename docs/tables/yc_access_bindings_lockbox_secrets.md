# Table: yc_access_bindings_lockbox_secrets

This table shows data for YC Access Bindings Lockbox Secrets.

https://cloud.yandex.ru/docs/lockbox/api-ref/grpc/secret_service#AccessBinding

The composite primary key for this table is (**id**, **role_id**, **subject**).

## Relations

This table depends on [yc_lockbox_secrets](yc_lockbox_secrets.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|role_id (PK)|`utf8`|
|subject (PK)|`json`|