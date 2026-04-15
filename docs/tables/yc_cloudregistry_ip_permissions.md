# Table: yc_cloudregistry_ip_permissions

This table shows data for YC Registry IP Permissions.

https://yandex.cloud/ru/docs/cloud-registry/api-ref/grpc/Registry/listIpPermissions#yandex.cloud.cloudregistry.v1.IpPermission

The primary key for this table is **id**.

## Relations

This table depends on [yc_cloudregistry_registries](yc_cloudregistry_registries.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|action|`utf8`|
|ip|`utf8`|