# Table: yandex_compute_images



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
|created_at|Timestamp|
|name|String|
|description|String|
|labels|JSON|
|family|String|
|storage_size|Int|
|min_disk_size|Int|
|product_ids|StringArray|
|status|Int|
|os|JSON|
|pooled|Bool|