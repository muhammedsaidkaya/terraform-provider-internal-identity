terraform {
  required_providers {
    internalidentity = {
      source = "terraform.local/muhammedkaya/internal-identity"
      version = "0.0.1"
    }
  }
}

provider "internalidentity" {
  api_key  = var.internalidentity_api_key
  base_url = var.internalidentity_base_url
}

data "internalidentity_user" "example" {
  id = var.user_id
}

output "user" {
  value = data.internalidentity_user.example
}

variable "internalidentity_api_key" {}
variable "internalidentity_base_url" {}
variable "user_id" {}