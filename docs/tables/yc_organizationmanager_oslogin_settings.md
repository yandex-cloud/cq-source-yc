# Table: yc_organizationmanager_oslogin_settings

This table shows data for YC Cloud Organization Oslogin Settings.

https://yandex.cloud/ru/docs/organization/api-ref/grpc/OsLogin/getSettings#yandex.cloud.organizationmanager.v1.OsLoginSettings

The primary key for this table is **organization_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|organization_id (PK)|`utf8`|
|user_ssh_key_settings|`json`|
|ssh_certificate_settings|`json`|