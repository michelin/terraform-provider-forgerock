package forgerockprovider

import (
	"context"
	"fmt"
	"log"
	"strings"

	"terraform-provider-forgerock/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"time"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource = &OAuthClientResourceContext{}
)

// NewOrderResource is a helper function to simplify the provider implementation.
func NewOauthClientResource() resource.Resource {
	return &OAuthClientResourceContext{}
}

// orderResource is the resource implementation.
type OAuthClientResourceContext struct {
	Client     *sdk.Client
	MailSender *MailSender
}

// Schema defines the schema for the resource.
func (r *OAuthClientResourceContext) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = createOAuthClientSchema()
}

// Metadata returns the resource type name.
func (r *OAuthClientResourceContext) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_oauth2Client"
}

// Call each time to initialize the client.
func (r *OAuthClientResourceContext) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.Client = req.ProviderData.(*OAuthClientResourceContext).Client
	r.MailSender = req.ProviderData.(*OAuthClientResourceContext).MailSender

}

// Create creates the resource and sets the initial Terraform state.
func (r *OAuthClientResourceContext) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan OAuth2ClientTF

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.Client.ReadClient(plan.Name.ValueString())
	if err == nil {
		resp.Diagnostics.AddError(
			"Client already exists", "A Forgerock client with the name "+plan.Name.ValueString()+" already exists")
		return
	}

	coreOAuth2ClientConfig, err := terraformResourceToSdkModel(ctx, &plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to convert terraform resource to sdk model",
			err.Error(),
		)
		return
	}

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	generatedPassword := CreatePassword(16, true, true)
	coreOAuth2ClientConfig.CoreOAuth2ClientConfig.Userpassword = NewWrappedValue(&generatedPassword)

	forgerockResponse, err := r.Client.UpdateClient(plan.Name.ValueString(), *coreOAuth2ClientConfig)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create client",
			err.Error(),
		)
		return
	}

	if r.MailSender != nil {

		wellknownEndpoint := fmt.Sprintf("%s/%s%s/%s", r.Client.ApiUrl, "realms/root", r.Client.RealmPath, ".well-known/openid-configuration")
		wellknownEndpoint = strings.Replace(wellknownEndpoint, "json", "oauth2", -1)

		err = r.MailSender.SendClientInfo(ctx, *forgerockResponse, plan.Name.ValueString(), generatedPassword, plan.AdminMail.ValueString(), wellknownEndpoint)
		if err != nil {
			resp.Diagnostics.AddWarning(
				"Unable to send client creation email",
				err.Error(),
			)
		}
	}

	err = conditionalUpdateStateFromSdkModel(ctx, &plan, forgerockResponse, false, &resp.State)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to update plan from sdk model",
			err.Error(),
		)
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *OAuthClientResourceContext) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state OAuth2ClientTF

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	oAuth2ClientConfig, err := r.Client.ReadClient(state.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to get client information: "+state.Name.ValueString(),
			err.Error(),
		)
		return
	}

	err = conditionalUpdateStateFromSdkModel(ctx, &state, oAuth2ClientConfig, false, &resp.State)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to update plan from sdk model",
			err.Error(),
		)
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *OAuthClientResourceContext) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan OAuth2ClientTF
	var lastState OAuth2ClientTF

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = req.State.Get(ctx, &lastState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	oAuth2Client, err := terraformResourceToSdkModel(ctx, &plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to convert terraform resource to sdk model",
			err.Error(),
		)
		return
	}

	lastPasswordValue := lastState.UserPasswordVersion.ValueInt64Pointer()
	nextPasswordValue := plan.UserPasswordVersion.ValueInt64Pointer()

	sendMail := false
	generatedPassword := ""
	if nextPasswordValue != nil && (lastPasswordValue == nil || *nextPasswordValue != *lastPasswordValue) {
		tflog.Info(ctx, "Client secret has changed, generating a new one")
		log.Println("[INFO] Client secret has changed, generating a new one")
		generatedPassword = CreatePassword(16, true, true)
		oAuth2Client.CoreOAuth2ClientConfig.Userpassword = NewWrappedValue(&generatedPassword)
		sendMail = true
	}

	forgerockResponse, err := r.Client.UpdateClient(plan.Name.ValueString(), *oAuth2Client)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to update client",
			err.Error(),
		)
		return
	}

	if sendMail && r.MailSender != nil {

		wellknownEndpoint := fmt.Sprintf("%s/%s%s/%s", r.Client.ApiUrl, "realms/root", r.Client.RealmPath, ".well-known/openid-configuration")
		wellknownEndpoint = strings.Replace(wellknownEndpoint, "json", "oauth2", -1)

		err = r.MailSender.SendClientInfo(ctx, *forgerockResponse, plan.Name.ValueString(), generatedPassword, plan.AdminMail.ValueString(), wellknownEndpoint)
		if err != nil {
			resp.Diagnostics.AddWarning(
				"Unable to send client update email",
				err.Error(),
			)
		}
	}

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *OAuthClientResourceContext) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OAuth2ClientTF

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.Client.DeleteClient(state.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to delete client", err.Error())
		return
	}
}

func (r *OAuthClientResourceContext) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	clientId := req.ID

	oAuth2ClientConfig, err := r.Client.ReadClient(clientId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to import client information: "+clientId,
			err.Error(),
		)
		return
	}

	oAuth2ClientConfigTemplate, err := r.Client.GetClientTemplate()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to get client template",
			err.Error(),
		)
		return
	}

	modifiedClientFields, err := sdk.ExtractChangedData(oAuth2ClientConfigTemplate, oAuth2ClientConfig)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to extract changed data",
			err.Error(),
		)
		return
	}

	var newClientResource OAuth2ClientTF
	newClientResource.Name = types.StringValue(clientId)
	err = conditionalUpdateStateFromSdkModel(ctx, &newClientResource, modifiedClientFields, true, &resp.State)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to update plan from sdk model",
			err.Error(),
		)
		return
	}

}
