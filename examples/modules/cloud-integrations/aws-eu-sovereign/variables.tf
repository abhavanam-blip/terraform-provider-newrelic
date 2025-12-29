# Required Variables
variable "account_name" {
  description = "The name of the AWS EU Sovereign account in New Relic"
  type        = string
}

variable "aws_role_arn" {
  description = "The ARN of the IAM role in AWS EU Sovereign for New Relic integrations"
  type        = string
}

# Optional Variables
variable "metric_collection_mode" {
  description = "How metrics are collected. Either PULL or PUSH"
  type        = string
  default     = "PUSH"
  validation {
    condition     = contains(["PULL", "PUSH"], var.metric_collection_mode)
    error_message = "metric_collection_mode must be either 'PULL' or 'PUSH'."
  }
}

variable "enable_integrations" {
  description = "Whether to enable AWS integrations"
  type        = bool
  default     = true
}

# Integration Configuration Variables
variable "cloudtrail_integration" {
  description = "Configuration for CloudTrail integration"
  type = object({
    metrics_polling_interval = optional(number, 300)
    aws_regions              = optional(list(string), ["eusc-de-east-1"])
    fetch_extended_inventory = optional(bool, true)
    fetch_tags               = optional(bool, true)
    tag_key                  = optional(string, null)
    tag_value                = optional(string, null)
  })
  default = null
}

variable "health_integration" {
  description = "Configuration for AWS Health integration"
  type = object({
    metrics_polling_interval = optional(number, 300)
    fetch_extended_inventory = optional(bool, true)
    fetch_tags               = optional(bool, true)
    tag_key                  = optional(string, null)
    tag_value                = optional(string, null)
  })
  default = null
}

variable "trusted_advisor_integration" {
  description = "Configuration for AWS Trusted Advisor integration"
  type = object({
    metrics_polling_interval = optional(number, 300)
    fetch_extended_inventory = optional(bool, true)
    fetch_tags               = optional(bool, true)
    tag_key                  = optional(string, null)
    tag_value                = optional(string, null)
  })
  default = null
}

variable "xray_integration" {
  description = "Configuration for AWS X-Ray integration"
  type = object({
    metrics_polling_interval = optional(number, 300)
    aws_regions              = optional(list(string), ["eusc-de-east-1"])
    fetch_extended_inventory = optional(bool, true)
    fetch_tags               = optional(bool, true)
    tag_key                  = optional(string, null)
    tag_value                = optional(string, null)
  })
  default = null
}