# Table: yc_datalens_reports

This table shows data for YC DataLens Reports.

DataLens entries with scope=report. https://yandex.cloud/ru/docs/datalens/operations/api-start

The primary key for this table is **entry_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|organization_id|`utf8`|
|collection_id|`utf8`|
|collection_title|`utf8`|
|created_at|`utf8`|
|created_by|`utf8`|
|data|`json`|
|entry_id (PK)|`utf8`|
|hidden|`bool`|
|is_favorite|`bool`|
|is_locked|`bool`|
|key|`utf8`|
|links|`json`|
|meta|`json`|
|name|`utf8`|
|permissions|`json`|
|published_id|`utf8`|
|saved_id|`utf8`|
|scope|`utf8`|
|type|`utf8`|
|updated_at|`utf8`|
|updated_by|`utf8`|
|workbook_id|`utf8`|
|workbook_title|`utf8`|