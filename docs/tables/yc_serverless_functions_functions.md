# Table: yc_serverless_functions_functions

This table shows data for YC Serverless Functions Functions.

The primary key for this table is **id**.

## Relations

The following tables depend on yc_serverless_functions_functions:
  - [yc_access_bindings_serverless_functions](yc_access_bindings_serverless_functions.md)

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
|http_invoke_url|`utf8`|
|status|`utf8`|