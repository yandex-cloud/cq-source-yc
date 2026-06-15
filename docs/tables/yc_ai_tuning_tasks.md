# Table: yc_ai_tuning_tasks

This table shows data for YC AI Tuning Tasks.

https://yandex.cloud/docs/foundation-models/tuning/api-ref/grpc/Tuning/list#yandex.cloud.ai.tuning.v1.TuningTask

The primary key for this table is **task_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|task_id (PK)|`utf8`|
|operation_id|`utf8`|
|status|`utf8`|
|folder_id|`utf8`|
|created_by|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|started_at|`timestamp[us, tz=UTC]`|
|finished_at|`timestamp[us, tz=UTC]`|
|source_model_uri|`utf8`|
|target_model_uri|`utf8`|
|name|`utf8`|
|description|`utf8`|
|labels|`json`|