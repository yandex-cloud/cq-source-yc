# Table: yandex_compute_instances


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
|zone_id|String|
|platform_id|String|
|resources|JSON|
|status|Int|
|metadata|JSON|
|metadata_options|JSON|
|boot_disk|JSON|
|secondary_disks|JSON|
|local_disks|JSON|
|filesystems|JSON|
|network_interfaces|JSON|
|fqdn|String|
|scheduling_policy|JSON|
|service_account_id|String|
|network_settings|JSON|
|placement_policy|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|