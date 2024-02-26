Authentication is done via IAM tokens/keys. You can read more about it [here](https://cloud.yandex.com/en/docs/iam/concepts/authorization/iam-token)

It is recommended to grant read-only permissions/roles:

- `auditor` role – primitive role which gives access to all configurations, but not the data.
- `*.viewer` roles – per-sevice roles to read resource. Sometimes it is necessary to grant this role, if the service doesn't support `auditor` role.

You can read more about roles [here](https://cloud.yandex.com/en/docs/iam/concepts/access-control/roles)

The plugin will get the following environment variables:

- `YC_SERVICE_ACCOUNT_KEY` – Service Account key (in json format)
- `YC_TOKEN` – IAM or OAuth token

If none of the variables are set, plugin will try to use Compute Metadata API to get IAM token.