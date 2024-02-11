# Table: yc_serverless_apigateway_openapi_specs

This table shows data for YC Serverless API Gateway Openapi Specs.

https://cloud.yandex.ru/docs/api-gateway/apigateway/api-ref/grpc/apigateway_service#GetOpenapiSpecResponse

The primary key for this table is **id**.

## Relations

This table depends on [yc_serverless_apigateway_gateways](yc_serverless_apigateway_gateways.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|openapi_spec|`utf8`|