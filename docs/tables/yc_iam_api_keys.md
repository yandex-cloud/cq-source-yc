# Table: yc_iam_api_keys

This table shows data for YC IAM API Keys.

https://cloud.yandex.ru/docs/iam/api-ref/grpc/api_key_service#ApiKey

The primary key for this table is **id**.

## Relations

This table depends on [yc_iam_service_accounts](yc_iam_service_accounts.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|id (PK)|`utf8`|
|service_account_id|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|last_used_at|`timestamp[us, tz=UTC]`|
|scope|`utf8`|
|scopes|`list<item: utf8, nullable>`|
|expires_at|`timestamp[us, tz=UTC]`|