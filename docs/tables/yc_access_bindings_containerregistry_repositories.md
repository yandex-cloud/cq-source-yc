# Table: yc_access_bindings_containerregistry_repositories

This table shows data for YC Access Bindings Container Registry Repositories.

The composite primary key for this table is (**id**, **role_id**, **subject**).

## Relations

This table depends on [yc_containerregistry_repositories](yc_containerregistry_repositories.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|role_id (PK)|`utf8`|
|subject (PK)|`json`|