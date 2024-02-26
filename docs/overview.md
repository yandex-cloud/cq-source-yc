---
name: Yandex Cloud
stage: Preview
title: Yandex Cloud Source Plugin
description: CloudQuery Yandex Cloud Plugin documentation
---

# Yandex Cloud Source Plugin

:badge

The CloudQuery Yandex Cloud plugin pulls configuration out of Yandex Cloud resources and loads it into any supported CloudQuery destination (e.g. PostgreSQL, BigQuery, Snowflake, and [more](/docs/plugins/destinations/overview)).

## Authentication

:authentication

## Query Examples

### Find all compute instances having public IPv4 address

```sql copy
with instances as (select yci.*,
                          iface #>> '{primary_v4_address, one_to_one_nat, address}' address
                   from (select id,
                                folder_id,
                                name,
                                labels,
                                fqdn,
                                jsonb_array_elements(network_interfaces) as iface
                         from yc_compute_instances) yci)
select *
from instances
where address is not null
```

### Find all public object storage buckets (access flags only)

```sql copy
with buckets as (select *,
                        (anonymous_access_flags #> '{list, value}')::bool        as flags_list,
                        (anonymous_access_flags #> '{read, value}')::bool        as flags_read,
                        (anonymous_access_flags #> '{config_read, value}')::bool as flags_config_read
                 from yc_storage_buckets)
select *
from buckets
where flags_list
   or flags_read
   or flags_config_read
```
