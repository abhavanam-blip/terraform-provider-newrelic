---
layout: "newrelic"
page_title: "New Relic: newrelic_cloud_aws_eu_sovereign_integrations"
sidebar_current: "docs-newrelic-resource-cloud-aws-eu-sovereign-integrations"
description: |-
  Configure AWS EU Sovereign integrations for a linked AWS EU Sovereign account.
---

# Resource: newrelic_cloud_aws_eu_sovereign_integrations

Use this resource to enable and configure New Relic's integrations with AWS EU Sovereign Cloud services.

## Prerequisites

* You must have an AWS EU Sovereign account linked to New Relic using the `newrelic_cloud_aws_eu_sovereign_link_account` resource
* Your AWS EU Sovereign IAM role must have the appropriate permissions for the services you want to monitor
* The AWS services must be available in your target EU Sovereign regions

See [New Relic's AWS EU Sovereign integration documentation](https://docs.newrelic.com/docs/infrastructure/amazon-integrations/aws-integrations/aws-eu-sovereign-cloud-integrations/) for setup instructions.

## Example Usage

```hcl
resource "newrelic_cloud_aws_eu_sovereign_link_account" "account" {
  name = "my-eu-sovereign-account"
  arn  = "arn:aws-eu-iso:iam::123456789012:role/NewRelicInfrastructure-Integrations"
}

resource "newrelic_cloud_aws_eu_sovereign_integrations" "integrations" {
  linked_account_id = newrelic_cloud_aws_eu_sovereign_link_account.account.id

  cloudtrail {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1"]
  }

  health {
    metrics_polling_interval = 300
  }

  trusted_advisor {
    metrics_polling_interval = 300
  }

  xray {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1", "eu-isob-west-1"]
  }
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Optional) The account ID for the New Relic account. If omitted, this defaults to the account ID specified in the provider configuration.
* `linked_account_id` - (Required) The ID of the AWS EU Sovereign linked account.

The following integration types are supported:

### `cloudtrail`
* `metrics_polling_interval` - (Optional) The data polling interval in seconds.
* `aws_regions` - (Optional) List of AWS EU Sovereign regions that include the resources you want to monitor.

### `health`
* `metrics_polling_interval` - (Optional) The data polling interval in seconds.

### `trusted_advisor`
* `metrics_polling_interval` - (Optional) The data polling interval in seconds.

### `xray`
* `metrics_polling_interval` - (Optional) The data polling interval in seconds.
* `aws_regions` - (Optional) List of AWS EU Sovereign regions that include the resources you want to monitor.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the integration configuration.

## Import

Integrations can be imported using the `linked_account_id`, e.g.

```bash
$ terraform import newrelic_cloud_aws_eu_sovereign_integrations.foo <linked_account_id>
```

## Notes

* **Limited Service Support**: EU Sovereign Cloud integration supports only four AWS services: CloudTrail, Health, Trusted Advisor, and X-Ray. These services support polling mode only via the `metrics_polling_interval` parameter.

* **Polling Mode Only**: Unlike regular AWS integrations that support both PUSH (Metric Streams) and PULL (Polling) collection modes, EU Sovereign integrations use polling exclusively for data collection from the supported services.

* **Service Availability**: Check AWS documentation for service availability in EU Sovereign regions (`eu-isob-east-1` and `eu-isob-west-1`).

* **Regional Considerations**: When specifying `aws_regions`, ensure you're using the correct EU Sovereign region identifiers (`eu-isob-east-1`, `eu-isob-west-1`, etc.).

* **Permissions**: Your AWS EU Sovereign IAM role must have appropriate permissions for each service you enable. Refer to New Relic's documentation for specific IAM policy requirements.

* **Polling Intervals**: Consider the impact of polling intervals on AWS API rate limits and costs. Lower intervals provide more frequent updates but consume more API calls.