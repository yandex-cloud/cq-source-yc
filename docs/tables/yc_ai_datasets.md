# Table: yc_ai_datasets

This table shows data for YC AI Datasets.

https://yandex.cloud/docs/foundation-models/dataset/api-ref/grpc/Dataset/list#yandex.cloud.ai.dataset.v1.DatasetInfo

The primary key for this table is **dataset_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|dataset_id (PK)|`utf8`|
|folder_id|`utf8`|
|name|`utf8`|
|description|`utf8`|
|metadata|`utf8`|
|status|`utf8`|
|task_type|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|rows|`int64`|
|size_bytes|`int64`|
|created_by_id|`utf8`|
|labels|`json`|
|created_by|`utf8`|
|updated_by|`utf8`|
|validation_error|`json`|
|allow_data_log|`bool`|