# Table: yc_compute_disks_images

This table shows data for YC Compute Disks Images.

This table is exact copy of [yc_compute_images](yc_compute_images.md), but contains images used in [yc_compute_disks](yc_compute_disks.md)
https://cloud.yandex.ru/docs/compute/api-ref/grpc/image_service#Image2

The primary key for this table is **id**.

## Relations

This table depends on [yc_compute_disks](yc_compute_disks.md).

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