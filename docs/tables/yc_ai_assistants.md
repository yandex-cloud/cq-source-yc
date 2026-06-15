# Table: yc_ai_assistants

This table shows data for YC AI Assistants.

https://yandex.cloud/docs/foundation-models/assistants/api-ref/grpc/Assistant/list#yandex.cloud.ai.assistants.v1.Assistant

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|id (PK)|`utf8`|
|folder_id|`utf8`|
|name|`utf8`|
|description|`utf8`|
|created_by|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_by|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|
|expiration_config|`json`|
|expires_at|`timestamp[us, tz=UTC]`|
|labels|`json`|
|model_uri|`utf8`|
|instruction|`utf8`|
|prompt_truncation_options|`json`|
|completion_options|`json`|
|tools|`json`|
|response_format|`json`|