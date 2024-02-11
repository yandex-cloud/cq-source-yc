# Table: yc_serverless_containers_containers

This table shows data for YC Serverless Containers Containers.

https://cloud.yandex.ru/docs/serverless-containers/containers/api-ref/grpc/container_service#Container1

The primary key for this table is **id**.

## Relations

The following tables depend on yc_serverless_containers_containers:
  - [yc_access_bindings_serverless_containers](yc_access_bindings_serverless_containers.md)

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
|url|`utf8`|
|status|`utf8`|