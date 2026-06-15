# Table: yc_access_policy_bindings_organizationmanager_organizations

This table shows data for YC Access Policy Bindings for Organizations.

https://yandex.cloud/docs/iam/concepts/access-control/#access-policies

The composite primary key for this table is (**id**, **access_policy_template_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|access_policy_template_id (PK)|`utf8`|
|parameters|`json`|