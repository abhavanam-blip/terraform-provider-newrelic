# AWS EU Sovereign Cloud Integration Module
# This module creates a New Relic AWS EU Sovereign cloud integration
# EU Sovereign only supports: cloudtrail, health, trusted_advisor, x_ray

# IAM policy document for New Relic to assume the role
data "aws_iam_policy_document" "newrelic_assume_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type = "AWS"
      # This is the unique identifier for New Relic account on AWS EU Sovereign Cloud
      identifiers = ["049749093707"]
    }

    condition {
      test     = "StringEquals"
      variable = "sts:ExternalId"
      values   = [var.newrelic_account_id]
    }
  }
}

# Create IAM role for New Relic
resource "aws_iam_role" "newrelic_aws_role" {
  name               = "NewRelicInfrastructure-Integrations-${var.name}"
  description        = "New Relic Cloud integration role for EU Sovereign"
  assume_role_policy = data.aws_iam_policy_document.newrelic_assume_policy.json
}

# IAM policy with permissions for New Relic integrations
resource "aws_iam_policy" "newrelic_aws_permissions" {
  name        = "NewRelicCloudStreamReadPermissions-${var.name}"
  description = "Permissions for New Relic AWS EU Sovereign Cloud integrations"
  policy      = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "cloudtrail:LookupEvents",
        "health:DescribeAffectedEntities",
        "health:DescribeEventDetails",
        "health:DescribeEvents",
        "support:DescribeTrustedAdvisorCheckRefreshStatuses",
        "support:DescribeTrustedAdvisorCheckResult",
        "support:DescribeTrustedAdvisorCheckSummaries",
        "support:DescribeTrustedAdvisorChecks",
        "support:RefreshTrustedAdvisorCheck",
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

# Attach the policy to the role
resource "aws_iam_role_policy_attachment" "newrelic_aws_policy_attach" {
  role       = aws_iam_role.newrelic_aws_role.name
  policy_arn = aws_iam_policy.newrelic_aws_permissions.arn
}

# Attach ReadOnlyAccess policy for comprehensive monitoring
resource "aws_iam_role_policy_attachment" "readonly_access_policy_attach" {
  role       = aws_iam_role.newrelic_aws_role.name
  policy_arn = "arn:aws-eusc:iam::aws:policy/ReadOnlyAccess"
}

# Link the AWS EU Sovereign account to New Relic
resource "newrelic_cloud_aws_eu_sovereign_link_account" "this" {
  account_id             = var.newrelic_account_id
  name                   = var.name
  arn                    = aws_iam_role.newrelic_aws_role.arn
  metric_collection_mode = var.metric_collection_mode
  depends_on             = [aws_iam_role_policy_attachment.newrelic_aws_policy_attach, aws_iam_role_policy_attachment.readonly_access_policy_attach]
}

# Configure AWS EU Sovereign integrations
# Only cloudtrail, health, trusted_advisor, and x_ray are supported for EU Sovereign
resource "newrelic_cloud_aws_eu_sovereign_integrations" "this" {
  count             = var.enable_integrations ? 1 : 0
  account_id        = var.newrelic_account_id
  linked_account_id = newrelic_cloud_aws_eu_sovereign_link_account.this.id

  cloudtrail {}
  health {}
  trusted_advisor {}
  x_ray {}
}
