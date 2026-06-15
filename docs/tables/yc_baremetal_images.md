# Table: yc_baremetal_images

This table shows data for YC Baremetal Images.

https://yandex.cloud/docs/baremetal/api-ref/grpc/Image/list#yandex.cloud.baremetal.v1alpha.Image

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cloud_id|`utf8`|
|id (PK)|`utf8`|
|folder_id|`utf8`|
|name|`utf8`|
|description|`utf8`|
|checksum|`utf8`|
|status|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|labels|`json`|