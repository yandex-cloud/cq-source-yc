# Table: yandex_k8s_node_groups


The primary key for this table is **id**.

## Relations
The following tables depend on `yandex_k8s_node_groups`:
  - [`yandex_k8s_nodes`](yandex_k8s_nodes.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|id (PK)|String|
|cluster_id|String|
|created_at|Timestamp|
|name|String|
|description|String|
|labels|JSON|
|status|Int|
|node_template|JSON|
|scale_policy|JSON|
|allocation_policy|JSON|
|deploy_policy|JSON|
|instance_group_id|String|
|node_version|String|
|version_info|JSON|
|maintenance_policy|JSON|
|allowed_unsafe_sysctls|StringArray|
|node_taints|JSON|
|node_labels|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|