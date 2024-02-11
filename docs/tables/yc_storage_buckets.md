# Table: yc_storage_buckets

This table shows data for YC Object Storage Buckets.

The composite primary key for this table is (**name**, **folder_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|name (PK)|`utf8`|
|folder_id (PK)|`utf8`|
|anonymous_access_flags|`json`|
|default_storage_class|`utf8`|
|versioning|`utf8`|
|max_size|`int64`|
|policy|`json`|
|acl|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|cors|`json`|
|website_settings|`json`|
|lifecycle_rules|`json`|
|tags|`json`|
|object_lock|`json`|
|encryption|`json`|