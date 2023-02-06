# Table: yandex_containerregistry_scan_results



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
|image_id|String|
|scanned_at|Timestamp|
|status|Int|
|vulnerabilities|JSON|