# Table: yandex_compute_disks


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
|type_id|String|
|zone_id|String|
|size|Int|
|block_size|Int|
|product_ids|StringArray|
|status|Int|
|instance_ids|StringArray|
|disk_placement_policy|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|