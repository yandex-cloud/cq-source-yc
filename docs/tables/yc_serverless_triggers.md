# Table: yc_serverless_triggers

This table shows data for YC Serverless Triggers.

https://yandex.cloud/docs/functions/triggers/api-ref/grpc/Trigger/list#yandex.cloud.serverless.triggers.v1.Trigger

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|id (PK)|`utf8`|
|folder_id|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|description|`utf8`|
|labels|`json`|
|rule|`json`|
|status|`utf8`|