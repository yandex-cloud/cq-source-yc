# Table: yc_compute_disks

This table shows data for YC Compute Disks.

https://cloud.yandex.ru/docs/compute/api-ref/grpc/disk_service#Disk1

The primary key for this table is **id**.

## Relations

The following tables depend on yc_compute_disks:
  - [yc_compute_disks_images](yc_compute_disks_images.md)

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
|type_id|`utf8`|
|zone_id|`utf8`|
|size|`int64`|
|block_size|`int64`|
|product_ids|`list<item: utf8, nullable>`|
|status|`utf8`|
|source|`json`|
|instance_ids|`list<item: utf8, nullable>`|
|disk_placement_policy|`json`|