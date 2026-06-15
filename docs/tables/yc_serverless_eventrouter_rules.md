# Table: yc_serverless_eventrouter_rules

This table shows data for YC Serverless Eventrouter Rules.

https://yandex.cloud/docs/serverless-integrations/eventrouter/api-ref/grpc/Rule/list#yandex.cloud.serverless.eventrouter.v1.Rule

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|id (PK)|`utf8`|
|bus_id|`utf8`|
|folder_id|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|description|`utf8`|
|labels|`json`|
|filter|`json`|
|targets|`json`|
|deletion_protection|`bool`|
|status|`utf8`|