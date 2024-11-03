# Table: yc_compute_images

This table shows data for YC Compute Images.

https://cloud.yandex.ru/docs/compute/api-ref/grpc/image_service#Image2

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
|family|`utf8`|
|storage_size|`int64`|
|min_disk_size|`int64`|
|product_ids|`list<item: utf8, nullable>`|
|status|`utf8`|
|os|`json`|
|pooled|`bool`|
|hardware_generation|`json`|
|kms_key|`json`|