# Table: yandex_containerregistry_scan_results


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|id (PK)|String|
|folder_id|String|
|image_id|String|
|scanned_at|Timestamp|
|status|Int|
|vulnerabilities|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|