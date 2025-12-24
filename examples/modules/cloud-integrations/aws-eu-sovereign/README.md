# AWS EU Sovereign Cloud Integration Module

This Terraform module creates and configures New Relic AWS EU Sovereign Cloud integrations. It simplifies the setup of monitoring for AWS services running in the EU Sovereign Cloud (aws-eu-iso partition).

## Prerequisites

Before using this module, ensure you have:

1. An AWS EU Sovereign account with the necessary permissions
2. A New Relic account with cloud integration capabilities
3. An IAM role in AWS EU Sovereign with permissions for New Relic integrations
4. The New Relic Terraform provider configured

## Usage

### Basic Example

```hcl
module "aws_eu_sovereign_integration" {
  source = "./modules/cloud-integrations/aws-eu-sovereign"

  account_name  = "my-eu-sovereign-account"
  aws_role_arn  = "arn:aws-eu-iso:iam::123456789012:role/NewRelicInfrastructure-Integrations"

  # Enable specific integrations
  ec2_integration = {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1", "eu-isob-west-1"]
    fetch_extended_inventory = true
  }

  lambda_integration = {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1"]
    fetch_extended_inventory = true
    fetch_tags               = true
  }

  s3_integration = {
    metrics_polling_interval = 3600
    fetch_extended_inventory = true
    fetch_tags               = true
  }
}
```

### Advanced Example with Multiple Integrations

```hcl
module "aws_eu_sovereign_integration" {
  source = "./modules/cloud-integrations/aws-eu-sovereign"

  account_name           = "production-eu-sovereign"
  aws_role_arn          = "arn:aws-eu-iso:iam::123456789012:role/NewRelicInfrastructure-Integrations"
  metric_collection_mode = "PULL"

  # ALB Integration
  alb_integration = {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1", "eu-isob-west-1"]
    fetch_extended_inventory = true
    fetch_tags               = true
    load_balancer_prefixes   = ["prod-", "staging-"]
    tag_key                  = "Environment"
    tag_value                = "Production"
  }

  # API Gateway Integration
  api_gateway_integration = {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1"]
    stage_prefixes           = ["prod", "v1"]
  }

  # Auto Scaling Integration
  auto_scaling_integration = {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1", "eu-isob-west-1"]
  }

  # DynamoDB Integration
  dynamodb_integration = {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1"]
    fetch_extended_inventory = true
    fetch_tags               = true
    tag_key                  = "Application"
    tag_value                = "MyApp"
  }

  # EC2 Integration
  ec2_integration = {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1", "eu-isob-west-1"]
    fetch_extended_inventory = true
    tag_key                  = "Environment"
    tag_value                = "Production"
  }

  # Lambda Integration
  lambda_integration = {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1"]
    fetch_extended_inventory = true
    fetch_tags               = true
    tag_key                  = "Application"
    tag_value                = "MyApp"
  }

  # RDS Integration
  rds_integration = {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1", "eu-isob-west-1"]
    fetch_extended_inventory = true
    fetch_tags               = true
    tag_key                  = "Environment"
    tag_value                = "Production"
  }

  # S3 Integration
  s3_integration = {
    metrics_polling_interval = 3600
    fetch_extended_inventory = true
    fetch_tags               = true
    tag_key                  = "Backup"
    tag_value                = "Important"
  }

  # SQS Integration
  sqs_integration = {
    metrics_polling_interval = 300
    aws_regions              = ["eu-isob-east-1"]
    fetch_extended_inventory = true
    fetch_tags               = true
    queue_prefixes           = ["prod-", "app-"]
    tag_key                  = "Application"
    tag_value                = "MyApp"
  }
}
```

## Requirements

| Name | Version |
|------|---------|
| terraform | >= 1.0 |
| newrelic | ~> 3.0 |

## Providers

| Name | Version |
|------|---------|
| newrelic | ~> 3.0 |

## Resources

| Name | Type |
|------|------|
| [newrelic_cloud_aws_eu_sovereign_link_account.this](https://registry.terraform.io/providers/newrelic/newrelic/latest/docs/resources/cloud_aws_eu_sovereign_link_account) | resource |
| [newrelic_cloud_aws_eu_sovereign_integrations.this](https://registry.terraform.io/providers/newrelic/newrelic/latest/docs/resources/cloud_aws_eu_sovereign_integrations) | resource |

## Inputs

### Required Inputs

| Name | Description | Type |
|------|-------------|------|
| account_name | The name of the AWS EU Sovereign account in New Relic | `string` |
| aws_role_arn | The ARN of the IAM role in AWS EU Sovereign for New Relic integrations | `string` |

### Optional Inputs

| Name | Description | Type | Default |
|------|-------------|------|---------|
| metric_collection_mode | How metrics are collected. Either PULL or PUSH | `string` | `"PULL"` |
| enable_integrations | Whether to enable AWS integrations | `bool` | `true` |

### Integration Configuration Inputs

Each integration can be configured with the following pattern. Set to `null` to disable an integration:

| Name | Description | Type | Default |
|------|-------------|------|---------|
| alb_integration | Configuration for ALB integration | `object` | `null` |
| api_gateway_integration | Configuration for API Gateway integration | `object` | `null` |
| auto_scaling_integration | Configuration for Auto Scaling integration | `object` | `null` |
| cloudtrail_integration | Configuration for CloudTrail integration | `object` | `null` |
| dynamodb_integration | Configuration for DynamoDB integration | `object` | `null` |
| ebs_integration | Configuration for EBS integration | `object` | `null` |
| ec2_integration | Configuration for EC2 integration | `object` | `null` |
| elasticsearch_integration | Configuration for Elasticsearch integration | `object` | `null` |
| elb_integration | Configuration for ELB integration | `object` | `null` |
| lambda_integration | Configuration for Lambda integration | `object` | `null` |
| rds_integration | Configuration for RDS integration | `object` | `null` |
| s3_integration | Configuration for S3 integration | `object` | `null` |
| sns_integration | Configuration for SNS integration | `object` | `null` |
| sqs_integration | Configuration for SQS integration | `object` | `null` |

## Outputs

| Name | Description |
|------|-------------|
| linked_account_id | The ID of the linked AWS EU Sovereign account |
| linked_account_name | The name of the linked AWS EU Sovereign account |
| integrations_id | The ID of the AWS EU Sovereign integrations configuration |

## Notes

1. **EU Sovereign Regions**: AWS EU Sovereign Cloud operates in specific regions like `eu-isob-east-1` and `eu-isob-west-1`. Ensure you're using the correct region identifiers.

2. **Service Availability**: Not all AWS services are available in EU Sovereign regions. Check AWS documentation for service availability before enabling integrations.

3. **IAM Permissions**: Your AWS EU Sovereign IAM role must have appropriate permissions for each service you want to monitor. Refer to New Relic's documentation for required permissions.

4. **Polling Intervals**: Consider the impact of polling intervals on AWS API rate limits and costs. Lower intervals provide more frequent data but consume more API calls.

5. **Tag-based Filtering**: Use `tag_key` and `tag_value` parameters to limit monitoring to specific resources, which can help reduce costs and improve performance.

6. **Metric Collection Mode**: The `metric_collection_mode` cannot be changed after the account is linked. Choose carefully between `PULL` and `PUSH` modes based on your requirements.

## Support

For issues related to this module, please refer to:
- [New Relic Terraform Provider Documentation](https://registry.terraform.io/providers/newrelic/newrelic/latest/docs)
- [New Relic AWS EU Sovereign Integration Guide](https://docs.newrelic.com/docs/infrastructure/amazon-integrations/aws-integrations/aws-eu-sovereign-cloud-integrations/)