package forgerockprovider

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"terraform-provider-forgerock/sdk"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func terraformResourceToSdkModel(ctx context.Context, terraformOAuthClientModel *OAuth2ClientTF) (*sdk.OAuth2Client, error) {

	var advancedOAuth2ClientConfigTF AdvancedOAuth2ClientConfigTF
	var coreOAuth2ClientConfigTF CoreOAuth2ClientConfigTF
	var coreOpenIDClientConfigTF CoreOpenIDClientConfigTF
	var signEncOAuth2ClientConfigTF SignEncOAuth2ClientConfigTF
	var coreUmaClientConfigTF CoreUmaClientConfigTF
	var overrideOAuth2ClientConfigTF OverrideOAuth2ClientConfigTF

	terraformOAuthClientModel.AdvancedOAuth2ClientConfigTF.As(ctx, &advancedOAuth2ClientConfigTF, basetypes.ObjectAsOptions{})
	terraformOAuthClientModel.CoreOAuth2ClientConfigTF.As(ctx, &coreOAuth2ClientConfigTF, basetypes.ObjectAsOptions{})
	terraformOAuthClientModel.CoreOpenIDClientConfigTF.As(ctx, &coreOpenIDClientConfigTF, basetypes.ObjectAsOptions{})
	terraformOAuthClientModel.SignEncOAuth2ClientConfigTF.As(ctx, &signEncOAuth2ClientConfigTF, basetypes.ObjectAsOptions{})
	terraformOAuthClientModel.CoreUmaClientConfigTF.As(ctx, &coreUmaClientConfigTF, basetypes.ObjectAsOptions{})
	terraformOAuthClientModel.OverrideOAuth2ClientConfigTF.As(ctx, &overrideOAuth2ClientConfigTF, basetypes.ObjectAsOptions{})

	var overrideableOIDCCLaims []string
	overrideableOIDCCLaimsPtr := terraformListToStringList(overrideOAuth2ClientConfigTF.OverrideableOIDCClaims)
	if overrideableOIDCCLaimsPtr != nil {
		overrideableOIDCCLaims = *overrideableOIDCCLaimsPtr
	}

	sdkOAuthClient := &sdk.OAuth2Client{
		AdvancedOAuth2ClientConfig: sdk.AdvancedOAuth2ClientConfig{
			LogoURI:                            NewWrappedValue(terraformListToStringList(advancedOAuth2ClientConfigTF.LogoUri)),
			SubjectType:                        NewWrappedValue(advancedOAuth2ClientConfigTF.SubjectType.ValueStringPointer()),
			ClientURI:                          NewWrappedValue(terraformListToStringList(advancedOAuth2ClientConfigTF.ClientUri)),
			TokenExchangeAuthLevel:             NewWrappedValue(advancedOAuth2ClientConfigTF.TokenExchangeAuthLevel.ValueInt64Pointer()),
			ResponseTypes:                      NewWrappedValue(terraformListToStringList(advancedOAuth2ClientConfigTF.ResponseTypes)),
			MixUpMitigation:                    NewWrappedValue(advancedOAuth2ClientConfigTF.MixUpMitigation.ValueBoolPointer()),
			CustomProperties:                   NewWrappedValue(terraformListToStringList(advancedOAuth2ClientConfigTF.CustomProperties)),
			JavascriptOrigins:                  NewWrappedValue(terraformListToStringList(advancedOAuth2ClientConfigTF.JavascriptOrigins)),
			PolicyURI:                          NewWrappedValue(terraformListToStringList(advancedOAuth2ClientConfigTF.PolicyUri)),
			SoftwareVersion:                    NewWrappedValue(advancedOAuth2ClientConfigTF.SoftwareVersion.ValueStringPointer()),
			TosURI:                             NewWrappedValue(terraformListToStringList(advancedOAuth2ClientConfigTF.TosURI)),
			SectorIdentifierUri:                NewWrappedValue(advancedOAuth2ClientConfigTF.SectorIdentifierUri.ValueStringPointer()),
			TokenEndpointAuthMethod:            NewWrappedValue(advancedOAuth2ClientConfigTF.TokenEndpointAuthMethod.ValueStringPointer()),
			RefreshTokenGracePeriod:            NewWrappedValue(advancedOAuth2ClientConfigTF.RefreshTokenGracePeriod.ValueInt64Pointer()),
			IsConsentImplied:                   NewWrappedValue(advancedOAuth2ClientConfigTF.IsConsentImplied.ValueBoolPointer()),
			SoftwareIdentity:                   NewWrappedValue(advancedOAuth2ClientConfigTF.SoftwareIdentity.ValueStringPointer()),
			GrantTypes:                         NewWrappedValue(terraformListToStringList(advancedOAuth2ClientConfigTF.GrantTypes)),
			RequirePushedAuthorizationRequests: NewWrappedValue(advancedOAuth2ClientConfigTF.Require_pushed_authorization_requests.ValueBoolPointer()),
			Descriptions:                       NewWrappedValue(terraformListToStringList(advancedOAuth2ClientConfigTF.Descriptions)),
			RequestUris:                        NewWrappedValue(terraformListToStringList(advancedOAuth2ClientConfigTF.RequestUris)),
			Name:                               NewWrappedValue(terraformListToStringList(advancedOAuth2ClientConfigTF.Name)),
			Contacts:                           NewWrappedValue(terraformListToStringList(advancedOAuth2ClientConfigTF.Contacts)),
			UpdateAccessToken:                  NewWrappedValue(advancedOAuth2ClientConfigTF.UpdateAccessToken.ValueStringPointer()),
		},
		CoreOAuth2ClientConfig: sdk.CoreOAuth2ClientConfig{
			Userpassword:                 nil,
			Status:                       NewWrappedValue(coreOAuth2ClientConfigTF.Status.ValueStringPointer()),
			ClientName:                   NewWrappedValue(terraformListToStringList(coreOAuth2ClientConfigTF.ClientName)),
			ClientType:                   NewWrappedValue(coreOAuth2ClientConfigTF.ClientType.ValueStringPointer()),
			LoopbackInterfaceRedirection: NewWrappedValue(coreOAuth2ClientConfigTF.LoopbackInterfaceRedirection.ValueBoolPointer()),
			DefaultScopes:                NewWrappedValue(terraformListToStringList(coreOAuth2ClientConfigTF.DefaultScopes)),
			Agentgroup:                   NewWrappedValue(coreOAuth2ClientConfigTF.Agentgroup.ValueStringPointer()),
			RefreshTokenLifetime:         NewWrappedValue(coreOAuth2ClientConfigTF.RefreshTokenLifetime.ValueInt64Pointer()),
			Scopes:                       NewWrappedValue(terraformListToStringList(coreOAuth2ClientConfigTF.Scopes)),
			AccessTokenLifetime:          NewWrappedValue(coreOAuth2ClientConfigTF.AccessTokenLifetime.ValueInt64Pointer()),
			RedirectionUris:              NewWrappedValue(terraformListToStringList(coreOAuth2ClientConfigTF.RedirectionUris)),
			AuthorizationCodeLifetime:    NewWrappedValue(coreOAuth2ClientConfigTF.AuthorizationCodeLifetime.ValueInt64Pointer()),
		},
		CoreOpenIDClientConfig: sdk.CoreOpenIDClientConfig{
			Claims:                           NewWrappedValue(terraformListToStringList(coreOpenIDClientConfigTF.Claims)),
			BackchannelLogoutUri:             NewWrappedValue(coreOpenIDClientConfigTF.BackchannelLogoutUri.ValueStringPointer()),
			DefaultAcrValues:                 NewWrappedValue(terraformListToStringList(coreOpenIDClientConfigTF.DefaultAcrValues)),
			JwtTokenLifetime:                 NewWrappedValue(coreOpenIDClientConfigTF.JwtTokenLifetime.ValueInt64Pointer()),
			DefaultMaxAgeEnabled:             NewWrappedValue(coreOpenIDClientConfigTF.DefaultMaxAgeEnabled.ValueBoolPointer()),
			ClientSessionUri:                 NewWrappedValue(coreOpenIDClientConfigTF.ClientSessionUri.ValueStringPointer()),
			DefaultMaxAge:                    NewWrappedValue(coreOpenIDClientConfigTF.DefaultMaxAge.ValueInt64Pointer()),
			PostLogoutRedirectUri:            NewWrappedValue(terraformListToStringList(coreOpenIDClientConfigTF.PostLogoutRedirectUri)),
			BackchannelLogoutSessionRequired: NewWrappedValue(coreOpenIDClientConfigTF.Backchannel_logout_session_required.ValueBoolPointer()),
		},
		SignEncOAuth2ClientConfig: sdk.SignEncOAuth2ClientConfig{
			TokenEndpointAuthSigningAlgorithm:                      NewWrappedValue(signEncOAuth2ClientConfigTF.TokenEndpointAuthSigningAlgorithm.ValueStringPointer()),
			IDTokenEncryptionEnabled:                               NewWrappedValue(signEncOAuth2ClientConfigTF.IDTokenEncryptionEnabled.ValueBoolPointer()),
			TokenIntrospectionEncryptedResponseEncryptionAlgorithm: NewWrappedValue(signEncOAuth2ClientConfigTF.TokenIntrospectionEncryptedResponseEncryptionAlgorithm.ValueStringPointer()),
			RequestParameterSignedAlg:                              NewWrappedValue(signEncOAuth2ClientConfigTF.RequestParameterSignedAlg.ValueStringPointer()),
			AuthorizationResponseSigningAlgorithm:                  NewWrappedValue(signEncOAuth2ClientConfigTF.AuthorizationResponseSigningAlgorithm.ValueStringPointer()),
			ClientJwtPublicKey:                                     NewWrappedValue(signEncOAuth2ClientConfigTF.ClientJwtPublicKey.ValueStringPointer()),
			IDTokenPublicEncryptionKey:                             NewWrappedValue(signEncOAuth2ClientConfigTF.IDTokenPublicEncryptionKey.ValueStringPointer()),
			MTLSSubjectDN:                                          NewWrappedValue(signEncOAuth2ClientConfigTF.MTLSSubjectDN.ValueStringPointer()),
			JwkStoreCacheMissCacheTime:                             NewWrappedValue(signEncOAuth2ClientConfigTF.JwkStoreCacheMissCacheTime.ValueInt64Pointer()),
			JwkSet:                                                 NewWrappedValue(signEncOAuth2ClientConfigTF.JwkSet.ValueStringPointer()),
			IDTokenEncryptionMethod:                                NewWrappedValue(signEncOAuth2ClientConfigTF.IDTokenEncryptionMethod.ValueStringPointer()),
			JwksUri:                                                NewWrappedValue(signEncOAuth2ClientConfigTF.JwksUri.ValueStringPointer()),
			TokenIntrospectionEncryptedResponseAlg:                 NewWrappedValue(signEncOAuth2ClientConfigTF.TokenIntrospectionEncryptedResponseAlg.ValueStringPointer()),
			AuthorizationResponseEncryptionMethod:                  NewWrappedValue(signEncOAuth2ClientConfigTF.AuthorizationResponseEncryptionMethod.ValueStringPointer()),
			MTLSCertificateBoundAccessTokens:                       NewWrappedValue(signEncOAuth2ClientConfigTF.MTLSCertificateBoundAccessTokens.ValueBoolPointer()),
			UserinfoResponseFormat:                                 NewWrappedValue(signEncOAuth2ClientConfigTF.UserinfoResponseFormat.ValueStringPointer()),
			PublicKeyLocation:                                      NewWrappedValue(signEncOAuth2ClientConfigTF.PublicKeyLocation.ValueStringPointer()),
			TokenIntrospectionResponseFormat:                       NewWrappedValue(signEncOAuth2ClientConfigTF.TokenIntrospectionResponseFormat.ValueStringPointer()),
			RequestParameterEncryptedEncryptionAlgorithm:           NewWrappedValue(signEncOAuth2ClientConfigTF.RequestParameterEncryptedEncryptionAlgorithm.ValueStringPointer()),
			UserinfoSignedResponseAlg:                              NewWrappedValue(signEncOAuth2ClientConfigTF.UserinfoSignedResponseAlg.ValueStringPointer()),
			IDTokenEncryptionAlgorithm:                             NewWrappedValue(signEncOAuth2ClientConfigTF.IDTokenEncryptionAlgorithm.ValueStringPointer()),
			RequestParameterEncryptedAlg:                           NewWrappedValue(signEncOAuth2ClientConfigTF.RequestParameterEncryptedAlg.ValueStringPointer()),
			AuthorizationResponseEncryptionAlgorithm:               NewWrappedValue(signEncOAuth2ClientConfigTF.AuthorizationResponseEncryptionAlgorithm.ValueStringPointer()),
			MTLSTrustedCert:                                        NewWrappedValue(signEncOAuth2ClientConfigTF.MTLSTrustedCert.ValueStringPointer()),
			JwksCacheTimeout:                                       NewWrappedValue(signEncOAuth2ClientConfigTF.JwksCacheTimeout.ValueInt64Pointer()),
			UserinfoEncryptedResponseAlg:                           NewWrappedValue(signEncOAuth2ClientConfigTF.UserinfoEncryptedResponseAlg.ValueStringPointer()),
			IDTokenSignedResponseAlg:                               NewWrappedValue(signEncOAuth2ClientConfigTF.IDTokenSignedResponseAlg.ValueStringPointer()),
			UserinfoEncryptedResponseEncryptionAlgorithm:           NewWrappedValue(signEncOAuth2ClientConfigTF.UserinfoEncryptedResponseEncryptionAlgorithm.ValueStringPointer()),
			TokenIntrospectionSignedResponseAlg:                    NewWrappedValue(signEncOAuth2ClientConfigTF.TokenIntrospectionSignedResponseAlg.ValueStringPointer()),
		},
		CoreUmaClientConfig: sdk.CoreUmaClientConfig{
			ClaimsRedirectionUris: NewWrappedValue(terraformListToStringList(coreUmaClientConfigTF.ClaimsRedirectionUris)),
		},
		OverrideOAuth2ClientConfig: sdk.OverrideOAuth2ClientConfig{
			AccessTokenMayActScript:                 overrideOAuth2ClientConfigTF.AccessTokenMayActScript.ValueString(),
			AccessTokenModificationPluginType:       overrideOAuth2ClientConfigTF.AccessTokenModificationPluginType.ValueString(),
			AccessTokenModificationScript:           overrideOAuth2ClientConfigTF.AccessTokenModificationScript.ValueString(),
			AccessTokenModifierClass:                overrideOAuth2ClientConfigTF.AccessTokenModifierClass.ValueString(),
			AuthorizeEndpointDataProviderClass:      overrideOAuth2ClientConfigTF.AuthorizeEndpointDataProviderClass.ValueString(),
			AuthorizeEndpointDataProviderPluginType: overrideOAuth2ClientConfigTF.AuthorizeEndpointDataProviderPluginType.ValueString(),
			AuthorizeEndpointDataProviderScript:     overrideOAuth2ClientConfigTF.AuthorizeEndpointDataProviderScript.ValueString(),
			ClientsCanSkipConsent:                   overrideOAuth2ClientConfigTF.ClientsCanSkipConsent.ValueBool(),
			CustomLoginUrlTemplate:                  overrideOAuth2ClientConfigTF.CustomLoginUrlTemplate.ValueString(),
			EvaluateScopeClass:                      overrideOAuth2ClientConfigTF.EvaluateScopeClass.ValueString(),
			EvaluateScopePluginType:                 overrideOAuth2ClientConfigTF.EvaluateScopePluginType.ValueString(),
			EvaluateScopeScript:                     overrideOAuth2ClientConfigTF.EvaluateScopeScript.ValueString(),
			EnableRemoteConsent:                     overrideOAuth2ClientConfigTF.EnableRemoteConsent.ValueBool(),
			IssueRefreshToken:                       overrideOAuth2ClientConfigTF.IssueRefreshToken.ValueBool(),
			IssueRefreshTokenOnRefreshedToken:       overrideOAuth2ClientConfigTF.IssueRefreshTokenOnRefreshedToken.ValueBool(),
			OidcClaimsClass:                         overrideOAuth2ClientConfigTF.OidcClaimsClass.ValueString(),
			OidcClaimsPluginType:                    overrideOAuth2ClientConfigTF.OidcClaimsPluginType.ValueString(),
			OidcClaimsScript:                        overrideOAuth2ClientConfigTF.OidcClaimsScript.ValueString(),
			OidcMayActScript:                        overrideOAuth2ClientConfigTF.OidcMayActScript.ValueString(),
			OverrideableOIDCClaims:                  overrideableOIDCCLaims,
			ProviderOverridesEnabled:                overrideOAuth2ClientConfigTF.ProviderOverridesEnabled.ValueBool(),
			RemoteConsentServiceId:                  overrideOAuth2ClientConfigTF.RemoteConsentServiceId.ValueString(),
			ScopesPolicySet:                         overrideOAuth2ClientConfigTF.ScopesPolicySet.ValueString(),
			StatelessTokensEnabled:                  overrideOAuth2ClientConfigTF.StatelessTokensEnabled.ValueBool(),
			TokenEncryptionEnabled:                  overrideOAuth2ClientConfigTF.TokenEncryptionEnabled.ValueBool(),
			UsePolicyEngineForScope:                 overrideOAuth2ClientConfigTF.UsePolicyEngineForScope.ValueBool(),
			ValidateScopeClass:                      overrideOAuth2ClientConfigTF.ValidateScopeClass.ValueString(),
			ValidateScopePluginType:                 overrideOAuth2ClientConfigTF.ValidateScopePluginType.ValueString(),
			ValidateScopeScript:                     overrideOAuth2ClientConfigTF.ValidateScopeScript.ValueString(),
		},
	}
	return sdkOAuthClient, nil
}

// We only retain values that have been declared by the provider.
// The rest are set to default values. Therefore, it is advisable to consider updating the plan
// rather than creating a new complete object.
func conditionalUpdateStateFromSdkModel(ctx context.Context, plan *OAuth2ClientTF, sdkOAuthClient *sdk.OAuth2Client, forceReplace bool, state *tfsdk.State) (err error) {

	defer func() {
		if r := recover(); r != nil {
			tflog.Error(ctx, "UpdatePlanFromSdkModel: failed %v")
			err = fmt.Errorf("Failed to update plan from sdk model: %v", r)
		}
	}()

	var advancedOAuth2ClientConfigTF AdvancedOAuth2ClientConfigTF
	var coreOAuth2ClientConfigTF CoreOAuth2ClientConfigTF
	var coreOpenIDClientConfigTF CoreOpenIDClientConfigTF
	var signEncOAuth2ClientConfigTF SignEncOAuth2ClientConfigTF
	var coreUmaClientConfigTF CoreUmaClientConfigTF
	var overrideOAuth2ClientConfigTF OverrideOAuth2ClientConfigTF

	plan.AdvancedOAuth2ClientConfigTF.As(ctx, &advancedOAuth2ClientConfigTF, basetypes.ObjectAsOptions{})
	plan.CoreOAuth2ClientConfigTF.As(ctx, &coreOAuth2ClientConfigTF, basetypes.ObjectAsOptions{})
	plan.CoreOpenIDClientConfigTF.As(ctx, &coreOpenIDClientConfigTF, basetypes.ObjectAsOptions{})
	plan.SignEncOAuth2ClientConfigTF.As(ctx, &signEncOAuth2ClientConfigTF, basetypes.ObjectAsOptions{})
	plan.CoreUmaClientConfigTF.As(ctx, &coreUmaClientConfigTF, basetypes.ObjectAsOptions{})
	plan.OverrideOAuth2ClientConfigTF.As(ctx, &overrideOAuth2ClientConfigTF, basetypes.ObjectAsOptions{})

	advancedOAuth2ClientConfigTF.LogoUri = updateTerraformValue(ctx, advancedOAuth2ClientConfigTF.LogoUri, sdkOAuthClient.AdvancedOAuth2ClientConfig.LogoURI, forceReplace).(types.List)
	advancedOAuth2ClientConfigTF.SubjectType = updateTerraformValue(ctx, advancedOAuth2ClientConfigTF.SubjectType, sdkOAuthClient.AdvancedOAuth2ClientConfig.SubjectType, forceReplace).(types.String)
	advancedOAuth2ClientConfigTF.ClientUri = updateTerraformValue(ctx, advancedOAuth2ClientConfigTF.ClientUri, sdkOAuthClient.AdvancedOAuth2ClientConfig.ClientURI, forceReplace).(types.List)
	advancedOAuth2ClientConfigTF.TokenExchangeAuthLevel = updateTerraformValue(ctx, advancedOAuth2ClientConfigTF.TokenExchangeAuthLevel, sdkOAuthClient.AdvancedOAuth2ClientConfig.TokenExchangeAuthLevel, forceReplace).(types.Int64)
	advancedOAuth2ClientConfigTF.ResponseTypes = updateTerraformValue(ctx, advancedOAuth2ClientConfigTF.ResponseTypes, sdkOAuthClient.AdvancedOAuth2ClientConfig.ResponseTypes, forceReplace).(types.List)
	advancedOAuth2ClientConfigTF.MixUpMitigation = updateTerraformValue(ctx, advancedOAuth2ClientConfigTF.MixUpMitigation, sdkOAuthClient.AdvancedOAuth2ClientConfig.MixUpMitigation, forceReplace).(types.Bool)
	advancedOAuth2ClientConfigTF.CustomProperties = updateTerraformValue(ctx, advancedOAuth2ClientConfigTF.CustomProperties, sdkOAuthClient.AdvancedOAuth2ClientConfig.CustomProperties, forceReplace).(types.List)
	advancedOAuth2ClientConfigTF.JavascriptOrigins = updateTerraformValue(ctx, advancedOAuth2ClientConfigTF.JavascriptOrigins, sdkOAuthClient.AdvancedOAuth2ClientConfig.JavascriptOrigins, forceReplace).(types.List)
	advancedOAuth2ClientConfigTF.PolicyUri = updateTerraformValue(ctx, advancedOAuth2ClientConfigTF.PolicyUri, sdkOAuthClient.AdvancedOAuth2ClientConfig.PolicyURI, forceReplace).(types.List)
	advancedOAuth2ClientConfigTF.SoftwareVersion = updateTerraformValue(ctx, advancedOAuth2ClientConfigTF.SoftwareVersion, sdkOAuthClient.AdvancedOAuth2ClientConfig.SoftwareVersion, forceReplace).(types.String)
	advancedOAuth2ClientConfigTF.TosURI = updateTerraformValue(ctx, advancedOAuth2ClientConfigTF.TosURI, sdkOAuthClient.AdvancedOAuth2ClientConfig.TosURI, forceReplace).(types.List)
	advancedOAuth2ClientConfigTF.SectorIdentifierUri = updateTerraformValue(ctx, advancedOAuth2ClientConfigTF.SectorIdentifierUri, sdkOAuthClient.AdvancedOAuth2ClientConfig.SectorIdentifierUri, forceReplace).(types.String)
	advancedOAuth2ClientConfigTF.TokenEndpointAuthMethod = updateTerraformValue(ctx, advancedOAuth2ClientConfigTF.TokenEndpointAuthMethod, sdkOAuthClient.AdvancedOAuth2ClientConfig.TokenEndpointAuthMethod, forceReplace).(types.String)
	advancedOAuth2ClientConfigTF.RefreshTokenGracePeriod = updateTerraformValue(ctx, advancedOAuth2ClientConfigTF.RefreshTokenGracePeriod, sdkOAuthClient.AdvancedOAuth2ClientConfig.RefreshTokenGracePeriod, forceReplace).(types.Int64)
	advancedOAuth2ClientConfigTF.IsConsentImplied = updateTerraformValue(ctx, advancedOAuth2ClientConfigTF.IsConsentImplied, sdkOAuthClient.AdvancedOAuth2ClientConfig.IsConsentImplied, forceReplace).(types.Bool)
	advancedOAuth2ClientConfigTF.SoftwareIdentity = updateTerraformValue(ctx, advancedOAuth2ClientConfigTF.SoftwareIdentity, sdkOAuthClient.AdvancedOAuth2ClientConfig.SoftwareIdentity, forceReplace).(types.String)
	advancedOAuth2ClientConfigTF.GrantTypes = updateTerraformValue(ctx, advancedOAuth2ClientConfigTF.GrantTypes, sdkOAuthClient.AdvancedOAuth2ClientConfig.GrantTypes, forceReplace).(types.List)
	advancedOAuth2ClientConfigTF.Require_pushed_authorization_requests = updateTerraformValue(ctx, advancedOAuth2ClientConfigTF.Require_pushed_authorization_requests, sdkOAuthClient.AdvancedOAuth2ClientConfig.RequirePushedAuthorizationRequests, forceReplace).(types.Bool)
	advancedOAuth2ClientConfigTF.Descriptions = updateTerraformValue(ctx, advancedOAuth2ClientConfigTF.Descriptions, sdkOAuthClient.AdvancedOAuth2ClientConfig.Descriptions, forceReplace).(types.List)
	advancedOAuth2ClientConfigTF.RequestUris = updateTerraformValue(ctx, advancedOAuth2ClientConfigTF.RequestUris, sdkOAuthClient.AdvancedOAuth2ClientConfig.RequestUris, forceReplace).(types.List)
	advancedOAuth2ClientConfigTF.Name = updateTerraformValue(ctx, advancedOAuth2ClientConfigTF.Name, sdkOAuthClient.AdvancedOAuth2ClientConfig.Name, forceReplace).(types.List)
	advancedOAuth2ClientConfigTF.Contacts = updateTerraformValue(ctx, advancedOAuth2ClientConfigTF.Contacts, sdkOAuthClient.AdvancedOAuth2ClientConfig.Contacts, forceReplace).(types.List)
	advancedOAuth2ClientConfigTF.UpdateAccessToken = updateTerraformValue(ctx, advancedOAuth2ClientConfigTF.UpdateAccessToken, sdkOAuthClient.AdvancedOAuth2ClientConfig.UpdateAccessToken, forceReplace).(types.String)

	plan.UserPasswordVersion = updateTerraformValue(ctx, plan.UserPasswordVersion, nil, forceReplace).(types.Int64)

	coreOAuth2ClientConfigTF.Status = updateTerraformValue(ctx, coreOAuth2ClientConfigTF.Status, sdkOAuthClient.CoreOAuth2ClientConfig.Status, forceReplace).(types.String)
	coreOAuth2ClientConfigTF.ClientName = updateTerraformValue(ctx, coreOAuth2ClientConfigTF.ClientName, sdkOAuthClient.CoreOAuth2ClientConfig.ClientName, forceReplace).(types.List)
	coreOAuth2ClientConfigTF.ClientType = updateTerraformValue(ctx, coreOAuth2ClientConfigTF.ClientType, sdkOAuthClient.CoreOAuth2ClientConfig.ClientType, forceReplace).(types.String)
	coreOAuth2ClientConfigTF.LoopbackInterfaceRedirection = updateTerraformValue(ctx, coreOAuth2ClientConfigTF.LoopbackInterfaceRedirection, sdkOAuthClient.CoreOAuth2ClientConfig.LoopbackInterfaceRedirection, forceReplace).(types.Bool)
	coreOAuth2ClientConfigTF.DefaultScopes = updateTerraformValue(ctx, coreOAuth2ClientConfigTF.DefaultScopes, sdkOAuthClient.CoreOAuth2ClientConfig.DefaultScopes, forceReplace).(types.List)
	coreOAuth2ClientConfigTF.Agentgroup = updateTerraformValue(ctx, coreOAuth2ClientConfigTF.Agentgroup, sdkOAuthClient.CoreOAuth2ClientConfig.Agentgroup, forceReplace).(types.String)
	coreOAuth2ClientConfigTF.RefreshTokenLifetime = updateTerraformValue(ctx, coreOAuth2ClientConfigTF.RefreshTokenLifetime, sdkOAuthClient.CoreOAuth2ClientConfig.RefreshTokenLifetime, forceReplace).(types.Int64)
	coreOAuth2ClientConfigTF.Scopes = updateTerraformValue(ctx, coreOAuth2ClientConfigTF.Scopes, sdkOAuthClient.CoreOAuth2ClientConfig.Scopes, forceReplace).(types.List)
	coreOAuth2ClientConfigTF.AccessTokenLifetime = updateTerraformValue(ctx, coreOAuth2ClientConfigTF.AccessTokenLifetime, sdkOAuthClient.CoreOAuth2ClientConfig.AccessTokenLifetime, forceReplace).(types.Int64)
	coreOAuth2ClientConfigTF.RedirectionUris = updateTerraformValue(ctx, coreOAuth2ClientConfigTF.RedirectionUris, sdkOAuthClient.CoreOAuth2ClientConfig.RedirectionUris, forceReplace).(types.List)
	coreOAuth2ClientConfigTF.AuthorizationCodeLifetime = updateTerraformValue(ctx, coreOAuth2ClientConfigTF.AuthorizationCodeLifetime, sdkOAuthClient.CoreOAuth2ClientConfig.AuthorizationCodeLifetime, forceReplace).(types.Int64)

	coreOpenIDClientConfigTF.Claims = updateTerraformValue(ctx, coreOpenIDClientConfigTF.Claims, sdkOAuthClient.CoreOpenIDClientConfig.Claims, forceReplace).(types.List)
	coreOpenIDClientConfigTF.BackchannelLogoutUri = updateTerraformValue(ctx, coreOpenIDClientConfigTF.BackchannelLogoutUri, sdkOAuthClient.CoreOpenIDClientConfig.BackchannelLogoutUri, forceReplace).(types.String)
	coreOpenIDClientConfigTF.DefaultAcrValues = updateTerraformValue(ctx, coreOpenIDClientConfigTF.DefaultAcrValues, sdkOAuthClient.CoreOpenIDClientConfig.DefaultAcrValues, forceReplace).(types.List)
	coreOpenIDClientConfigTF.JwtTokenLifetime = updateTerraformValue(ctx, coreOpenIDClientConfigTF.JwtTokenLifetime, sdkOAuthClient.CoreOpenIDClientConfig.JwtTokenLifetime, forceReplace).(types.Int64)
	coreOpenIDClientConfigTF.DefaultMaxAgeEnabled = updateTerraformValue(ctx, coreOpenIDClientConfigTF.DefaultMaxAgeEnabled, sdkOAuthClient.CoreOpenIDClientConfig.DefaultMaxAgeEnabled, forceReplace).(types.Bool)
	coreOpenIDClientConfigTF.ClientSessionUri = updateTerraformValue(ctx, coreOpenIDClientConfigTF.ClientSessionUri, sdkOAuthClient.CoreOpenIDClientConfig.ClientSessionUri, forceReplace).(types.String)
	coreOpenIDClientConfigTF.DefaultMaxAge = updateTerraformValue(ctx, coreOpenIDClientConfigTF.DefaultMaxAge, sdkOAuthClient.CoreOpenIDClientConfig.DefaultMaxAge, forceReplace).(types.Int64)
	coreOpenIDClientConfigTF.PostLogoutRedirectUri = updateTerraformValue(ctx, coreOpenIDClientConfigTF.PostLogoutRedirectUri, sdkOAuthClient.CoreOpenIDClientConfig.PostLogoutRedirectUri, forceReplace).(types.List)
	coreOpenIDClientConfigTF.Backchannel_logout_session_required = updateTerraformValue(ctx, coreOpenIDClientConfigTF.Backchannel_logout_session_required, sdkOAuthClient.CoreOpenIDClientConfig.BackchannelLogoutSessionRequired, forceReplace).(types.Bool)

	signEncOAuth2ClientConfigTF.AuthorizationResponseEncryptionAlgorithm = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.AuthorizationResponseEncryptionAlgorithm, sdkOAuthClient.SignEncOAuth2ClientConfig.AuthorizationResponseEncryptionAlgorithm, forceReplace).(types.String)
	signEncOAuth2ClientConfigTF.AuthorizationResponseEncryptionMethod = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.AuthorizationResponseEncryptionMethod, sdkOAuthClient.SignEncOAuth2ClientConfig.AuthorizationResponseEncryptionMethod, forceReplace).(types.String)
	signEncOAuth2ClientConfigTF.AuthorizationResponseSigningAlgorithm = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.AuthorizationResponseSigningAlgorithm, sdkOAuthClient.SignEncOAuth2ClientConfig.AuthorizationResponseSigningAlgorithm, forceReplace).(types.String)
	signEncOAuth2ClientConfigTF.ClientJwtPublicKey = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.ClientJwtPublicKey, sdkOAuthClient.SignEncOAuth2ClientConfig.ClientJwtPublicKey, forceReplace).(types.String)
	signEncOAuth2ClientConfigTF.IDTokenEncryptionAlgorithm = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.IDTokenEncryptionAlgorithm, sdkOAuthClient.SignEncOAuth2ClientConfig.IDTokenEncryptionAlgorithm, forceReplace).(types.String)
	signEncOAuth2ClientConfigTF.IDTokenEncryptionEnabled = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.IDTokenEncryptionEnabled, sdkOAuthClient.SignEncOAuth2ClientConfig.IDTokenEncryptionEnabled, forceReplace).(types.Bool)
	signEncOAuth2ClientConfigTF.IDTokenEncryptionMethod = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.IDTokenEncryptionMethod, sdkOAuthClient.SignEncOAuth2ClientConfig.IDTokenEncryptionMethod, forceReplace).(types.String)
	signEncOAuth2ClientConfigTF.IDTokenPublicEncryptionKey = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.IDTokenPublicEncryptionKey, sdkOAuthClient.SignEncOAuth2ClientConfig.IDTokenPublicEncryptionKey, forceReplace).(types.String)
	signEncOAuth2ClientConfigTF.IDTokenSignedResponseAlg = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.IDTokenSignedResponseAlg, sdkOAuthClient.SignEncOAuth2ClientConfig.IDTokenSignedResponseAlg, forceReplace).(types.String)
	signEncOAuth2ClientConfigTF.JwkSet = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.JwkSet, sdkOAuthClient.SignEncOAuth2ClientConfig.JwkSet, forceReplace).(types.String)
	signEncOAuth2ClientConfigTF.JwkStoreCacheMissCacheTime = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.JwkStoreCacheMissCacheTime, sdkOAuthClient.SignEncOAuth2ClientConfig.JwkStoreCacheMissCacheTime, forceReplace).(types.Int64)
	signEncOAuth2ClientConfigTF.JwksCacheTimeout = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.JwksCacheTimeout, sdkOAuthClient.SignEncOAuth2ClientConfig.JwksCacheTimeout, forceReplace).(types.Int64)
	signEncOAuth2ClientConfigTF.JwksUri = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.JwksUri, sdkOAuthClient.SignEncOAuth2ClientConfig.JwksUri, forceReplace).(types.String)
	signEncOAuth2ClientConfigTF.MTLSCertificateBoundAccessTokens = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.MTLSCertificateBoundAccessTokens, sdkOAuthClient.SignEncOAuth2ClientConfig.MTLSCertificateBoundAccessTokens, forceReplace).(types.Bool)
	signEncOAuth2ClientConfigTF.MTLSSubjectDN = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.MTLSSubjectDN, sdkOAuthClient.SignEncOAuth2ClientConfig.MTLSSubjectDN, forceReplace).(types.String)
	signEncOAuth2ClientConfigTF.MTLSTrustedCert = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.MTLSTrustedCert, sdkOAuthClient.SignEncOAuth2ClientConfig.MTLSTrustedCert, forceReplace).(types.String)
	signEncOAuth2ClientConfigTF.PublicKeyLocation = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.PublicKeyLocation, sdkOAuthClient.SignEncOAuth2ClientConfig.PublicKeyLocation, forceReplace).(types.String)
	signEncOAuth2ClientConfigTF.RequestParameterEncryptedAlg = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.RequestParameterEncryptedAlg, sdkOAuthClient.SignEncOAuth2ClientConfig.RequestParameterEncryptedAlg, forceReplace).(types.String)
	signEncOAuth2ClientConfigTF.RequestParameterEncryptedEncryptionAlgorithm = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.RequestParameterEncryptedEncryptionAlgorithm, sdkOAuthClient.SignEncOAuth2ClientConfig.RequestParameterEncryptedEncryptionAlgorithm, forceReplace).(types.String)
	signEncOAuth2ClientConfigTF.RequestParameterSignedAlg = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.RequestParameterSignedAlg, sdkOAuthClient.SignEncOAuth2ClientConfig.RequestParameterSignedAlg, forceReplace).(types.String)
	signEncOAuth2ClientConfigTF.TokenEndpointAuthSigningAlgorithm = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.TokenEndpointAuthSigningAlgorithm, sdkOAuthClient.SignEncOAuth2ClientConfig.TokenEndpointAuthSigningAlgorithm, forceReplace).(types.String)
	signEncOAuth2ClientConfigTF.TokenIntrospectionEncryptedResponseAlg = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.TokenIntrospectionEncryptedResponseAlg, sdkOAuthClient.SignEncOAuth2ClientConfig.TokenIntrospectionEncryptedResponseAlg, forceReplace).(types.String)
	signEncOAuth2ClientConfigTF.TokenIntrospectionEncryptedResponseEncryptionAlgorithm = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.TokenIntrospectionEncryptedResponseEncryptionAlgorithm, sdkOAuthClient.SignEncOAuth2ClientConfig.TokenIntrospectionEncryptedResponseEncryptionAlgorithm, forceReplace).(types.String)
	signEncOAuth2ClientConfigTF.TokenIntrospectionResponseFormat = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.TokenIntrospectionResponseFormat, sdkOAuthClient.SignEncOAuth2ClientConfig.TokenIntrospectionResponseFormat, forceReplace).(types.String)
	signEncOAuth2ClientConfigTF.TokenIntrospectionSignedResponseAlg = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.TokenIntrospectionSignedResponseAlg, sdkOAuthClient.SignEncOAuth2ClientConfig.TokenIntrospectionSignedResponseAlg, forceReplace).(types.String)
	signEncOAuth2ClientConfigTF.UserinfoEncryptedResponseAlg = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.UserinfoEncryptedResponseAlg, sdkOAuthClient.SignEncOAuth2ClientConfig.UserinfoEncryptedResponseAlg, forceReplace).(types.String)
	signEncOAuth2ClientConfigTF.UserinfoEncryptedResponseEncryptionAlgorithm = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.UserinfoEncryptedResponseEncryptionAlgorithm, sdkOAuthClient.SignEncOAuth2ClientConfig.UserinfoEncryptedResponseEncryptionAlgorithm, forceReplace).(types.String)
	signEncOAuth2ClientConfigTF.UserinfoResponseFormat = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.UserinfoResponseFormat, sdkOAuthClient.SignEncOAuth2ClientConfig.UserinfoResponseFormat, forceReplace).(types.String)
	signEncOAuth2ClientConfigTF.UserinfoSignedResponseAlg = updateTerraformValue(ctx, signEncOAuth2ClientConfigTF.UserinfoSignedResponseAlg, sdkOAuthClient.SignEncOAuth2ClientConfig.UserinfoSignedResponseAlg, forceReplace).(types.String)

	coreUmaClientConfigTF.ClaimsRedirectionUris = updateTerraformValue(ctx, coreUmaClientConfigTF.ClaimsRedirectionUris, sdkOAuthClient.CoreUmaClientConfig.ClaimsRedirectionUris, forceReplace).(types.List)

	overrideOAuth2ClientConfigTF.AccessTokenMayActScript = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.AccessTokenMayActScript, sdkOAuthClient.OverrideOAuth2ClientConfig.AccessTokenMayActScript, forceReplace).(types.String)
	overrideOAuth2ClientConfigTF.AccessTokenModificationPluginType = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.AccessTokenModificationPluginType, sdkOAuthClient.OverrideOAuth2ClientConfig.AccessTokenModificationPluginType, forceReplace).(types.String)
	overrideOAuth2ClientConfigTF.AccessTokenModificationScript = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.AccessTokenModificationScript, sdkOAuthClient.OverrideOAuth2ClientConfig.AccessTokenModificationScript, forceReplace).(types.String)
	overrideOAuth2ClientConfigTF.AccessTokenModifierClass = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.AccessTokenModifierClass, sdkOAuthClient.OverrideOAuth2ClientConfig.AccessTokenModifierClass, forceReplace).(types.String)
	overrideOAuth2ClientConfigTF.AuthorizeEndpointDataProviderClass = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.AuthorizeEndpointDataProviderClass, sdkOAuthClient.OverrideOAuth2ClientConfig.AuthorizeEndpointDataProviderClass, forceReplace).(types.String)
	overrideOAuth2ClientConfigTF.AuthorizeEndpointDataProviderPluginType = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.AuthorizeEndpointDataProviderPluginType, sdkOAuthClient.OverrideOAuth2ClientConfig.AuthorizeEndpointDataProviderPluginType, forceReplace).(types.String)
	overrideOAuth2ClientConfigTF.AuthorizeEndpointDataProviderScript = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.AuthorizeEndpointDataProviderScript, sdkOAuthClient.OverrideOAuth2ClientConfig.AuthorizeEndpointDataProviderScript, forceReplace).(types.String)
	overrideOAuth2ClientConfigTF.ClientsCanSkipConsent = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.ClientsCanSkipConsent, sdkOAuthClient.OverrideOAuth2ClientConfig.ClientsCanSkipConsent, forceReplace).(types.Bool)
	overrideOAuth2ClientConfigTF.CustomLoginUrlTemplate = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.CustomLoginUrlTemplate, sdkOAuthClient.OverrideOAuth2ClientConfig.CustomLoginUrlTemplate, forceReplace).(types.String)
	overrideOAuth2ClientConfigTF.EvaluateScopeClass = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.EvaluateScopeClass, sdkOAuthClient.OverrideOAuth2ClientConfig.EvaluateScopeClass, forceReplace).(types.String)
	overrideOAuth2ClientConfigTF.EvaluateScopePluginType = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.EvaluateScopePluginType, sdkOAuthClient.OverrideOAuth2ClientConfig.EvaluateScopePluginType, forceReplace).(types.String)
	overrideOAuth2ClientConfigTF.EvaluateScopeScript = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.EvaluateScopeScript, sdkOAuthClient.OverrideOAuth2ClientConfig.EvaluateScopeScript, forceReplace).(types.String)
	overrideOAuth2ClientConfigTF.EnableRemoteConsent = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.EnableRemoteConsent, sdkOAuthClient.OverrideOAuth2ClientConfig.EnableRemoteConsent, forceReplace).(types.Bool)
	overrideOAuth2ClientConfigTF.IssueRefreshToken = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.IssueRefreshToken, sdkOAuthClient.OverrideOAuth2ClientConfig.IssueRefreshToken, forceReplace).(types.Bool)
	overrideOAuth2ClientConfigTF.IssueRefreshTokenOnRefreshedToken = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.IssueRefreshTokenOnRefreshedToken, sdkOAuthClient.OverrideOAuth2ClientConfig.IssueRefreshTokenOnRefreshedToken, forceReplace).(types.Bool)
	overrideOAuth2ClientConfigTF.OidcClaimsClass = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.OidcClaimsClass, sdkOAuthClient.OverrideOAuth2ClientConfig.OidcClaimsClass, forceReplace).(types.String)
	overrideOAuth2ClientConfigTF.OidcClaimsPluginType = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.OidcClaimsPluginType, sdkOAuthClient.OverrideOAuth2ClientConfig.OidcClaimsPluginType, forceReplace).(types.String)
	overrideOAuth2ClientConfigTF.OidcClaimsScript = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.OidcClaimsScript, sdkOAuthClient.OverrideOAuth2ClientConfig.OidcClaimsScript, forceReplace).(types.String)
	overrideOAuth2ClientConfigTF.OidcMayActScript = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.OidcMayActScript, sdkOAuthClient.OverrideOAuth2ClientConfig.OidcMayActScript, forceReplace).(types.String)
	overrideOAuth2ClientConfigTF.OverrideableOIDCClaims = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.OverrideableOIDCClaims, sdkOAuthClient.OverrideOAuth2ClientConfig.OverrideableOIDCClaims, forceReplace).(types.List)
	overrideOAuth2ClientConfigTF.ProviderOverridesEnabled = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.ProviderOverridesEnabled, sdkOAuthClient.OverrideOAuth2ClientConfig.ProviderOverridesEnabled, forceReplace).(types.Bool)
	overrideOAuth2ClientConfigTF.RemoteConsentServiceId = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.RemoteConsentServiceId, sdkOAuthClient.OverrideOAuth2ClientConfig.RemoteConsentServiceId, forceReplace).(types.String)
	overrideOAuth2ClientConfigTF.ScopesPolicySet = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.ScopesPolicySet, sdkOAuthClient.OverrideOAuth2ClientConfig.ScopesPolicySet, forceReplace).(types.String)
	overrideOAuth2ClientConfigTF.StatelessTokensEnabled = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.StatelessTokensEnabled, sdkOAuthClient.OverrideOAuth2ClientConfig.StatelessTokensEnabled, forceReplace).(types.Bool)
	overrideOAuth2ClientConfigTF.TokenEncryptionEnabled = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.TokenEncryptionEnabled, sdkOAuthClient.OverrideOAuth2ClientConfig.TokenEncryptionEnabled, forceReplace).(types.Bool)
	overrideOAuth2ClientConfigTF.UsePolicyEngineForScope = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.UsePolicyEngineForScope, sdkOAuthClient.OverrideOAuth2ClientConfig.UsePolicyEngineForScope, forceReplace).(types.Bool)
	overrideOAuth2ClientConfigTF.ValidateScopeClass = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.ValidateScopeClass, sdkOAuthClient.OverrideOAuth2ClientConfig.ValidateScopeClass, forceReplace).(types.String)
	overrideOAuth2ClientConfigTF.ValidateScopePluginType = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.ValidateScopePluginType, sdkOAuthClient.OverrideOAuth2ClientConfig.ValidateScopePluginType, forceReplace).(types.String)
	overrideOAuth2ClientConfigTF.ValidateScopeScript = updateTerraformValue(ctx, overrideOAuth2ClientConfigTF.ValidateScopeScript, sdkOAuthClient.OverrideOAuth2ClientConfig.ValidateScopeScript, forceReplace).(types.String)

	if !checkEmptyTFObject(coreOpenIDClientConfigTF) {
		diags := state.SetAttribute(ctx, path.Root("core_open_id_client_config"), &coreOpenIDClientConfigTF)
		if diags.HasError() {
			return errors.New("Failed to set core_open_id_client_config")
		}
	}

	if !checkEmptyTFObject(coreOAuth2ClientConfigTF) {
		diags := state.SetAttribute(ctx, path.Root("core_oauth2_client_config"), &coreOAuth2ClientConfigTF)
		if diags.HasError() {
			return errors.New("Failed to set core_oauth2_client_config")
		}
	}

	if !checkEmptyTFObject(advancedOAuth2ClientConfigTF) {
		diags := state.SetAttribute(ctx, path.Root("advanced_oauth2_client_config"), &advancedOAuth2ClientConfigTF)
		if diags.HasError() {
			return errors.New("Failed to set advanced_oauth2_client_config")
		}
	}

	if !checkEmptyTFObject(signEncOAuth2ClientConfigTF) {
		diags := state.SetAttribute(ctx, path.Root("sign_enc_oauth2_client_config"), &signEncOAuth2ClientConfigTF)
		if diags.HasError() {
			return errors.New("Failed to set sign_enc_oauth2_client_config")
		}
	}

	if !checkEmptyTFObject(coreUmaClientConfigTF) {
		diags := state.SetAttribute(ctx, path.Root("core_uma_client_config"), &coreUmaClientConfigTF)
		if diags.HasError() {
			return errors.New("Failed to set core_uma_client_config")
		}
	}

	if !checkEmptyTFObject(overrideOAuth2ClientConfigTF) {
		diags := state.SetAttribute(ctx, path.Root("override_oauth2_client_config"), &overrideOAuth2ClientConfigTF)
		if diags.HasError() {
			return errors.New("Failed to set override_oauth2_client_config")
		}
	}

	state.SetAttribute(ctx, path.Root("name"), plan.Name)
	state.SetAttribute(ctx, path.Root("admin_mail"), plan.AdminMail)
	state.SetAttribute(ctx, path.Root("user_password_version"), plan.UserPasswordVersion)

	return nil
}

func updateWithReflection(ctx context.Context, tfObject any, sdkObject any, forceReplace bool) any {

	tfObjectReflect := reflect.Indirect(reflect.ValueOf(&tfObject))

	sdkObjectReflect := reflect.ValueOf(sdkObject)
	typeOfSdkObject := sdkObjectReflect.Type()
	for j := 0; j < sdkObjectReflect.NumField(); j++ {
		foundField := tfObjectReflect.Elem().FieldByName(typeOfSdkObject.Field(j).Tag.Get("mapping"))
		if foundField.IsValid() {
			settableField := reflect.ValueOf(&foundField).Elem()
			if settableField.CanSet() {
				updatedValue := updateTerraformValue(ctx, foundField.Interface().(attr.Value), sdkObjectReflect.Field(j).Interface(), forceReplace)
				// TODO : fix type error on assignation
				settableField.Set(reflect.ValueOf(updatedValue))
			}
		}
	}
	return tfObject
}

func checkListIsTyped(ctx context.Context, tfType attr.Value) attr.Value {
	if reflect.TypeOf(tfType) == reflect.TypeOf((*types.List)(nil)).Elem() {
		typedList := tfType.(types.List)
		if typedList.ElementType(ctx) == nil {
			return types.ListNull(types.StringType)
		}
	}
	return nil

}

func updateTerraformValue(ctx context.Context, tfType attr.Value, sdkValue interface{}, forceReplace bool) attr.Value {
	// If the value in Terraform is null, it indicates that we should not update it from the ForgeRock response,
	// as we do not intend to manage it via Terraform.

	nilList := checkListIsTyped(ctx, tfType)
	if nilList != nil && reflect.ValueOf(sdkValue).IsNil() {
		return nilList
	}

	if tfType.IsNull() && !forceReplace {
		return tfType
	}

	if sdkValue == nil || (reflect.TypeOf(sdkValue) == reflect.TypeOf((*sdk.WrappedValue[any])(nil)) && reflect.ValueOf(sdkValue).IsNil()) {
		return tfType
	}

	switch reflect.TypeOf(tfType) {

	case reflect.TypeOf((*types.String)(nil)).Elem():
		if reflect.TypeOf(sdkValue) == reflect.TypeOf((*sdk.WrappedValue[string])(nil)) {
			return types.StringValue(sdkValue.(*sdk.WrappedValue[string]).Value)
		} else {
			return types.StringValue(sdkValue.(string))
		}

	case reflect.TypeOf((*types.Bool)(nil)).Elem():
		if reflect.TypeOf(sdkValue) == reflect.TypeOf((*sdk.WrappedValue[bool])(nil)) {
			return types.BoolValue(sdkValue.(*sdk.WrappedValue[bool]).Value)
		} else {
			return types.BoolValue(sdkValue.(bool))
		}

	case reflect.TypeOf((*types.Int64)(nil)).Elem():
		if reflect.TypeOf(sdkValue) == reflect.TypeOf((*sdk.WrappedValue[int64])(nil)) {
			return types.Int64Value(sdkValue.(*sdk.WrappedValue[int64]).Value)
		} else {
			return types.Int64Value(sdkValue.(int64))
		}

	case reflect.TypeOf((*types.List)(nil)).Elem():
		if reflect.TypeOf(sdkValue) == reflect.TypeOf((*sdk.WrappedValue[[]string])(nil)) {
			return stringListToTerraformList(sdkValue.(*sdk.WrappedValue[[]string]).Value)
		} else {
			return stringListToTerraformList(sdkValue.([]string))
		}

	default:
		panic("Unknown type used in updateTerraformValue " + reflect.TypeOf(tfType).String())
	}
}

/*
 * Convert a list of string to a terraform list
 */
func stringListToTerraformList(list []string) types.List {
	var valueList []attr.Value
	//transform list to a list of types.StringValue
	for _, s := range list {
		valueList = append(valueList, types.StringValue(s))
	}

	finalList, _ := types.ListValue(types.StringType, valueList)

	return finalList
}

/**
 * Convert a terraform list to a list of string
 */
func terraformListToStringList(list types.List) *[]string {

	if list.IsNull() {
		return nil
	}

	var valueList []string
	for _, s := range list.Elements() {
		valueList = append(valueList, s.(types.String).ValueString())
	}

	return &valueList
}

/**
 * Create a Terraform WrappedValue from a pointer to a value
 */
func NewWrappedValue[T any](value *T) *sdk.WrappedValue[T] {

	if value == nil {
		return nil
	}

	return &sdk.WrappedValue[T]{
		Inherited: false,
		Value:     *value,
	}
}
