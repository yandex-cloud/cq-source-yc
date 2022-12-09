# Table: yandex_vpc_subnets



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
|network_id|String|
|zone_id|String|
|v4_cidr_blocks|StringArray|
|v6_cidr_blocks|StringArray|
|route_table_id|String|
|dhcp_options|JSON|