# Table: yc_mdb_kafka_hosts

This table shows data for YC Managed Service for Kafka Hosts.

https://yandex.cloud/ru/docs/managed-kafka/api-ref/grpc/Cluster/listHosts#yandex.cloud.mdb.kafka.v1.Host

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
|zone_id|`utf8`|
|role|`utf8`|
|resources|`json`|
|health|`utf8`|
|subnet_id|`utf8`|
|assign_public_ip|`bool`|