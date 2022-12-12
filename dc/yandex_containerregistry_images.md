# Table: yandex_containerregistry_images



The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|folder_id|String|
|name|String|
|digest|String|
|compressed_size|Int|
|config|JSON|
|layers|JSON|
|tags|StringArray|
|created_at|Timestamp|