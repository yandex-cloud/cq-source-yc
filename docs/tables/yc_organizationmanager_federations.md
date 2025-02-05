# Table: yc_organizationmanager_federations

This table shows data for YC Cloud Organization Federations.

https://yandex.cloud/ru/docs/organization/saml/api-ref/grpc/Federation/list#yandex.cloud.organizationmanager.v1.saml.Federation

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|organization_id|`utf8`|
|name|`utf8`|
|description|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|cookie_max_age|`json`|
|auto_create_account_on_login|`bool`|
|issuer|`utf8`|
|sso_binding|`utf8`|
|sso_url|`utf8`|
|security_settings|`json`|
|case_insensitive_name_ids|`bool`|
|labels|`json`|