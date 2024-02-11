# Table: yc_access_bindings_serverless_containers

This table shows data for YC Access Bindings Serverless Containers.

The composite primary key for this table is (**id**, **role_id**, **subject**).

## Relations

This table depends on [yc_serverless_containers_containers](yc_serverless_containers_containers.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|role_id (PK)|`utf8`|
|subject (PK)|`json`|