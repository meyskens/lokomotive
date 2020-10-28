# Terraform version and plugin versions

terraform {
  required_version = ">= 0.13"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "3.12.0"
    }
    ct = {
      source  = "poseidon/ct"
      version = "0.6.1"
    }
    template = {
      source  = "hashicorp/template"
      version = "2.1.2"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.0.0"
    }
  }
}
