package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tftypes "github.com/hashicorp/terraform-plugin-framework/types"
	"log"
)

// userModel maps postgres schema data.
type usersModel struct {
	Email tftypes.String `tfsdk:"email"`
	Users []userModel    `tfsdk:"users"`
}

// NewUsersDataSource is a helper function to simplify the provider implementation.
func NewUsersDataSource() datasource.DataSource {
	return &usersDataSource{}
}

// userDataSource is the data source implementation.
type usersDataSource struct {
	client *Client
}

func (d *usersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_users"
}

func (d *usersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"email": schema.StringAttribute{
				Optional: true,
			},
			"users": &schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"email": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *usersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state usersModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	endpoint := "/users"
	if state.Email.ValueString() != "" {
		endpoint = endpoint + "?email=" + state.Email.ValueString()
	}
	data, err := d.client.GetData(endpoint)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read User data",
			err.Error(),
		)
		return
	}

	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Fatalf("Error unmarshalling data: %v", err)
	}

	for _, user := range users {
		userState := userModel{
			ID:    types.StringValue(user.UserId),
			Name:  types.StringValue(user.Name),
			Email: types.StringValue(user.Email),
		}
		state.Users = append(state.Users, userState)
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *usersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *internalIdentity.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}
	d.client = client
}
