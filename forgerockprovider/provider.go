package forgerockprovider

import (
	"context"

	"terraform-provider-forgerock/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &forgerockProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New() func() provider.Provider {
	return func() provider.Provider {
		return &forgerockProvider{}
	}
}

// hashicupsProvider is the provider implementation.
type forgerockProvider struct{}

type forgerockProviderModel struct {
	Username      types.String `tfsdk:"username"`
	Password      types.String `tfsdk:"password"`
	Forgerock_api types.String `tfsdk:"forgerock_api"`
	RealmPath     types.String `tfsdk:"realm_path"`

	MailSender *forgerockProviderModelMail `tfsdk:"mail_sender"`
}

type forgerockProviderModelMail struct {
	SendClientSecretMail types.Bool   `tfsdk:"send_client_secret_mail"`
	SMTPServer           types.String `tfsdk:"smtp_server"`
	SMTPPort             types.Int64  `tfsdk:"smtp_port"`
	SenderEmail          types.String `tfsdk:"sender_email"`
	SenderUsername       types.String `tfsdk:"sender_username"`
	SenderPassword       types.String `tfsdk:"sender_password"`
}

// Metadata returns the provider type name.
func (p *forgerockProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "forgerock"
}

// Schema defines the provider-level schema for configuration data.
func (p *forgerockProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
			},
			"password": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
			},
			"realm_path": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
			},
			"forgerock_api": schema.StringAttribute{
				Required:  true,
				Sensitive: false,
			},

			"mail_sender": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"send_client_secret_mail": schema.BoolAttribute{
						Optional: true,
					},
					"smtp_server": schema.StringAttribute{
						Optional: true,
					},
					"smtp_port": schema.Int64Attribute{
						Optional: true,
					},
					"sender_email": schema.StringAttribute{
						Optional: true,
					},
					"sender_username": schema.StringAttribute{
						Optional: true,
					},
					"sender_password": schema.StringAttribute{
						Optional: true,
					},
				},
			},
		},
	}
}

// Configure prepares a HashiCups API client for data sources and resources.
func (p *forgerockProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	tflog.Info(ctx, "Configuring Forgerock client")
	var config forgerockProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create a new forgerock client using the configuration values
	client, err := sdk.NewClient(config.Forgerock_api.ValueString(), config.RealmPath.ValueString(), config.Username.ValueString(), config.Password.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create FORGEROCK API Client",
			"An unexpected error occurred when creating the FORGEROCK API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"FORGEROCK Client Error: "+err.Error(),
		)
		return
	}

	var mailSender *MailSender

	if config.MailSender != nil && config.MailSender.SendClientSecretMail.ValueBool() {

		mailSender = &MailSender{
			SMTPServer:     config.MailSender.SMTPServer.ValueString(),
			SMTPPort:       config.MailSender.SMTPPort.ValueInt64(),
			SenderEmail:    config.MailSender.SenderEmail.ValueString(),
			SenderUsername: config.MailSender.SenderUsername.ValueString(),
			SenderPassword: config.MailSender.SenderPassword.ValueString(),
			smtpMailer:    NewDefaultSMTPMailer(),
		}

	}

	providerContext := &OAuthClientResourceContext{
		Client:     client,
		MailSender: mailSender,
	}

	resp.DataSourceData = providerContext
	resp.ResourceData = providerContext
	tflog.Info(ctx, "Configured FORGEROCK client", map[string]any{"success": true})
}

// DataSources defines the data sources implemented in the provider.
func (p *forgerockProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

// Resources defines the resources implemented in the provider.
func (p *forgerockProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewOauthClientResource,
	}
}
