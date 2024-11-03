# Table: yc_iam_service_accounts

This table shows data for YC IAM Service Accounts.

https://cloud.yandex.ru/docs/iam/api-ref/grpc/service_account_service#ServiceAccount1

The primary key for this table is **id**.

## Relations

The following tables depend on yc_iam_service_accounts:
  - [yc_access_iam_service_accounts](yc_access_iam_service_accounts.md)
  - [yc_iam_access_keys](yc_iam_access_keys.md)
  - [yc_iam_api_keys](yc_iam_api_keys.md)
  - [yc_iam_keys](yc_iam_keys.md)

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
|last_authenticated_at|`timestamp[us, tz=UTC]`|