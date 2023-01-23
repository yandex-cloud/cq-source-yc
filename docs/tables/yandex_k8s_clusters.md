# Table: yandex_k8s_clusters



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
|status|Int|
|health|Int|
|network_id|String|
|master|JSON|
|ip_allocation_policy|JSON|
|service_account_id|String|
|node_service_account_id|String|
|release_channel|Int|
|network_policy|JSON|
|kms_provider|JSON|
|log_group_id|String|
|internet_gateway|JSON|
|network_implementation|JSON|