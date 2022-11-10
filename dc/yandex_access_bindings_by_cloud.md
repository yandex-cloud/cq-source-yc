# Table: yandex_access_bindings_by_cloud


The composite primary key for this table is (**cloud_id**, **role_id**, **subject_id**).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|cloud_id (PK)|String|
|role_id (PK)|String|
|subject_id (PK)|String|
|subject_type|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|