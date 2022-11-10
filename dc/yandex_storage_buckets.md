# Table: yandex_storage_buckets


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|id (PK)|String|
|name|String|
|folder_id|String|
|anonymous_access_flags|JSON|
|default_storage_class|String|
|versioning|Int|
|max_size|Int|
|policy|JSON|
|acl|JSON|
|created_at|Timestamp|
|cors|JSON|
|website_settings|JSON|
|lifecycle_rules|JSON|
|server_side_encryption|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|