# Yandex Cloud Source Plugin Configuration Reference

## Example

```yaml copy
kind: source
spec:
  name: "yc"
  registry: cloudquery
  path: yandex-cloud/yc
  version: "v1.0.0"
  destinations: ["postgresql"]
  tables: 
    ["*"]
  spec:
    organization_ids: # sync these organizations only
      - bpf...
    cloud_ids: # sync these clouds only
      - b1g... 
    folder_ids: # sync these folders only
      - b1g... 
---
kind: destination
spec:
  name: "postgresql"
  path: "cloudquery/postgresql"
  registry: "cloudquery"
  version: "v7.1.2"
  spec:
    connection_string: "${PG_CONNECTION_STRING}"
```

## YC Spec

- `organization_ids` (`[]string`, optional, default: empty):
  List of Organization IDs to target. If empty, all available Organization will be targeted.

- `cloud_ids` (`[]string`, optional, default: empty):
  List of Cloud IDs to target. If empty, all available Clouds will be targeted.

- `folder_ids` (`[]string`, optional, default: empty):
  List of Folder IDs to target. If empty, all available Folders will be targeted.

- `debug_grpc` (`bool`, default `false`):
  If true, will log all GRPC calls

- `max_retries` (`int`, default `3`):
  Maxiumum number of retries for YC Client

- `endpoint` (`string`, default `api.cloud.yandex.net:443`):
  Yandex Cloud endpoint

- `concurrency` (`int`, optional, default: `10000`):
  A best effort maximum number of Go routines to use. Lower this number to reduce memory usage.

- `scheduler` (`string`, default: `shuffle`):
  The scheduler to use when determining the priority of resources to sync. Currently, the only supported values are dfs (depth-first search), round-robin and shuffle.
