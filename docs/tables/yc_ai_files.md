# Table: yc_ai_files

This table shows data for YC AI Files.

https://yandex.cloud/docs/foundation-models/assistants/api-ref/grpc/Files/list#yandex.cloud.ai.files.v1.File

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
|mime_type|`utf8`|
|created_by|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_by|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|
|expiration_config|`json`|
|expires_at|`timestamp[us, tz=UTC]`|
|labels|`json`|