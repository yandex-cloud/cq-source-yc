# Table: yandex_access_bindings_by_folder



The composite primary key for this table is (**folder_id**, **role_id**, **subject_id**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|folder_id (PK)|String|
|role_id (PK)|String|
|subject_id (PK)|String|
|subject_type|String|