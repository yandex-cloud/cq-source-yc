# Table: yc_access_bindings_datasphere_projects

This table shows data for YC Access Bindings DataSphere Projects.

The composite primary key for this table is (**id**, **role_id**, **subject**).

## Relations

This table depends on [yc_datasphere_projects](yc_datasphere_projects.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|role_id (PK)|`utf8`|
|subject (PK)|`json`|