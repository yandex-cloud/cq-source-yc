# Table: yandex_access_bindings_by_organization


The composite primary key for this table is (**organization_id**, **role_id**, **subject_id**).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|organization_id (PK)|String|
|role_id (PK)|String|
|subject_id (PK)|String|
|subject_type|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|