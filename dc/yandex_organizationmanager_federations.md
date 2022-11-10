# Table: yandex_organizationmanager_federations


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|id (PK)|String|
|organization_id|String|
|name|String|
|description|String|
|created_at|Timestamp|
|cookie_max_age|JSON|
|auto_create_account_on_login|Bool|
|issuer|String|
|sso_binding|Int|
|sso_url|String|
|security_settings|JSON|
|case_insensitive_name_ids|Bool|
|labels|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|