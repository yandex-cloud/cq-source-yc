# Table: yc_ydb_databases

This table shows data for YC Managed Service for YDB Databases.

https://cloud.yandex.ru/docs/ydb/api-ref/grpc/database_service#Database1

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
|status|`utf8`|
|endpoint|`utf8`|
|resource_preset_id|`utf8`|
|storage_config|`json`|
|scale_policy|`json`|
|network_id|`utf8`|
|subnet_ids|`list<item: utf8, nullable>`|
|database_type|`json`|
|assign_public_ips|`bool`|
|location_id|`utf8`|
|labels|`json`|
|backup_config|`json`|
|document_api_endpoint|`utf8`|
|kinesis_api_endpoint|`utf8`|
|kafka_api_endpoint|`utf8`|
|monitoring_config|`json`|
|deletion_protection|`bool`|