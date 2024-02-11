# Table: yc_kubernetes_clusters

This table shows data for YC Kubernetes Clusters.

https://cloud.yandex.ru/docs/managed-kubernetes/api-ref/grpc/cluster_service#Cluster1

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|id (PK)|`utf8`|
|folder_id|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|description|`utf8`|
|labels|`json`|
|status|`utf8`|
|health|`utf8`|
|network_id|`utf8`|
|master|`json`|
|ip_allocation_policy|`json`|
|internet_gateway|`json`|
|service_account_id|`utf8`|
|node_service_account_id|`utf8`|
|release_channel|`utf8`|
|network_policy|`json`|
|kms_provider|`json`|
|log_group_id|`utf8`|
|network_implementation|`json`|