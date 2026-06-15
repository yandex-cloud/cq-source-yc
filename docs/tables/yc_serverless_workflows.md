# Table: yc_serverless_workflows

This table shows data for YC Serverless Workflows.

https://yandex.cloud/docs/serverless-integrations/workflows/api-ref/grpc/Workflow/list#yandex.cloud.serverless.workflows.v1.Workflow

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|id (PK)|`utf8`|
|folder_id|`utf8`|
|specification|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|description|`utf8`|
|labels|`json`|
|status|`utf8`|
|log_options|`json`|
|network_id|`utf8`|
|service_account_id|`utf8`|
|express|`bool`|
|schedule|`json`|
|is_public|`bool`|
|execution_url|`utf8`|