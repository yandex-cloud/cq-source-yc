# Yandex Cloud Source Plugin

The Yandex Cloud Source Plugin extracts information from [Yandex Cloud API](https://cloud.yandex.ru/ru/docs/api-design-guide/)

> Previous version is available [here](https://github.com/yandex-cloud/cq-source-yandex/tree/v0.3.8)

---

## Configuration

Example configuration:

```yaml
kind: source
spec:
  name: "yc"
  registry: github
  path: yandex-cloud/cq-source-yc
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

## Authentication

Authentication is done via IAM tokens/keys.

You can set following environment variables:

- `YC_SERVICE_ACCOUNT_KEY` – Service Account key (in json format)
- `YC_TOKEN` – IAM or OAuth token

If none of the variables are set, plugin will try to use Compute Metadata API to get IAM token.
