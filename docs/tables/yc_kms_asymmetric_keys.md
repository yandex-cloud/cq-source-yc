# Table: yc_kms_asymmetric_keys

This table shows data for YC Key Management Service Asymmetric Keys.

https://yandex.cloud/ru/docs/kms/api-ref/grpc/asymmetric_encryption_key_service#AsymmetricEncryptionKey2

The primary key for this table is **id**.

## Relations

The following tables depend on yc_kms_asymmetric_keys:
  - [yc_access_bindings_kms_asymmetric_keys](yc_access_bindings_kms_asymmetric_keys.md)

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
|status|`utf8`|
|encryption_algorithm|`utf8`|
|deletion_protection|`bool`|