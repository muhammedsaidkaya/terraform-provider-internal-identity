terraform {
  required_providers {
    internalidentity = {
      source = "registry.terraform.io/muhammedkaya/internalidentity"
    }
  }
}

provider "internalidentity" {
  api_key  = var.internalidentity_api_key
  base_url = var.internalidentity_base_url
}

data "internalidentity_users" "example" {
}

output "users" {
  value = data.internalidentity_users.example
}

variable "internalidentity_api_key" {}
variable "internalidentity_base_url" {}