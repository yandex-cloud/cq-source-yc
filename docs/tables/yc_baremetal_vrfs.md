# Table: yc_baremetal_vrfs

This table shows data for YC Baremetal VRFs.

https://yandex.cloud/docs/baremetal/api-ref/grpc/Vrf/list#yandex.cloud.baremetal.v1alpha.Vrf

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
|status|`utf8`|
|static_routes|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|labels|`json`|