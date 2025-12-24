# AWS EU Sovereign Cloud Integration Module
# This module creates a New Relic AWS EU Sovereign cloud integration with commonly used services

terraform {
  required_providers {
    newrelic = {
      source  = "newrelic/newrelic"
      version = "~> 3.0"
    }
  }
}

# Link the AWS EU Sovereign account to New Relic
resource "newrelic_cloud_aws_eu_sovereign_link_account" "this" {
  name                   = var.account_name
  arn                    = var.aws_role_arn
  metric_collection_mode = var.metric_collection_mode
}

# Configure AWS EU Sovereign integrations
resource "newrelic_cloud_aws_eu_sovereign_integrations" "this" {
  count             = var.enable_integrations ? 1 : 0
  linked_account_id = newrelic_cloud_aws_eu_sovereign_link_account.this.id

  dynamic "alb" {
    for_each = var.alb_integration != null ? [var.alb_integration] : []
    content {
      metrics_polling_interval = alb.value.metrics_polling_interval
      aws_regions              = alb.value.aws_regions
      fetch_extended_inventory = alb.value.fetch_extended_inventory
      fetch_tags               = alb.value.fetch_tags
      load_balancer_prefixes   = alb.value.load_balancer_prefixes
      tag_key                  = alb.value.tag_key
      tag_value                = alb.value.tag_value
    }
  }

  dynamic "api_gateway" {
    for_each = var.api_gateway_integration != null ? [var.api_gateway_integration] : []
    content {
      metrics_polling_interval = api_gateway.value.metrics_polling_interval
      aws_regions              = api_gateway.value.aws_regions
      stage_prefixes           = api_gateway.value.stage_prefixes
      tag_key                  = api_gateway.value.tag_key
      tag_value                = api_gateway.value.tag_value
    }
  }

  dynamic "auto_scaling" {
    for_each = var.auto_scaling_integration != null ? [var.auto_scaling_integration] : []
    content {
      metrics_polling_interval = auto_scaling.value.metrics_polling_interval
      aws_regions              = auto_scaling.value.aws_regions
    }
  }

  dynamic "cloudtrail" {
    for_each = var.cloudtrail_integration != null ? [var.cloudtrail_integration] : []
    content {
      metrics_polling_interval = cloudtrail.value.metrics_polling_interval
      aws_regions              = cloudtrail.value.aws_regions
    }
  }

  dynamic "dynamodb" {
    for_each = var.dynamodb_integration != null ? [var.dynamodb_integration] : []
    content {
      metrics_polling_interval = dynamodb.value.metrics_polling_interval
      aws_regions              = dynamodb.value.aws_regions
      fetch_extended_inventory = dynamodb.value.fetch_extended_inventory
      fetch_tags               = dynamodb.value.fetch_tags
      tag_key                  = dynamodb.value.tag_key
      tag_value                = dynamodb.value.tag_value
    }
  }

  dynamic "ebs" {
    for_each = var.ebs_integration != null ? [var.ebs_integration] : []
    content {
      metrics_polling_interval = ebs.value.metrics_polling_interval
      aws_regions              = ebs.value.aws_regions
      fetch_extended_inventory = ebs.value.fetch_extended_inventory
      tag_key                  = ebs.value.tag_key
      tag_value                = ebs.value.tag_value
    }
  }

  dynamic "ec2" {
    for_each = var.ec2_integration != null ? [var.ec2_integration] : []
    content {
      metrics_polling_interval = ec2.value.metrics_polling_interval
      aws_regions              = ec2.value.aws_regions
      fetch_extended_inventory = ec2.value.fetch_extended_inventory
      tag_key                  = ec2.value.tag_key
      tag_value                = ec2.value.tag_value
    }
  }

  dynamic "elasticsearch" {
    for_each = var.elasticsearch_integration != null ? [var.elasticsearch_integration] : []
    content {
      metrics_polling_interval = elasticsearch.value.metrics_polling_interval
      aws_regions              = elasticsearch.value.aws_regions
      fetch_extended_inventory = elasticsearch.value.fetch_extended_inventory
      fetch_tags               = elasticsearch.value.fetch_tags
      tag_key                  = elasticsearch.value.tag_key
      tag_value                = elasticsearch.value.tag_value
    }
  }

  dynamic "elb" {
    for_each = var.elb_integration != null ? [var.elb_integration] : []
    content {
      metrics_polling_interval = elb.value.metrics_polling_interval
      aws_regions              = elb.value.aws_regions
      fetch_extended_inventory = elb.value.fetch_extended_inventory
      fetch_tags               = elb.value.fetch_tags
    }
  }

  dynamic "lambda" {
    for_each = var.lambda_integration != null ? [var.lambda_integration] : []
    content {
      metrics_polling_interval = lambda.value.metrics_polling_interval
      aws_regions              = lambda.value.aws_regions
      fetch_extended_inventory = lambda.value.fetch_extended_inventory
      fetch_tags               = lambda.value.fetch_tags
      tag_key                  = lambda.value.tag_key
      tag_value                = lambda.value.tag_value
    }
  }

  dynamic "rds" {
    for_each = var.rds_integration != null ? [var.rds_integration] : []
    content {
      metrics_polling_interval = rds.value.metrics_polling_interval
      aws_regions              = rds.value.aws_regions
      fetch_extended_inventory = rds.value.fetch_extended_inventory
      fetch_tags               = rds.value.fetch_tags
      tag_key                  = rds.value.tag_key
      tag_value                = rds.value.tag_value
    }
  }

  dynamic "s3" {
    for_each = var.s3_integration != null ? [var.s3_integration] : []
    content {
      metrics_polling_interval = s3.value.metrics_polling_interval
      fetch_extended_inventory = s3.value.fetch_extended_inventory
      fetch_tags               = s3.value.fetch_tags
      tag_key                  = s3.value.tag_key
      tag_value                = s3.value.tag_value
    }
  }

  dynamic "sns" {
    for_each = var.sns_integration != null ? [var.sns_integration] : []
    content {
      metrics_polling_interval = sns.value.metrics_polling_interval
      aws_regions              = sns.value.aws_regions
      fetch_extended_inventory = sns.value.fetch_extended_inventory
    }
  }

  dynamic "sqs" {
    for_each = var.sqs_integration != null ? [var.sqs_integration] : []
    content {
      metrics_polling_interval = sqs.value.metrics_polling_interval
      aws_regions              = sqs.value.aws_regions
      fetch_extended_inventory = sqs.value.fetch_extended_inventory
      fetch_tags               = sqs.value.fetch_tags
      queue_prefixes           = sqs.value.queue_prefixes
      tag_key                  = sqs.value.tag_key
      tag_value                = sqs.value.tag_value
    }
  }
}