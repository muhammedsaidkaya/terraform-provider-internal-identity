terraform {
  required_providers {
    internalidentity = {
      source = "registry.terraform.io/picusnext/internalidentity"
    }
  }
}

provider "internalidentity" {
  api_key  = var.internalidentity_api_key
  base_url = var.internalidentity_base_url
}

data "internalidentity_user" "example" {
  id = "a6b0af4a-71d2-4fd0-9122-1dd41647c79d"
}

output "user" {
  value = data.internalidentity_user.example
}

variable "internalidentity_api_key" {}
variable "internalidentity_base_url" {}