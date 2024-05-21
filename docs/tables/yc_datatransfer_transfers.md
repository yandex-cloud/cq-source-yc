# Table: yc_datatransfer_transfers

This table shows data for YC Data Transfer Transfers.

https://yandex.cloud/ru/docs/data-transfer/api-ref/grpc/transfer_service#Transfer

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
|labels|`json`|
|source|`json`|
|target|`json`|
|runtime|`json`|
|status|`utf8`|
|type|`utf8`|
|warning|`utf8`|
|transformation|`json`|
|prestable|`bool`|