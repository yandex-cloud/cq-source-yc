# Table: yc_audittrails_trails

This table shows data for YC Audit Trails Trails.

https://yandex.cloud/ru/docs/audit-trails/api-ref/grpc/trail_service#Trail1

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
|updated_at|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|description|`utf8`|
|labels|`json`|
|destination|`json`|
|service_account_id|`utf8`|
|status|`utf8`|
|filter|`json`|
|status_error_message|`utf8`|