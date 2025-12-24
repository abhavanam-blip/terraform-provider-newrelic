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

  alb {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1", "eu-isob-west-1"]
    fetch_extended_inventory = true
    fetch_tags               = true
  }

  api_gateway {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1"]
  }

  auto_scaling {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1", "eu-isob-west-1"]
  }

  cloudtrail {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1"]
  }

  dynamodb {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1", "eu-isob-west-1"]
    fetch_extended_inventory = true
    fetch_tags               = true
  }

  ebs {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1"]
    fetch_extended_inventory = true
  }

  ec2 {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1", "eu-isob-west-1"]
    fetch_extended_inventory = true
  }

  elasticsearch {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1"]
    fetch_extended_inventory = true
    fetch_tags               = true
  }

  elb {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1", "eu-isob-west-1"]
    fetch_extended_inventory = true
  }

  lambda {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1"]
    fetch_extended_inventory = true
    fetch_tags               = true
  }

  rds {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1", "eu-isob-west-1"]
    fetch_extended_inventory = true
    fetch_tags               = true
  }

  s3 {
    metrics_polling_interval = 3600
    fetch_extended_inventory = true
    fetch_tags               = true
  }

  sns {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1"]
    fetch_extended_inventory = true
  }

  sqs {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1", "eu-isob-west-1"]
    fetch_extended_inventory = true
    fetch_tags               = true
  }
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Optional) The account ID for the New Relic account. If omitted, this defaults to the account ID specified in the provider configuration.
* `linked_account_id` - (Required) The ID of the AWS EU Sovereign linked account.

Each of the integration blocks supports the following common arguments:

* `metrics_polling_interval` - (Optional) The data polling interval in seconds.
* `aws_regions` - (Optional) List of AWS EU Sovereign regions that include the resources you want to monitor.

Additional arguments vary by integration type:

### `alb`
* `fetch_extended_inventory` - (Optional) Determine if extra inventory data be collected or not.
* `fetch_tags` - (Optional) Specify if tags should be collected.
* `load_balancer_prefixes` - (Optional) List of Load Balancer name prefixes to monitor.
* `tag_key` - (Optional) Tag key associated with the resources that you want to monitor.
* `tag_value` - (Optional) Tag value associated with the resources that you want to monitor.

### `api_gateway`
* `stage_prefixes` - (Optional) List of API Gateway stage prefixes to monitor.
* `tag_key` - (Optional) Tag key associated with the resources that you want to monitor.
* `tag_value` - (Optional) Tag value associated with the resources that you want to monitor.

### `auto_scaling`
No additional arguments.

### `aws_direct_connect`
No additional arguments.

### `aws_states`
No additional arguments.

### `cloudtrail`
No additional arguments.

### `dynamodb`
* `fetch_extended_inventory` - (Optional) Determine if extra inventory data be collected or not.
* `fetch_tags` - (Optional) Specify if tags should be collected.
* `tag_key` - (Optional) Tag key associated with the resources that you want to monitor.
* `tag_value` - (Optional) Tag value associated with the resources that you want to monitor.

### `ebs`
* `fetch_extended_inventory` - (Optional) Determine if extra inventory data be collected or not.
* `tag_key` - (Optional) Tag key associated with the resources that you want to monitor.
* `tag_value` - (Optional) Tag value associated with the resources that you want to monitor.

### `ec2`
* `fetch_extended_inventory` - (Optional) Determine if extra inventory data be collected or not.
* `tag_key` - (Optional) Tag key associated with the resources that you want to monitor.
* `tag_value` - (Optional) Tag value associated with the resources that you want to monitor.

### `ecs`
* `fetch_tags` - (Optional) Specify if tags should be collected.
* `tag_key` - (Optional) Tag key associated with the resources that you want to monitor.
* `tag_value` - (Optional) Tag value associated with the resources that you want to monitor.

### `efs`
* `fetch_tags` - (Optional) Specify if tags should be collected.
* `tag_key` - (Optional) Tag key associated with the resources that you want to monitor.
* `tag_value` - (Optional) Tag value associated with the resources that you want to monitor.

### `elasticache`
* `fetch_tags` - (Optional) Specify if tags should be collected.
* `tag_key` - (Optional) Tag key associated with the resources that you want to monitor.
* `tag_value` - (Optional) Tag value associated with the resources that you want to monitor.

### `elasticsearch`
* `fetch_extended_inventory` - (Optional) Determine if extra inventory data be collected or not.
* `fetch_tags` - (Optional) Specify if tags should be collected.
* `tag_key` - (Optional) Tag key associated with the resources that you want to monitor.
* `tag_value` - (Optional) Tag value associated with the resources that you want to monitor.

### `elb`
* `fetch_extended_inventory` - (Optional) Determine if extra inventory data be collected or not.
* `fetch_tags` - (Optional) Specify if tags should be collected.

### `emr`
* `fetch_tags` - (Optional) Specify if tags should be collected.
* `tag_key` - (Optional) Tag key associated with the resources that you want to monitor.
* `tag_value` - (Optional) Tag value associated with the resources that you want to monitor.

### `iam`
* `tag_key` - (Optional) Tag key associated with the resources that you want to monitor.
* `tag_value` - (Optional) Tag value associated with the resources that you want to monitor.

### `lambda`
* `fetch_extended_inventory` - (Optional) Determine if extra inventory data be collected or not.
* `fetch_tags` - (Optional) Specify if tags should be collected.
* `tag_key` - (Optional) Tag key associated with the resources that you want to monitor.
* `tag_value` - (Optional) Tag value associated with the resources that you want to monitor.

### `rds`
* `fetch_extended_inventory` - (Optional) Determine if extra inventory data be collected or not.
* `fetch_tags` - (Optional) Specify if tags should be collected.
* `tag_key` - (Optional) Tag key associated with the resources that you want to monitor.
* `tag_value` - (Optional) Tag value associated with the resources that you want to monitor.

### `redshift`
* `tag_key` - (Optional) Tag key associated with the resources that you want to monitor.
* `tag_value` - (Optional) Tag value associated with the resources that you want to monitor.

### `route53`
* `fetch_extended_inventory` - (Optional) Determine if extra inventory data be collected or not.

### `s3`
* `fetch_extended_inventory` - (Optional) Determine if extra inventory data be collected or not.
* `fetch_tags` - (Optional) Specify if tags should be collected.
* `tag_key` - (Optional) Tag key associated with the resources that you want to monitor.
* `tag_value` - (Optional) Tag value associated with the resources that you want to monitor.

### `sns`
* `fetch_extended_inventory` - (Optional) Determine if extra inventory data be collected or not.

### `sqs`
* `fetch_extended_inventory` - (Optional) Determine if extra inventory data be collected or not.
* `fetch_tags` - (Optional) Specify if tags should be collected.
* `queue_prefixes` - (Optional) List of SQS queue name prefixes to monitor.
* `tag_key` - (Optional) Tag key associated with the resources that you want to monitor.
* `tag_value` - (Optional) Tag value associated with the resources that you want to monitor.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the integration configuration.

## Import

Integrations can be imported using the `linked_account_id`, e.g.

```bash
$ terraform import newrelic_cloud_aws_eu_sovereign_integrations.foo <linked_account_id>
```

## Notes

* **Service Availability**: Not all AWS services are available in AWS EU Sovereign regions. Check AWS documentation for service availability in `eu-isob-east-1` and `eu-isob-west-1` regions.

* **Regional Considerations**: When specifying `aws_regions`, ensure you're using the correct EU Sovereign region identifiers (`eu-isob-east-1`, `eu-isob-west-1`, etc.).

* **Permissions**: Your AWS EU Sovereign IAM role must have appropriate permissions for each service you enable. Refer to New Relic's documentation for specific IAM policy requirements.

* **Polling Intervals**: Consider the impact of polling intervals on AWS API rate limits and costs. Lower intervals provide more frequent updates but consume more API calls.

* **Tag Filtering**: Use `tag_key` and `tag_value` to limit monitoring to specific resources, which can help reduce data collection costs and improve performance.