# Table: yc_quotamanager_quota_limits

This table shows data for YC Quota Manager Quota Limits.

https://yandex.cloud/ru/docs/quota-manager/api-ref/grpc/QuotaLimit/list#yandex.cloud.quotamanager.v1.QuotaLimit

The composite primary key for this table is (**resource_id**, **quota_id**).

## Relations

This table depends on [yc_quotamanager_quota_services](yc_quotamanager_quota_services.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|resource_id (PK)|`utf8`|
|resource_type|`utf8`|
|quota_id (PK)|`utf8`|
|limit|`json`|
|usage|`json`|
