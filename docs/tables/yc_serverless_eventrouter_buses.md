# Table: yc_serverless_eventrouter_buses

This table shows data for YC Serverless Eventrouter Buses.

https://yandex.cloud/docs/serverless-integrations/eventrouter/api-ref/grpc/Bus/list#yandex.cloud.serverless.eventrouter.v1.Bus

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
|deletion_protection|`bool`|
|status|`utf8`|
|logging_enabled|`bool`|
|log_options|`json`|