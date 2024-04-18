# Table: yc_iam_access_keys

This table shows data for YC IAM Access Keys.

https://cloud.yandex.ru/docs/iam/api-ref/grpc/access_key_service#AccessKey

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
|key_id|`utf8`|
|last_used_at|`timestamp[us, tz=UTC]`|