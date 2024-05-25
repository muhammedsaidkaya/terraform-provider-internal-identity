package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"os"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &internalIdentityProvider{}
)

type internalIdentityProviderModel struct {
	ApiKey  types.String `tfsdk:"api_key"`
	BaseUrl types.String `tfsdk:"base_url"`
}

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &internalIdentityProvider{
			version: version,
		}
	}
}

// internalIdentityProvider is the provider implementation.
type internalIdentityProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// Metadata returns the provider type name.
func (p *internalIdentityProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "internalidentity"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *internalIdentityProvider) Schema(ctx context.Context, request provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_key": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
			},
			"base_url": schema.StringAttribute{
				Required:  true,
				Sensitive: false,
			},
		},
	}
}

func (p *internalIdentityProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config internalIdentityProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	if config.ApiKey.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Unknown Api Key value",
			"The provider cannot create the Internal Identity client as there is an unknown configuration value for the Api Key.",
		)
	}
	if config.BaseUrl.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("base_url"),
			"Unknown Base URL value",
			"The provider cannot create the Internal Identity client as there is an unknown configuration value for the Base URL.",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	apiKey := os.Getenv("INTERNAL_IDENTITY_API_KEY")
	baseUrl := os.Getenv("INTERNAL_IDENTITY_BASE_URL")

	if !config.ApiKey.IsNull() {
		apiKey = config.ApiKey.ValueString()
	}

	if apiKey == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Missing Api Key",
			"The provider cannot create the Internal Identity client as there is a missing or empty value for the Api Key. ",
		)
	}

	if !config.BaseUrl.IsNull() {
		baseUrl = config.BaseUrl.ValueString()
	}

	if baseUrl == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("base_url"),
			"Missing Base URL",
			"The provider cannot create the Internal Identity client as there is a missing or empty value for the Base URL. ",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	//Create HTTP Client
	client := NewClient(baseUrl, apiKey)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *internalIdentityProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewUserDataSource,
		NewUsersDataSource,
	}
}

func (p *internalIdentityProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}
