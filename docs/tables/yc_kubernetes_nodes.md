# Table: yc_kubernetes_nodes

This table shows data for YC Kubernetes Nodes.

The primary key for this table is **cloud_status_id**.

## Relations

This table depends on [yc_kubernetes_node_groups](yc_kubernetes_node_groups.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|node_group_id|`utf8`|
|status|`utf8`|
|spec|`json`|
|cloud_status_id (PK)|`utf8`|
|cloud_status_status|`utf8`|
|cloud_status_status_message|`utf8`|
|kubernetes_status|`json`|