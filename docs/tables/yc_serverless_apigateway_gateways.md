# Table: yc_serverless_apigateway_gateways

This table shows data for YC Serverless API Gateway Gateways.

https://cloud.yandex.ru/docs/api-gateway/apigateway/api-ref/grpc/apigateway_service#ApiGateway1

The primary key for this table is **id**.

## Relations

The following tables depend on yc_serverless_apigateway_gateways:
  - [yc_serverless_apigateway_openapi_specs](yc_serverless_apigateway_openapi_specs.md)

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
|status|`utf8`|
|domain|`utf8`|
|log_group_id|`utf8`|
|attached_domains|`json`|
|connectivity|`json`|
|log_options|`json`|
|variables|`json`|
|canary|`json`|
|execution_timeout|`json`|