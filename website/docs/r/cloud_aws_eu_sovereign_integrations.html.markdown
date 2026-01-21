---
layout: "newrelic"
page_title: "New Relic: newrelic_cloud_aws_eu_sovereign_integrations"
sidebar_current: "docs-newrelic-resource-cloud-aws-eu-sovereign-integrations"
description: |-
    Integrate AWS EU Sovereign services with New Relic.
---

# Resource: newrelic\_cloud\_aws\_eu\_sovereign\_integrations

Use this resource to integrate AWS EU Sovereign services with New Relic.

## Prerequisite

Setup is required for this resource to work properly. This resource assumes you have [linked an AWS EU Sovereign account](cloud_aws_eu_sovereign_link_account.html) to New Relic and configured it to push metrics using CloudWatch Metric Streams.

New Relic doesn't automatically receive metrics from AWS EU Sovereign for some services so this resource can be used to configure integrations to those services.

## Example Usage

The following example demonstrates the use of the `newrelic_cloud_aws_eu_sovereign_integrations` resource with all AWS EU Sovereign integrations supported by the resource.

```hcl
resource "newrelic_cloud_aws_eu_sovereign_link_account" "foo" {
  arn                    = "arn:aws-eusc:iam::123456789012:role/NewRelicInfrastructure-Integrations"
  metric_collection_mode = "PULL"
  name                   = "my-eu-sovereign-account"
}

resource "newrelic_cloud_aws_eu_sovereign_integrations" "bar" {
  linked_account_id = newrelic_cloud_aws_eu_sovereign_link_account.foo.id

  cloudtrail {
    metrics_polling_interval = 300
    aws_regions              = ["eusc-de-east-1"]
  }

  health {
    metrics_polling_interval = 300
  }

  trusted_advisor {
    metrics_polling_interval = 300
  }

  x_ray {
    metrics_polling_interval = 300
    aws_regions              = ["eusc-de-east-1"]
  }
}
```

## Supported AWS EU Sovereign Integrations

-> **NOTE:** AWS EU Sovereign Cloud only supports the following four integrations.

| Block             | Description                   |
|-------------------|-------------------------------|
| `cloudtrail`      | CloudTrail Integration        |
| `health`          | Health Integration            |
| `trusted_advisor` | Trusted Advisor Integration   |
| `x_ray`           | X-Ray Integration             |

## Argument Reference

-> **WARNING:** Starting with [v3.27.2](https://registry.terraform.io/providers/newrelic/newrelic/3.27.2) of the New Relic Terraform Provider, updating the `linked_account_id` of a `newrelic_cloud_aws_eu_sovereign_integrations` resource that has been applied would **force a replacement** of the resource (destruction of the resource, followed by the creation of a new resource). Please carefully review the output of `terraform plan`, which would clearly indicate a replacement of this resource, before performing a `terraform apply`.

* `account_id` - (Optional) The New Relic account ID to operate on. This allows the user to override the `account_id` attribute set on the provider. Defaults to the environment variable `NEW_RELIC_ACCOUNT_ID`.
* `linked_account_id` - (Required) The ID of the linked AWS EU Sovereign account in New Relic.

### Arguments to be Specified with Integration Blocks

The following arguments are intended to be used within ["integration blocks"](#integration-blocks) in the resource.

* `metrics_polling_interval` - (Optional) The data polling interval **in seconds**.
  * Supported by all integration blocks: `cloudtrail`, `health`, `trusted_advisor`, `x_ray`
  * Valid values: 300, 900, 1800, 3600 (seconds)

* `aws_regions` - (Optional) Specify each AWS EU Sovereign region that includes the resources that you want to monitor.
  * Supported by: `cloudtrail`, `x_ray`
  * Valid regions: `eusc-de-east-1`

## Integration Blocks

The following section lists out arguments which may be used with each AWS EU Sovereign integration supported by this resource.

<details>
  <summary>Expand this list to see all integration blocks supported by this resource.</summary>
  <details>
    <summary>cloudtrail</summary>

*  Supported Arguments: `aws_regions`, `metrics_polling_interval`
*  Valid `metrics_polling_interval` values: 300, 900, 1800, 3600 (seconds)
```hcl
  cloudtrail {
    metrics_polling_interval = 300
    aws_regions              = ["eusc-de-east-1"]
  }
```
  </details>
  <details>
    <summary>health</summary>

*  Supported Arguments: `metrics_polling_interval`
*  Valid `metrics_polling_interval` values: 300, 900, 1800, 3600 (seconds)
```hcl
  health {
    metrics_polling_interval = 300
  }
```
  </details>
  <details>
    <summary>trusted_advisor</summary>

*  Supported Arguments: `metrics_polling_interval`
*  Valid `metrics_polling_interval` values: 300, 900, 1800, 3600 (seconds)
```hcl
  trusted_advisor {
    metrics_polling_interval = 300
  }
```
  </details>
  <details>
    <summary>x_ray</summary>

*  Supported Arguments: `aws_regions`, `metrics_polling_interval`
*  Valid `metrics_polling_interval` values: 60, 300, 900, 1800, 3600 (seconds)
```hcl
  x_ray {
    metrics_polling_interval = 300
    aws_regions              = ["eusc-de-east-1"]
  }
```
  </details>
</details>

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the AWS EU Sovereign linked account.

## Additional Examples

### All Integrations with Default Polling Intervals

```hcl
resource "newrelic_cloud_aws_eu_sovereign_integrations" "example" {
  linked_account_id = newrelic_cloud_aws_eu_sovereign_link_account.foo.id

  cloudtrail {
    metrics_polling_interval = 900
    aws_regions              = ["eusc-de-east-1"]
  }

  health {
    metrics_polling_interval = 300
  }

  trusted_advisor {
    metrics_polling_interval = 3600
  }

  x_ray {
    metrics_polling_interval = 300
    aws_regions              = ["eusc-de-east-1"]
  }
}
```

### Minimal Configuration (Empty Blocks Use Defaults)

```hcl
resource "newrelic_cloud_aws_eu_sovereign_integrations" "minimal" {
  linked_account_id = newrelic_cloud_aws_eu_sovereign_link_account.foo.id

  cloudtrail {}
  health {}
  trusted_advisor {}
  x_ray {}
}
```

## Import

AWS EU Sovereign integrations can be imported using the `id`, e.g.

```bash
$ terraform import newrelic_cloud_aws_eu_sovereign_integrations.foo <id>
```