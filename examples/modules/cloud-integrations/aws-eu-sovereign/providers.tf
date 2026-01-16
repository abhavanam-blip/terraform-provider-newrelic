terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 6.0"
    }
    newrelic = {
      source = "newrelic/newrelic"
    }
  }
}

provider "newrelic" {
  account_id = var.newrelic_account_id
  api_key    = var.newrelic_api_key
  region     = var.newrelic_account_region
}

provider "aws" {
  region                      = "eusc-de-east-1"
  skip_region_validation      = true   # eusc-de-east-1 is not a standard AWS region
  skip_credentials_validation = true   # STS endpoint not available at standard URL
  skip_requesting_account_id  = true   # S3 Control API not supported in EU Sovereign
}

