# Output the linked account information
output "linked_account_id" {
  description = "The ID of the linked AWS EU Sovereign account"
  value       = newrelic_cloud_aws_eu_sovereign_link_account.this.id
}

output "linked_account_name" {
  description = "The name of the linked AWS EU Sovereign account"
  value       = newrelic_cloud_aws_eu_sovereign_link_account.this.name
}

output "integrations_id" {
  description = "The ID of the AWS EU Sovereign integrations configuration"
  value       = var.enable_integrations ? newrelic_cloud_aws_eu_sovereign_integrations.this[0].id : null
}