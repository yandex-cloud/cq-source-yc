# Table: yc_mdb_opensearch_auth_settings

This table shows data for YC Managed Service for Opensearch Auth Settings.

https://cloud.yandex.ru/docs/managed-opensearch/api-ref/grpc/cluster_service#AuthSettings

The primary key for this table is **cluster_id**.

## Relations

This table depends on [yc_mdb_opensearch_clusters](yc_mdb_opensearch_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|cluster_id (PK)|`utf8`|
|saml_enabled|`bool`|
|saml_idp_entity_id|`utf8`|
|saml_idp_metadata_file|`binary`|
|saml_sp_entity_id|`utf8`|
|saml_dashboards_url|`utf8`|
|saml_roles_key|`utf8`|
|saml_subject_key|`utf8`|
|saml_jwt_default_expiration_timeout|`json`|