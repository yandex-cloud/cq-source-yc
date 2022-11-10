# Table: yandex_vpc_addresses


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
|reserved|Bool|
|used|Bool|
|type|Int|
|ip_version|Int|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|