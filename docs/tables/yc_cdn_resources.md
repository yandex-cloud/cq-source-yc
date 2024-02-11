# Table: yc_cdn_resources

This table shows data for YC CDN Resources.

https://cloud.yandex.ru/docs/cdn/api-ref/grpc/resource_service#Resource1

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|id (PK)|`utf8`|
|folder_id|`utf8`|
|cname|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|active|`bool`|
|options|`json`|
|secondary_hostnames|`list<item: utf8, nullable>`|
|origin_group_id|`int64`|
|origin_group_name|`utf8`|
|origin_protocol|`utf8`|
|ssl_certificate|`json`|
|labels|`json`|