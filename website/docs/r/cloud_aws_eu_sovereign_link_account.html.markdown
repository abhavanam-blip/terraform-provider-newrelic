---
layout: "newrelic"
page_title: "New Relic: newrelic_cloud_aws_eu_sovereign_link_account"
sidebar_current: "docs-newrelic-cloud-resource-aws-eu-sovereign-link-account"
description: |-
  Link an AWS EU Sovereign account to New Relic.
---
# Resource: newrelic_cloud_aws_eu_sovereign_link_account

Use this resource to link an AWS EU Sovereign account to New Relic.
Is the
## Prerequisite

To link an AWS EU Sovereign account to New Relic, you need an AWS EU Sovereign Cloud account. AWS EU Sovereign Cloud is designed to address the specific regulatory needs of European Union public sector organizations and their supporting ecosystem. It is an isolated AWS region designed to host sensitive data and regulated workloads in the cloud, helping customers support their EU compliance requirements.

To pull data from AWS EU Sovereign, follow the [steps outlined here](https://docs.newrelic.com/docs/infrastructure/amazon-integrations/aws-integrations/aws-eu-sovereign-cloud-integrations/).

## Example Usage

```hcl
resource "newrelic_cloud_aws_eu_sovereign_link_account" "foo" {
  account_id             = 1234567
  name                   = "My New Relic - AWS EU Sovereign Linked Account"
  metric_collection_mode = "PULL"
  arn                    = "arn:aws-eusc:iam::123456789012:role/NewRelicInfrastructure-Integrations"
}
```

## Argument Reference

The following arguments are supported:

- `account_id` - (Optional) The New Relic account ID to operate on. This allows the user to override the `account_id` attribute set on the provider. Defaults to the environment variable `NEW_RELIC_ACCOUNT_ID`, if not specified in the configuration.
- `name` - (Required) The name/identifier of the AWS EU Sovereign - New Relic 'linked' account.
- `metric_collection_mode` - (Optional) The mode by which metric data is to be collected from the linked AWS EU Sovereign account. Defaults to `PULL`, if not specified in the configuration.
  - Use `PUSH` for Metric Streams and `PULL` for API Polling based metric collection respectively.
- `arn` - (Required) The Amazon Resource Name (ARN) of the IAM role.

-> **NOTE:** Altering the `account_id` (or) `metric_collection_mode` of an already applied `newrelic_cloud_aws_eu_sovereign_link_account` resource shall trigger a recreation of the resource, instead of an update.

-> **NOTE:** This resource requires the New Relic provider to be configured with `region = "EU"` or the `NEW_RELIC_REGION=EU` environment variable.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` - The ID of the AWS EU Sovereign linked account.

## Import

Linked AWS EU Sovereign accounts can be imported using the `id`, e.g.

```bash
$ terraform import newrelic_cloud_aws_eu_sovereign_link_account.foo <id>
```