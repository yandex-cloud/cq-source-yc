# Table: yandex_vpc_subnets


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
|network_id|String|
|status|Int|
|rules|JSON|
|default_for_network|Bool|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|