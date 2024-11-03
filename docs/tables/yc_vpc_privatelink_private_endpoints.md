# Table: yc_vpc_privatelink_private_endpoints

This table shows data for VPC Private Endpoints.

https://yandex.cloud/ru/docs/vpc/privatelink/api-ref/grpc/PrivateEndpoint/list#yandex.cloud.vpc.v1.privatelink.PrivateEndpoint

The primary key for this table is **id**.

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
|network_id|`utf8`|
|status|`utf8`|
|address|`json`|
|dns_options|`json`|
|service|`json`|