# Table: yc_access_bindings_kms_asymmetric_keys

This table shows data for YC Access Bindings Key Management Service Asymmetric Keys.

The composite primary key for this table is (**id**, **role_id**, **subject**).

## Relations

This table depends on [yc_kms_asymmetric_keys](yc_kms_asymmetric_keys.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|role_id (PK)|`utf8`|
|subject (PK)|`json`|