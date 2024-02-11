# Table: yc_mdb_kafka_clusters

This table shows data for YC Managed Service for Kafka Clusters.

https://cloud.yandex.ru/docs/managed-kafka/api-ref/grpc/cluster_service#Cluster1

The primary key for this table is **id**.

## Relations

The following tables depend on yc_mdb_kafka_clusters:
  - [yc_mdb_kafka_topics](yc_mdb_kafka_topics.md)
  - [yc_mdb_kafka_users](yc_mdb_kafka_users.md)

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
|environment|`utf8`|
|monitoring|`json`|
|config|`json`|
|network_id|`utf8`|
|health|`utf8`|
|status|`utf8`|
|security_group_ids|`list<item: utf8, nullable>`|
|host_group_ids|`list<item: utf8, nullable>`|
|deletion_protection|`bool`|
|maintenance_window|`json`|
|planned_operation|`json`|