
# Internal Identity Service Terraform Plugin/Provider 

* Only Read Operations implemented. (Data blocks)

## Local Development

### Terraform Configuration
```
user="muhammedkaya"
cat <<EOF > ~/.terraformrc
provider_installation {

  dev_overrides {
      "registry.terraform.io/$user/internalidentity" = "/Users/$user/go/bin"
  }

  direct {}
}
EOF
```

### Compile Plugin/Binary
```
user="muhammedkaya"
cd terraform-provider-internal-identity
export GOBIN="/Users/$user/go/bin"
go install .
```

### Test Terraform Module
```
cd terraform-provider-internal-identity/examples/data-sources/get_user_by_id
export TF_LOG=DEBUG
# You do not need to run terraform init
# Set Api key variable
export TF_VAR_internalidentity_api_key="" 
export TF_VAR_internalidentity_base_url=""
terraform plan 
```

### Import Provider in another Terraform Project
```
mkdir -p ~/.terraform.d/plugins/terraform.local/muhammedkaya/internal-identity/0.0.1/darwin_arm64/
cp ~/go/bin/terraform-provider-internalidentity ~/.terraform.d/plugins/terraform.local/muhammedkaya/internal-identity/0.0.1/darwin_arm64/terraform-provider-internal-identity_v0.0.1
cd terraform-provider-internal-identity/examples/data-sources/local-provider
terraform init
terraform plan
```

### Resources

* https://github.com/rapidappio/terraform-provider-rapidapp/blob/main/internal/provider/postgres_database_data_source.go
* https://spacelift.io/blog/terraform-custom-provider
* https://developer.hashicorp.com/terraform/plugin/framework/handling-data/attributes/list-nested