# Table: yc_mdb_mysql_hosts

This table shows data for YC Managed Service for Mysql Hosts.

The composite primary key for this table is (**name**, **cluster_id**).

## Relations

This table depends on [yc_mdb_mysql_clusters](yc_mdb_mysql_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|name (PK)|`utf8`|
|cluster_id (PK)|`utf8`|
|zone_id|`utf8`|
|resources|`json`|
|role|`utf8`|
|health|`utf8`|
|services|`json`|
|subnet_id|`utf8`|
|assign_public_ip|`bool`|
|replication_source|`utf8`|
|backup_priority|`int64`|
|priority|`int64`|