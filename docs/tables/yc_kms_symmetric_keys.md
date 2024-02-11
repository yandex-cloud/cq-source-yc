# Table: yc_kms_symmetric_keys

This table shows data for YC Key Management Service Symmetric Keys.

The primary key for this table is **id**.

## Relations

The following tables depend on yc_kms_symmetric_keys:
  - [yc_access_bindings_kms_symmetric_keys](yc_access_bindings_kms_symmetric_keys.md)

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
|primary_version|`json`|
|default_algorithm|`utf8`|
|rotated_at|`timestamp[us, tz=UTC]`|
|rotation_period|`json`|
|deletion_protection|`bool`|