# Table: yc_mdb_kafka_users

This table shows data for YC Managed Service for Kafka Users.

https://cloud.yandex.ru/docs/managed-kafka/api-ref/grpc/user_service#User1

The composite primary key for this table is (**name**, **cluster_id**).

## Relations

This table depends on [yc_mdb_kafka_clusters](yc_mdb_kafka_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|name (PK)|`utf8`|
|cluster_id (PK)|`utf8`|
|permissions|`json`|