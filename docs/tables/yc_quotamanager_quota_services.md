# Table: yc_quotamanager_quota_services

This table shows data for YC Quota Manager Quota Services.

https://yandex.cloud/ru/docs/quota-manager/api-ref/grpc/QuotaLimit/listServices#yandex.cloud.quotamanager.v1.Service

The composite primary key for this table is (**id**, **resource_type**).

## Relations

The following tables depend on yc_quotamanager_quota_services:
  - [yc_quotamanager_quota_limits](yc_quotamanager_quota_limits.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|resource_type (PK)|`utf8`|
