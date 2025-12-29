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

data "aws_iam_policy_document" "newrelic_assume_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type = "AWS"
      // This is the unique identifier for New Relic account on AWS EU Sovereign, there is no need to change this
      identifiers = [049749093707]
    }

    condition {
      test     = "StringEquals"
      variable = "sts:ExternalId"
      values   = [var.newrelic_account_id]
    }
  }
}

resource "aws_iam_role" "newrelic_aws_role" {
  name               = "NewRelicInfrastructure-Integrations-${var.account_name}"
  description        = "New Relic Cloud integration role"
  assume_role_policy = data.aws_iam_policy_document.newrelic_assume_policy.json
}

resource "aws_iam_policy" "newrelic_aws_permissions" {
  name        = "NewRelicCloudStreamReadPermissions-${var.account_name}"
  description = ""
  policy      = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "budgets:ViewBudget",
        "cloudtrail:LookupEvents",
        "config:BatchGetResourceConfig",
        "config:ListDiscoveredResources",
        "ec2:DescribeInternetGateways",
        "ec2:DescribeVpcs",
        "ec2:DescribeNatGateways",
        "ec2:DescribeVpcEndpoints",
        "ec2:DescribeSubnets",
        "ec2:DescribeNetworkAcls",
        "ec2:DescribeVpcAttribute",
        "ec2:DescribeRouteTables",
        "ec2:DescribeSecurityGroups",
        "ec2:DescribeVpcPeeringConnections",
        "ec2:DescribeNetworkInterfaces",
        "ec2:DescribeVpnConnections",
        "health:DescribeAffectedEntities",
        "health:DescribeEventDetails",
        "health:DescribeEvents",
        "tag:GetResources",
        "xray:BatchGet*",
        "xray:Get*"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "newrelic_aws_policy_attach" {
  role       = aws_iam_role.newrelic_aws_role.name
  policy_arn = aws_iam_policy.newrelic_aws_permissions.arn
}

resource "aws_iam_role_policy_attachment" "readonly_access_policy_attach" {
  role       = aws_iam_role.newrelic_aws_role.name
  policy_arn = "arn:aws-eusc:iam::aws:policy/ReadOnlyAccess"  # Updated for EU Sovereign
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

  dynamic "cloudtrail" {
    for_each = var.cloudtrail_integration != null ? [var.cloudtrail_integration] : []
    content {
      metrics_polling_interval = cloudtrail.value.metrics_polling_interval
      aws_regions              = cloudtrail.value.aws_regions
      fetch_extended_inventory = cloudtrail.value.fetch_extended_inventory
      fetch_tags               = cloudtrail.value.fetch_tags
      tag_key                  = cloudtrail.value.tag_key
      tag_value                = cloudtrail.value.tag_value
    }
  }

  dynamic "health" {
    for_each = var.health_integration != null ? [var.health_integration] : []
    content {
      metrics_polling_interval = health.value.metrics_polling_interval
      fetch_extended_inventory = health.value.fetch_extended_inventory
      fetch_tags               = health.value.fetch_tags
      tag_key                  = health.value.tag_key
      tag_value                = health.value.tag_value
    }
  }

  dynamic "trusted_advisor" {
    for_each = var.trusted_advisor_integration != null ? [var.trusted_advisor_integration] : []
    content {
      metrics_polling_interval = trusted_advisor.value.metrics_polling_interval
      fetch_extended_inventory = trusted_advisor.value.fetch_extended_inventory
      fetch_tags               = trusted_advisor.value.fetch_tags
      tag_key                  = trusted_advisor.value.tag_key
      tag_value                = trusted_advisor.value.tag_value
    }
  }

  dynamic "xray" {
    for_each = var.xray_integration != null ? [var.xray_integration] : []
    content {
      metrics_polling_interval = xray.value.metrics_polling_interval
      aws_regions              = xray.value.aws_regions
      fetch_extended_inventory = xray.value.fetch_extended_inventory
      fetch_tags               = xray.value.fetch_tags
      tag_key                  = xray.value.tag_key
      tag_value                = xray.value.tag_value
    }
  }
}

# PUSH mode integration for metric streaming
resource "newrelic_cloud_aws_eu_sovereign_link_account" "newrelic_cloud_integration_push" {
  name                   = "${var.account_name} metric stream"
  arn                    = aws_iam_role.newrelic_aws_role.arn
  metric_collection_mode = "PUSH"
  depends_on             = [aws_iam_role_policy_attachment.newrelic_aws_policy_attach, aws_iam_role_policy_attachment.readonly_access_policy_attach]
}

resource "newrelic_api_access_key" "newrelic_aws_access_key" {
  account_id  = var.newrelic_account_id
  key_type    = "INGEST"
  ingest_type = "LICENSE"
  name        = "Metric Stream Key for ${var.account_name}"
  notes       = "AWS Cloud Integrations Metric Stream Key"
}

resource "aws_iam_role" "firehose_newrelic_role" {
  name = "firehose_newrelic_role_${var.account_name}"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "firehose.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "readonly_access_policy_attach_2" {
  role       = aws_iam_role.firehose_newrelic_role.name
  policy_arn = "arn:aws-eusc:iam::aws:policy/ReadOnlyAccess"
}

resource "random_string" "s3-bucket-name" {
  length  = 8
  special = false
  upper   = false
}

resource "aws_s3_bucket" "newrelic_aws_bucket" {
  bucket        = "newrelic-aws-bucket-${random_string.s3-bucket-name.id}"
  force_destroy = true
}

resource "aws_s3_bucket_ownership_controls" "newrelic_ownership_controls" {
  bucket = aws_s3_bucket.newrelic_aws_bucket.id
  rule {
    object_ownership = "BucketOwnerEnforced"
  }
}

resource "aws_kinesis_firehose_delivery_stream" "newrelic_firehose_stream" {
  name        = "newrelic_firehose_stream_${var.account_name}"
  destination = "http_endpoint"
  http_endpoint_configuration {
    url                = "https://eu-aws-api.newrelic.com/cloudwatch-metrics/v1"  # Updated for EU Sovereign endpoint
    name               = "New Relic ${var.account_name}"
    access_key         = newrelic_api_access_key.newrelic_aws_access_key.key
    buffering_size     = 1
    buffering_interval = 60
    role_arn           = aws_iam_role.firehose_newrelic_role.arn
    s3_backup_mode     = "FailedDataOnly"
    s3_configuration {
      role_arn           = aws_iam_role.firehose_newrelic_role.arn
      bucket_arn         = aws_s3_bucket.newrelic_aws_bucket.arn
      buffering_size     = 10
      buffering_interval = 400
      compression_format = "GZIP"
    }
    request_configuration {
      content_encoding = "GZIP"
    }
  }
}

resource "aws_iam_role" "metric_stream_to_firehose" {
  name = "newrelic_metric_stream_to_firehose_role_${var.account_name}"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "streams.metrics.cloudwatch.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "metric_stream_to_firehose" {
  name = "default"
  role = aws_iam_role.metric_stream_to_firehose.id

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "firehose:PutRecord",
                "firehose:PutRecordBatch"
            ],
            "Resource": "${aws_kinesis_firehose_delivery_stream.newrelic_firehose_stream.arn}"
        }
    ]
}
EOF
}

resource "aws_cloudwatch_metric_stream" "newrelic_metric_stream" {
  name          = "newrelic-metric-stream-${var.account_name}"
  role_arn      = aws_iam_role.metric_stream_to_firehose.arn
  firehose_arn  = aws_kinesis_firehose_delivery_stream.newrelic_firehose_stream.arn
  output_format = "opentelemetry0.7"

  dynamic "exclude_filter" {
    for_each = var.exclude_metric_filters
    content {
      namespace    = exclude_filter.key
      metric_names = exclude_filter.value
    }
  }

  dynamic "include_filter" {
    for_each = var.include_metric_filters
    content {
      namespace    = include_filter.key
      metric_names = include_filter.value
    }
  }
}

# PULL mode integration with all AWS services
resource "newrelic_cloud_aws_eu_sovereign_link_account" "newrelic_cloud_integration_pull" {
  name                   = "${var.account_name} pull"
  arn                    = aws_iam_role.newrelic_aws_role.arn
  metric_collection_mode = "PULL"
  depends_on             = [aws_iam_role_policy_attachment.newrelic_aws_policy_attach, aws_iam_role_policy_attachment.readonly_access_policy_attach]
}

resource "newrelic_cloud_aws_eu_sovereign_integrations" "newrelic_cloud_integration_pull" {
  linked_account_id = newrelic_cloud_aws_eu_sovereign_link_account.newrelic_cloud_integration_pull.id

  # All available AWS services for EU Sovereign
  cloudtrail {}
  s3 {}
  sqs {}
  ebs {}
  alb {}
  api_gateway {}
  auto_scaling {}
  aws_direct_connect {}
  aws_states {}
  dynamo_db {}
  ec2 {}
  elastic_search {}
  elb {}
  emr {}
  iam {}
  lambda {}
  rds {}
  red_shift {}
  route53 {}
  sns {}
}

# AWS Config setup for comprehensive monitoring
resource "aws_s3_bucket" "newrelic_configuration_recorder_s3" {
  bucket        = "newrelic-configuration-recorder-${random_string.s3-bucket-name.id}"
  force_destroy = true
}

resource "aws_iam_role" "newrelic_configuration_recorder" {
  name               = "newrelic_configuration_recorder-${var.account_name}"
  assume_role_policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": "sts:AssumeRole",
            "Principal": {
                "Service": "config.amazonaws.com"
            },
            "Effect": "Allow",
            "Sid": ""
        }
      ]
    }
EOF
}

resource "aws_iam_role_policy" "newrelic_configuration_recorder_s3" {
  name = "newrelic-configuration-recorder-s3-${var.account_name}"
  role = aws_iam_role.newrelic_configuration_recorder.id

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "s3:*"
      ],
      "Effect": "Allow",
      "Resource": [
        "${aws_s3_bucket.newrelic_configuration_recorder_s3.arn}",
        "${aws_s3_bucket.newrelic_configuration_recorder_s3.arn}/*"
      ]
    }
  ]
}
POLICY
}

resource "aws_iam_role_policy_attachment" "newrelic_configuration_recorder" {
  role       = aws_iam_role.newrelic_configuration_recorder.name
  policy_arn = "arn:aws-eusc:iam::aws:policy/service-role/AWS_ConfigRole"  # Updated for EU Sovereign
}

resource "aws_config_configuration_recorder" "newrelic_recorder" {
  name     = "newrelic_configuration_recorder-${var.account_name}"
  role_arn = aws_iam_role.newrelic_configuration_recorder.arn
}

resource "aws_config_configuration_recorder_status" "newrelic_recorder_status" {
  name       = aws_config_configuration_recorder.newrelic_recorder.name
  is_enabled = true
  depends_on = [aws_config_delivery_channel.newrelic_recorder_delivery]
}

resource "aws_config_delivery_channel" "newrelic_recorder_delivery" {
  name           = "newrelic_configuration_recorder-${var.account_name}"
  s3_bucket_name = aws_s3_bucket.newrelic_configuration_recorder_s3.bucket
  depends_on = [
    aws_config_configuration_recorder.newrelic_recorder
  ]
}