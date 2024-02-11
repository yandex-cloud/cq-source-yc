# Table: yc_mdb_kafka_topics

This table shows data for YC Managed Service for Kafka Topics.

https://cloud.yandex.ru/docs/managed-kafka/api-ref/grpc/topic_service#Topic1

The composite primary key for this table is (**name**, **cluster_id**).

## Relations

This table depends on [yc_mdb_kafka_clusters](yc_mdb_kafka_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|name (PK)|`utf8`|
|cluster_id (PK)|`utf8`|
|partitions|`json`|
|replication_factor|`json`|
|topic_config|`json`|