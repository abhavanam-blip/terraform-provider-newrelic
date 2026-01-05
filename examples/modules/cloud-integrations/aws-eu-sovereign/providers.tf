terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
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
  region = "eusc-de-east-1"  # EU Sovereign region
  # Credentials can be set via AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables
}

