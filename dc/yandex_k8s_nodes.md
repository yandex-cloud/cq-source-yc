# Table: yandex_k8s_nodes


The primary key for this table is **id**.

## Relations
This table depends on [`yandex_k8s_node_groups`](yandex_k8s_node_groups.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|id (PK)|String|
|status|Int|
|spec|JSON|
|cloud_status|JSON|
|kubernetes_status|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|