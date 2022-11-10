# Table: yandex_compute_images


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
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
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|