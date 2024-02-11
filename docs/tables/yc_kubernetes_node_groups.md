# Table: yc_kubernetes_node_groups

This table shows data for YC Kubernetes Node Groups.

The primary key for this table is **id**.

## Relations

The following tables depend on yc_kubernetes_node_groups:
  - [yc_kubernetes_nodes](yc_kubernetes_nodes.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|id (PK)|`utf8`|
|cluster_id|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|description|`utf8`|
|labels|`json`|
|status|`utf8`|
|node_template|`json`|
|scale_policy|`json`|
|allocation_policy|`json`|
|deploy_policy|`json`|
|instance_group_id|`utf8`|
|node_version|`utf8`|
|version_info|`json`|
|maintenance_policy|`json`|
|allowed_unsafe_sysctls|`list<item: utf8, nullable>`|
|node_taints|`json`|
|node_labels|`json`|