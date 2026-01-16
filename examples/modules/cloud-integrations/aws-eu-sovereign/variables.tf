# Required Variables
variable "newrelic_account_id" {
  description = "The New Relic account ID"
  type        = number
}

variable "name" {
  description = "The name for the AWS EU Sovereign integration resources"
  type        = string
}

# Optional Variables
variable "newrelic_region" {
  description = "The New Relic region. EU Sovereign only supports EU region."
  type        = string
  default     = "EU"

  validation {
    condition     = var.newrelic_region == "EU"
    error_message = "EU Sovereign integrations only support the EU region."
  }
}

variable "aws_region" {
  description = "The AWS EU Sovereign region"
  type        = string
  default     = "eusc-de-east-1"
}

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

