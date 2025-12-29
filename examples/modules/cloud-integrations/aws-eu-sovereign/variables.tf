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
variable "alb_integration" {
  description = "Configuration for ALB integration"
  type = object({
    metrics_polling_interval = optional(number, 300)
    aws_regions              = optional(list(string), ["eu-isob-east-1"])
    fetch_extended_inventory = optional(bool, true)
    fetch_tags               = optional(bool, true)
    load_balancer_prefixes   = optional(list(string), null)
    tag_key                  = optional(string, null)
    tag_value                = optional(string, null)
  })
  default = null
}

variable "api_gateway_integration" {
  description = "Configuration for API Gateway integration"
  type = object({
    metrics_polling_interval = optional(number, 300)
    aws_regions              = optional(list(string), ["eu-isob-east-1"])
    stage_prefixes           = optional(list(string), null)
    tag_key                  = optional(string, null)
    tag_value                = optional(string, null)
  })
  default = null
}

variable "auto_scaling_integration" {
  description = "Configuration for Auto Scaling integration"
  type = object({
    metrics_polling_interval = optional(number, 300)
    aws_regions              = optional(list(string), ["eu-isob-east-1"])
  })
  default = null
}

variable "cloudtrail_integration" {
  description = "Configuration for CloudTrail integration"
  type = object({
    metrics_polling_interval = optional(number, 300)
    aws_regions              = optional(list(string), ["eu-isob-east-1"])
  })
  default = null
}

variable "dynamodb_integration" {
  description = "Configuration for DynamoDB integration"
  type = object({
    metrics_polling_interval = optional(number, 300)
    aws_regions              = optional(list(string), ["eu-isob-east-1"])
    fetch_extended_inventory = optional(bool, true)
    fetch_tags               = optional(bool, true)
    tag_key                  = optional(string, null)
    tag_value                = optional(string, null)
  })
  default = null
}

variable "ebs_integration" {
  description = "Configuration for EBS integration"
  type = object({
    metrics_polling_interval = optional(number, 300)
    aws_regions              = optional(list(string), ["eu-isob-east-1"])
    fetch_extended_inventory = optional(bool, true)
    tag_key                  = optional(string, null)
    tag_value                = optional(string, null)
  })
  default = null
}

variable "ec2_integration" {
  description = "Configuration for EC2 integration"
  type = object({
    metrics_polling_interval = optional(number, 300)
    aws_regions              = optional(list(string), ["eu-isob-east-1"])
    fetch_extended_inventory = optional(bool, true)
    tag_key                  = optional(string, null)
    tag_value                = optional(string, null)
  })
  default = null
}

variable "elasticsearch_integration" {
  description = "Configuration for Elasticsearch integration"
  type = object({
    metrics_polling_interval = optional(number, 300)
    aws_regions              = optional(list(string), ["eu-isob-east-1"])
    fetch_extended_inventory = optional(bool, true)
    fetch_tags               = optional(bool, true)
    tag_key                  = optional(string, null)
    tag_value                = optional(string, null)
  })
  default = null
}

variable "elb_integration" {
  description = "Configuration for ELB integration"
  type = object({
    metrics_polling_interval = optional(number, 300)
    aws_regions              = optional(list(string), ["eu-isob-east-1"])
    fetch_extended_inventory = optional(bool, true)
    fetch_tags               = optional(bool, true)
  })
  default = null
}

variable "lambda_integration" {
  description = "Configuration for Lambda integration"
  type = object({
    metrics_polling_interval = optional(number, 300)
    aws_regions              = optional(list(string), ["eu-isob-east-1"])
    fetch_extended_inventory = optional(bool, true)
    fetch_tags               = optional(bool, true)
    tag_key                  = optional(string, null)
    tag_value                = optional(string, null)
  })
  default = null
}

variable "rds_integration" {
  description = "Configuration for RDS integration"
  type = object({
    metrics_polling_interval = optional(number, 300)
    aws_regions              = optional(list(string), ["eu-isob-east-1"])
    fetch_extended_inventory = optional(bool, true)
    fetch_tags               = optional(bool, true)
    tag_key                  = optional(string, null)
    tag_value                = optional(string, null)
  })
  default = null
}

variable "s3_integration" {
  description = "Configuration for S3 integration"
  type = object({
    metrics_polling_interval = optional(number, 3600)
    fetch_extended_inventory = optional(bool, true)
    fetch_tags               = optional(bool, true)
    tag_key                  = optional(string, null)
    tag_value                = optional(string, null)
  })
  default = null
}

variable "sns_integration" {
  description = "Configuration for SNS integration"
  type = object({
    metrics_polling_interval = optional(number, 300)
    aws_regions              = optional(list(string), ["eu-isob-east-1"])
    fetch_extended_inventory = optional(bool, true)
  })
  default = null
}

variable "sqs_integration" {
  description = "Configuration for SQS integration"
  type = object({
    metrics_polling_interval = optional(number, 300)
    aws_regions              = optional(list(string), ["eu-isob-east-1"])
    fetch_extended_inventory = optional(bool, true)
    fetch_tags               = optional(bool, true)
    queue_prefixes           = optional(list(string), null)
    tag_key                  = optional(string, null)
    tag_value                = optional(string, null)
  })
  default = null
}