# Table: yc_baremetal_images

This table shows data for YC Baremetal Images.

https://yandex.cloud/docs/baremetal/api-ref/grpc/BootImage/list#yandex.cloud.baremetal.v2.BootImage

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
|uri|`utf8`|
|checksum|`utf8`|
|state|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|annotations|`json`|