# Table: yandex_containerregistry_images


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|id (PK)|String|
|folder_id|String|
|name|String|
|digest|String|
|compressed_size|Int|
|config|JSON|
|layers|JSON|
|tags|StringArray|
|created_at|Timestamp|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|