# Table: yc_ai_batch_inference_tasks

This table shows data for YC AI Batch Inference Tasks.

https://yandex.cloud/docs/foundation-models/batch/api-ref/grpc/BatchInference/list#yandex.cloud.ai.batch_inference.v1.BatchInferenceTask

The primary key for this table is **task_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|task_id (PK)|`utf8`|
|operation_id|`utf8`|
|folder_id|`utf8`|
|model_uri|`utf8`|
|source_dataset_id|`utf8`|
|request|`json`|
|status|`utf8`|
|result_dataset_id|`utf8`|
|labels|`json`|
|created_by|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|started_at|`timestamp[us, tz=UTC]`|
|finished_at|`timestamp[us, tz=UTC]`|
|errors|`json`|