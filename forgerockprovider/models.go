package forgerockprovider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OAuth2ClientTF struct {
	LastUpdated         types.String `tfsdk:"last_updated"`
	Name                types.String `tfsdk:"name"`
	AdminMail           types.String `tfsdk:"admin_mail"`
	UserPasswordVersion types.Int64  `tfsdk:"user_password_version"`

	AdvancedOAuth2ClientConfigTF types.Object `tfsdk:"advanced_oauth2_client_config"`

	CoreOAuth2ClientConfigTF types.Object `tfsdk:"core_oauth2_client_config"`

	CoreOpenIDClientConfigTF types.Object `tfsdk:"core_open_id_client_config"`

	SignEncOAuth2ClientConfigTF types.Object `tfsdk:"sign_enc_oauth2_client_config"`

	CoreUmaClientConfigTF types.Object `tfsdk:"core_uma_client_config"`

	OverrideOAuth2ClientConfigTF types.Object `tfsdk:"override_oauth2_client_config"`
}

type AdvancedOAuth2ClientConfigTF struct {
	LogoUri                               types.List   `tfsdk:"logo_uri"`
	SubjectType                           types.String `tfsdk:"subject_type"`
	ClientUri                             types.List   `tfsdk:"client_uri"`
	TokenExchangeAuthLevel                types.Int64  `tfsdk:"token_exchange_auth_level"`
	ResponseTypes                         types.List   `tfsdk:"response_types"`
	MixUpMitigation                       types.Bool   `tfsdk:"mix_up_mitigation"`
	CustomProperties                      types.List   `tfsdk:"custom_properties"`
	JavascriptOrigins                     types.List   `tfsdk:"javascript_origins"`
	PolicyUri                             types.List   `tfsdk:"policy_uri"`
	SoftwareVersion                       types.String `tfsdk:"software_version"`
	TosURI                                types.List   `tfsdk:"tos_uri"`
	SectorIdentifierUri                   types.String `tfsdk:"sector_identifier_uri"`
	TokenEndpointAuthMethod               types.String `tfsdk:"token_endpoint_auth_method"`
	RefreshTokenGracePeriod               types.Int64  `tfsdk:"refresh_token_grace_period"`
	IsConsentImplied                      types.Bool   `tfsdk:"is_consent_implied"`
	SoftwareIdentity                      types.String `tfsdk:"software_identity"`
	GrantTypes                            types.List   `tfsdk:"grant_types"`
	Require_pushed_authorization_requests types.Bool   `tfsdk:"require_pushed_authorization_requests"`
	Descriptions                          types.List   `tfsdk:"descriptions"`
	RequestUris                           types.List   `tfsdk:"request_uris"`
	Name                                  types.List   `tfsdk:"name"`
	Contacts                              types.List   `tfsdk:"contacts"`
	UpdateAccessToken                     types.String `tfsdk:"update_access_token"`
}

type CoreOAuth2ClientConfigTF struct {
	Status                       types.String `tfsdk:"status"`
	ClientName                   types.List   `tfsdk:"client_name"`
	ClientType                   types.String `tfsdk:"client_type"`
	LoopbackInterfaceRedirection types.Bool   `tfsdk:"loopback_interface_redirection"`
	DefaultScopes                types.List   `tfsdk:"default_scopes"`
	Agentgroup                   types.String `tfsdk:"agentgroup"`
	RefreshTokenLifetime         types.Int64  `tfsdk:"refresh_token_lifetime"`
	Scopes                       types.List   `tfsdk:"scopes"`
	AccessTokenLifetime          types.Int64  `tfsdk:"access_token_lifetime"`
	RedirectionUris              types.List   `tfsdk:"redirection_uris"`
	AuthorizationCodeLifetime    types.Int64  `tfsdk:"authorization_code_lifetime"`
}

type CoreOpenIDClientConfigTF struct {
	Claims                              types.List   `tfsdk:"claims"`
	BackchannelLogoutUri                types.String `tfsdk:"backchannel_logout_uri"`
	DefaultAcrValues                    types.List   `tfsdk:"default_acr_values"`
	JwtTokenLifetime                    types.Int64  `tfsdk:"jwt_token_lifetime"`
	DefaultMaxAgeEnabled                types.Bool   `tfsdk:"default_max_age_enabled"`
	ClientSessionUri                    types.String `tfsdk:"client_session_uri"`
	DefaultMaxAge                       types.Int64  `tfsdk:"default_max_age"`
	PostLogoutRedirectUri               types.List   `tfsdk:"post_logout_redirect_uri"`
	Backchannel_logout_session_required types.Bool   `tfsdk:"backchannel_logout_session_required"`
}

type SignEncOAuth2ClientConfigTF struct {
	AuthorizationResponseEncryptionMethod                  types.String `tfsdk:"authorization_response_encryption_method"`
	AuthorizationResponseEncryptionAlgorithm               types.String `tfsdk:"authorization_response_encryption_algorithm"`
	AuthorizationResponseSigningAlgorithm                  types.String `tfsdk:"authorization_response_signing_algorithm"`
	ClientJwtPublicKey                                     types.String `tfsdk:"client_jwt_public_key"`
	IDTokenEncryptionAlgorithm                             types.String `tfsdk:"id_token_encryption_algorithm"`
	IDTokenEncryptionEnabled                               types.Bool   `tfsdk:"id_token_encryption_enabled"`
	IDTokenEncryptionMethod                                types.String `tfsdk:"id_token_encryption_method"`
	IDTokenPublicEncryptionKey                             types.String `tfsdk:"id_token_public_encryption_key"`
	IDTokenSignedResponseAlg                               types.String `tfsdk:"id_token_signed_response_alg"`
	JwkSet                                                 types.String `tfsdk:"jwk_set"`
	JwkStoreCacheMissCacheTime                             types.Int64  `tfsdk:"jwk_store_cache_miss_cache_time"`
	JwksCacheTimeout                                       types.Int64  `tfsdk:"jwks_cache_timeout"`
	JwksUri                                                types.String `tfsdk:"jwks_uri"`
	MTLSCertificateBoundAccessTokens                       types.Bool   `tfsdk:"mtls_certificate_bound_access_tokens"`
	MTLSSubjectDN                                          types.String `tfsdk:"mtls_subject_dn"`
	MTLSTrustedCert                                        types.String `tfsdk:"mtls_trusted_cert"`
	PublicKeyLocation                                      types.String `tfsdk:"public_key_location"`
	RequestParameterEncryptedAlg                           types.String `tfsdk:"request_parameter_encrypted_alg"`
	RequestParameterEncryptedEncryptionAlgorithm           types.String `tfsdk:"request_parameter_encrypted_encryption_algorithm"`
	RequestParameterSignedAlg                              types.String `tfsdk:"request_parameter_signed_alg"`
	TokenEndpointAuthSigningAlgorithm                      types.String `tfsdk:"token_endpoint_auth_signing_algorithm"`
	TokenIntrospectionEncryptedResponseAlg                 types.String `tfsdk:"token_introspection_encrypted_response_alg"`
	TokenIntrospectionEncryptedResponseEncryptionAlgorithm types.String `tfsdk:"token_introspection_encrypted_response_encryption_algorithm"`
	TokenIntrospectionResponseFormat                       types.String `tfsdk:"token_introspection_response_format"`
	TokenIntrospectionSignedResponseAlg                    types.String `tfsdk:"token_introspection_signed_response_alg"`
	UserinfoEncryptedResponseAlg                           types.String `tfsdk:"userinfo_encrypted_response_alg"`
	UserinfoEncryptedResponseEncryptionAlgorithm           types.String `tfsdk:"userinfo_encrypted_response_encryption_algorithm"`
	UserinfoResponseFormat                                 types.String `tfsdk:"userinfo_response_format"`
	UserinfoSignedResponseAlg                              types.String `tfsdk:"userinfo_signed_response_alg"`
}

type CoreUmaClientConfigTF struct {
	ClaimsRedirectionUris types.List `tfsdk:"claims_redirection_uris"`
}

type OverrideOAuth2ClientConfigTF struct {
	AccessTokenMayActScript                 types.String `tfsdk:"access_token_may_act_script"`
	AccessTokenModificationPluginType       types.String `tfsdk:"access_token_modification_plugin_type"`
	AccessTokenModificationScript           types.String `tfsdk:"access_token_modification_script"`
	AccessTokenModifierClass                types.String `tfsdk:"access_token_modifier_class"`
	AuthorizeEndpointDataProviderClass      types.String `tfsdk:"authorize_endpoint_data_provider_class"`
	AuthorizeEndpointDataProviderPluginType types.String `tfsdk:"authorize_endpoint_data_provider_plugin_type"`
	AuthorizeEndpointDataProviderScript     types.String `tfsdk:"authorize_endpoint_data_provider_script"`
	ClientsCanSkipConsent                   types.Bool   `tfsdk:"clients_can_skip_consent"`
	CustomLoginUrlTemplate                  types.String `tfsdk:"custom_login_url_template"`
	EvaluateScopeClass                      types.String `tfsdk:"evaluate_scope_class"`
	EvaluateScopePluginType                 types.String `tfsdk:"evaluate_scope_plugin_type"`
	EvaluateScopeScript                     types.String `tfsdk:"evaluate_scope_script"`
	EnableRemoteConsent                     types.Bool   `tfsdk:"enable_remote_consent"`
	IssueRefreshToken                       types.Bool   `tfsdk:"issue_refresh_token"`
	IssueRefreshTokenOnRefreshedToken       types.Bool   `tfsdk:"issue_refresh_token_on_refreshed_token"`
	OidcClaimsClass                         types.String `tfsdk:"oidc_claims_class"`
	OidcClaimsPluginType                    types.String `tfsdk:"oidc_claims_plugin_type"`
	OidcClaimsScript                        types.String `tfsdk:"oidc_claims_script"`
	OidcMayActScript                        types.String `tfsdk:"oidc_may_act_script"`
	OverrideableOIDCClaims                  types.List   `tfsdk:"overrideable_oidc_claims"`
	ProviderOverridesEnabled                types.Bool   `tfsdk:"provider_overrides_enabled"`
	RemoteConsentServiceId                  types.String `tfsdk:"remote_consent_service_id"`
	ScopesPolicySet                         types.String `tfsdk:"scopes_policy_set"`
	StatelessTokensEnabled                  types.Bool   `tfsdk:"stateless_tokens_enabled"`
	TokenEncryptionEnabled                  types.Bool   `tfsdk:"token_encryption_enabled"`
	UsePolicyEngineForScope                 types.Bool   `tfsdk:"use_policy_engine_for_scope"`
	ValidateScopeClass                      types.String `tfsdk:"validate_scope_class"`
	ValidateScopePluginType                 types.String `tfsdk:"validate_scope_plugin_type"`
	ValidateScopeScript                     types.String `tfsdk:"validate_scope_script"`
}
