<p align="center">
<a href="https://cloudquery.io">
<img alt="cloudquery logo" width=75% src="https://github.com/cloudquery/cloudquery/raw/main/docs/images/logo.png" />
</a>
</p>

CloudQuery Yandex.Cloud Provider ![BuildStatus](https://img.shields.io/github/workflow/status/yandex-cloud/cq-provider-yandex/test?style=flat-square) ![License](https://img.shields.io/github/license/cloudquery/cloudquery?style=flat-square)
==================================

This [CloudQuery](https://github.com/cloudquery/cloudquery)
provider transforms Yandex.Cloud resources to relational and graph databases.

### Credentials
You may specify either [`YC_TOKEN`](https://cloud.yandex.com/en-ru/docs/iam/operations/iam-token/create) or [`YC_SERVICE_ACCOUNT_KEY_FILE`](https://cloud.yandex.com/en-ru/docs/iam/operations/api-key/create) as environment variables. `YC_SERVICE_ACCOUNT_KEY_FILE` have higher priority, which means if you specified both of variables, provider will use `YC_SERVICE_ACCOUNT_KEY_FILE`.

### Config Example
``` yaml
kind: source
spec:
  # Source spec section
  name: "yandex"
  version: "v0.3.0"
  path: "yandex-cloud/yandex"
  destinations: ["postgresql"]

  skip_tables:
    - yandex_storage_buckets
  
  # yc provider spec
  spec:
    folder_ids: 
      - abcdefj1234567890xyz
    # cloud_ids:
    # folder_ids:
    # endpoint:
    # folder_filter:
```

For more parameters see [plugin-sdk](https://github.com/cloudquery/plugin-sdk/blob/main/specs/source.go#L17)

## What is CloudQuery

CloudQuery pulls, normalize, expose and monitor your cloud infrastructure and SaaS apps as SQL or Graph(Neo4j) database.
This abstracts various scattered APIs enabling you to define security,governance,cost and compliance policies with SQL
 or [Cypher(Neo4j)](https://neo4j.com/developer/cypher/).

cloudquery can be easily extended to more resources and SaaS providers (open an [Issue](https://github.com/cloudquery/cloudquery/issues)).

cloudquery comes with built-in policy packs such as: [AWS CIS](#running-policy-packs) (more is coming!).

Think about cloudquery as a compliance-as-code tool inspired by tools like [osquery](https://github.com/osquery/osquery)
and [terraform](https://github.com/hashicorp/terraform), cool right?

### Links
* Homepage: https://cloudquery.io
* Releases: https://github.com/cloudquery/cloudquery/releases
* Documentation: https://docs.cloudquery.io
* Schema explorer (schemaspy): https://schema.cloudquery.io/
* Database Configuration: https://docs.cloudquery.io/database-configuration

#### Note about previous history
CloudQuery providers where split into standalone repository. Previous history
is available at the main [cloudquery/cloudquery](https://github.com/cloudquery/cloudquery) repository.

