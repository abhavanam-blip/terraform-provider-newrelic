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