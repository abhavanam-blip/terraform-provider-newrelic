---
layout: "newrelic"
page_title: "New Relic: newrelic_cloud_aws_eu_sovereign_link_account"
sidebar_current: "docs-newrelic-resource-cloud-aws-eu-sovereign-link-account"
description: |-
  Link an AWS EU Sovereign account to New Relic.
---

# Resource: newrelic_cloud_aws_eu_sovereign_link_account

Use this resource to link an AWS EU Sovereign account to New Relic.

## Prerequisites

Ensure you have followed [New Relic's AWS EU Sovereign setup documentation](https://docs.newrelic.com/docs/infrastructure/amazon-integrations/aws-integrations/aws-eu-sovereign-cloud-integrations/) to set up your AWS EU Sovereign environment before using this resource.

To use this resource effectively, you'll need:

1. An AWS EU Sovereign account with appropriate permissions
2. A New Relic account with permissions to create cloud integrations
3. An IAM role in your AWS EU Sovereign account with the necessary permissions for New Relic integrations
4. The ARN of the IAM role created for New Relic

## Example Usage

```hcl
resource "newrelic_cloud_aws_eu_sovereign_link_account" "foo" {
  name                   = "my-eu-sovereign-account"
  arn                    = "arn:aws-eu-iso:iam::123456789012:role/NewRelicInfrastructure-Integrations"
  metric_collection_mode = "PULL"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the linked account.
* `arn` - (Required) The ARN of the IAM role.
* `metric_collection_mode` - (Optional) How metrics are collected. Either `PULL` or `PUSH`. Defaults to `PULL`. **Note**: This argument cannot be updated after resource creation.
* `account_id` - (Optional) The account ID for the New Relic account. If omitted, this defaults to the account ID specified in the provider configuration.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the linked account.

## Import

Linked accounts can be imported using the `id`, e.g.

```bash
$ terraform import newrelic_cloud_aws_eu_sovereign_link_account.foo <id>
```

## Notes

* **AWS EU Sovereign Cloud**: This resource is specifically for AWS EU Sovereign Cloud (aws-eu-iso partition) accounts. For regular AWS accounts, use `newrelic_cloud_aws_link_account`. For AWS GovCloud, use `newrelic_cloud_aws_govcloud_link_account`.

* **IAM Role Permissions**: Ensure your AWS EU Sovereign IAM role has the necessary permissions for the integrations you plan to enable. Refer to New Relic's documentation for specific permission requirements.

* **Metric Collection Mode**:
  - `PULL` mode (default): New Relic polls AWS APIs to collect metrics
  - `PUSH` mode: Uses AWS CloudWatch metric streams to push metrics to New Relic

  The collection mode cannot be changed after the account is linked. If you need to change it, you must destroy and recreate the resource.

* **Region Availability**: AWS EU Sovereign regions are limited compared to standard AWS regions. Ensure the services you want to monitor are available in your target EU Sovereign regions.