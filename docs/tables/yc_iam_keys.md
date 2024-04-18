# Table: yc_iam_keys

This table shows data for YC IAM Keys.

https://cloud.yandex.ru/docs/iam/api-ref/grpc/key_service#Key1

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
|subject|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|key_algorithm|`utf8`|
|public_key|`utf8`|
|last_used_at|`timestamp[us, tz=UTC]`|