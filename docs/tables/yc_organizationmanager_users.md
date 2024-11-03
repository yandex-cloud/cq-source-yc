# Table: yc_organizationmanager_users

This table shows data for YC Cloud Organization Users.

https://cloud.yandex.ru/docs/organization/api-ref/grpc/user_service#OrganizationUser

The primary key for this table is **subject_claims_sub**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subject_claims_sub (PK)|`utf8`|
|subject_claims_name|`utf8`|
|subject_claims_given_name|`utf8`|
|subject_claims_family_name|`utf8`|
|subject_claims_preferred_username|`utf8`|
|subject_claims_picture|`utf8`|
|subject_claims_email|`utf8`|
|subject_claims_zoneinfo|`utf8`|
|subject_claims_locale|`utf8`|
|subject_claims_phone_number|`utf8`|
|subject_claims_sub_type|`utf8`|
|subject_claims_federation|`json`|
|subject_claims_last_authenticated_at|`timestamp[us, tz=UTC]`|