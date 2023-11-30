package sdk

type OAuth2Client struct {
	AdvancedOAuth2ClientConfig AdvancedOAuth2ClientConfig `json:"advancedOAuth2ClientConfig,omitempty"`

	CoreOpenIDClientConfig CoreOpenIDClientConfig `json:"coreOpenIDClientConfig,omitempty"`

	CoreOAuth2ClientConfig CoreOAuth2ClientConfig `json:"coreOAuth2ClientConfig,omitempty"`

	SignEncOAuth2ClientConfig SignEncOAuth2ClientConfig `json:"signEncOAuth2ClientConfig,omitempty"`

	CoreUmaClientConfig CoreUmaClientConfig `json:"coreUmaClientConfig,omitempty"`

	OverrideOAuth2ClientConfig OverrideOAuth2ClientConfig `json:"overrideOAuth2ClientConfig,omitempty"`
}

type AdvancedOAuth2ClientConfig struct {
	LogoURI                            *WrappedValue[[]string] `json:"logoUri,omitempty" mapping:"LogoUri"`
	SubjectType                        *WrappedValue[string]   `json:"subjectType,omitempty" mapping:"SubjectType"`
	ClientURI                          *WrappedValue[[]string] `json:"clientUri,omitempty" mapping:"ClientUri"`
	TokenExchangeAuthLevel             *WrappedValue[int64]    `json:"tokenExchangeAuthLevel,omitempty" mapping:"TokenExchangeAuthLevel"`
	ResponseTypes                      *WrappedValue[[]string] `json:"responseTypes,omitempty" mapping:"ResponseTypes"`
	MixUpMitigation                    *WrappedValue[bool]     `json:"mixUpMitigation,omitempty" mapping:"MixUpMitigation"`
	CustomProperties                   *WrappedValue[[]string] `json:"customProperties,omitempty" mapping:"CustomProperties"`
	JavascriptOrigins                  *WrappedValue[[]string] `json:"javascriptOrigins,omitempty" mapping:"JavascriptOrigins"`
	PolicyURI                          *WrappedValue[[]string] `json:"policyUri,omitempty" mapping:"PolicyUri"`
	SoftwareVersion                    *WrappedValue[string]   `json:"softwareVersion,omitempty" mapping:"SoftwareVersion"`
	TosURI                             *WrappedValue[[]string] `json:"tosURI,omitempty" mapping:"TosURI"`
	SectorIdentifierUri                *WrappedValue[string]   `json:"sectorIdentifierUri,omitempty" mapping:"SectorIdentifierUri"`
	TokenEndpointAuthMethod            *WrappedValue[string]   `json:"tokenEndpointAuthMethod,omitempty" mapping:"TokenEndpointAuthMethod"`
	RefreshTokenGracePeriod            *WrappedValue[int64]    `json:"refreshTokenGracePeriod,omitempty" mapping:"RefreshTokenGracePeriod"`
	IsConsentImplied                   *WrappedValue[bool]     `json:"isConsentImplied,omitempty" mapping:"IsConsentImplied"`
	SoftwareIdentity                   *WrappedValue[string]   `json:"softwareIdentity,omitempty" mapping:"SoftwareIdentity"`
	GrantTypes                         *WrappedValue[[]string] `json:"grantTypes,omitempty" mapping:"GrantTypes"`
	RequirePushedAuthorizationRequests *WrappedValue[bool]     `json:"require_pushed_authorization_requests,omitempty" mapping:"Require_pushed_authorization_requests"`
	Descriptions                       *WrappedValue[[]string] `json:"descriptions,omitempty" mapping:"Descriptions"`
	RequestUris                        *WrappedValue[[]string] `json:"requestUris,omitempty" mapping:"RequestUris"`
	Name                               *WrappedValue[[]string] `json:"name,omitempty" mapping:"Name"`
	Contacts                           *WrappedValue[[]string] `json:"contacts,omitempty" mapping:"Contacts"`
	UpdateAccessToken                  *WrappedValue[string]   `json:"updateAccessToken,omitempty" mapping:"UpdateAccessToken"`
}

type CoreOAuth2ClientConfig struct {
	Userpassword                 *WrappedValue[string]   `json:"userpassword,omitempty"` // pas de champ correspondant dans CoreOAuth2ClientConfigTF
	Status                       *WrappedValue[string]   `json:"status,omitempty" mapping:"Status"`
	ClientName                   *WrappedValue[[]string] `json:"clientName,omitempty" mapping:"ClientName"`
	ClientType                   *WrappedValue[string]   `json:"clientType,omitempty" mapping:"ClientType"`
	LoopbackInterfaceRedirection *WrappedValue[bool]     `json:"loopbackInterfaceRedirection,omitempty" mapping:"LoopbackInterfaceRedirection"`
	DefaultScopes                *WrappedValue[[]string] `json:"defaultScopes,omitempty" mapping:"DefaultScopes"`
	Agentgroup                   *WrappedValue[string]   `json:"agentgroup,omitempty" mapping:"Agentgroup"`
	RefreshTokenLifetime         *WrappedValue[int64]    `json:"refreshTokenLifetime,omitempty" mapping:"RefreshTokenLifetime"`
	Scopes                       *WrappedValue[[]string] `json:"scopes,omitempty" mapping:"Scopes"`
	AccessTokenLifetime          *WrappedValue[int64]    `json:"accessTokenLifetime,omitempty" mapping:"AccessTokenLifetime"`
	RedirectionUris              *WrappedValue[[]string] `json:"redirectionUris,omitempty" mapping:"RedirectionUris"`
	AuthorizationCodeLifetime    *WrappedValue[int64]    `json:"authorizationCodeLifetime,omitempty" mapping:"AuthorizationCodeLifetime"`
}

type CoreOpenIDClientConfig struct {
	Claims                           *WrappedValue[[]string] `json:"claims,omitempty" mapping:"Claims"`
	BackchannelLogoutUri             *WrappedValue[string]   `json:"backchannel_logout_uri,omitempty" mapping:"BackchannelLogoutUri"`
	DefaultAcrValues                 *WrappedValue[[]string] `json:"defaultAcrValues,omitempty" mapping:"DefaultAcrValues"`
	JwtTokenLifetime                 *WrappedValue[int64]    `json:"jwtTokenLifetime,omitempty" mapping:"JwtTokenLifetime"`
	DefaultMaxAgeEnabled             *WrappedValue[bool]     `json:"defaultMaxAgeEnabled,omitempty" mapping:"DefaultMaxAgeEnabled"`
	ClientSessionUri                 *WrappedValue[string]   `json:"clientSessionUri,omitempty" mapping:"ClientSessionUri"`
	DefaultMaxAge                    *WrappedValue[int64]    `json:"defaultMaxAge,omitempty" mapping:"DefaultMaxAge"`
	PostLogoutRedirectUri            *WrappedValue[[]string] `json:"postLogoutRedirectUri,omitempty" mapping:"PostLogoutRedirectUri"`
	BackchannelLogoutSessionRequired *WrappedValue[bool]     `json:"backchannel_logout_session_required,omitempty" mapping:"Backchannel_logout_session_required"`
}

type SignEncOAuth2ClientConfig struct {
	TokenEndpointAuthSigningAlgorithm                      *WrappedValue[string] `json:"tokenEndpointAuthSigningAlgorithm,omitempty" mapping:"TokenEndpointAuthSigningAlgorithm"`
	IDTokenEncryptionEnabled                               *WrappedValue[bool]   `json:"idTokenEncryptionEnabled,omitempty" mapping:"IDTokenEncryptionEnabled"`
	TokenIntrospectionEncryptedResponseEncryptionAlgorithm *WrappedValue[string] `json:"tokenIntrospectionEncryptedResponseEncryptionAlgorithm,omitempty" mapping:"TokenIntrospectionEncryptedResponseEncryptionAlgorithm"`
	RequestParameterSignedAlg                              *WrappedValue[string] `json:"requestParameterSignedAlg,omitempty" mapping:"RequestParameterSignedAlg"`
	AuthorizationResponseSigningAlgorithm                  *WrappedValue[string] `json:"authorizationResponseSigningAlgorithm,omitempty" mapping:"AuthorizationResponseSigningAlgorithm"`
	ClientJwtPublicKey                                     *WrappedValue[string] `json:"clientJwtPublicKey,omitempty" mapping:"ClientJwtPublicKey"`
	IDTokenPublicEncryptionKey                             *WrappedValue[string] `json:"idTokenPublicEncryptionKey,omitempty" mapping:"IDTokenPublicEncryptionKey"`
	MTLSSubjectDN                                          *WrappedValue[string] `json:"mTLSSubjectDN,omitempty" mapping:"MTLSSubjectDN"`
	JwkStoreCacheMissCacheTime                             *WrappedValue[int64]  `json:"jwkStoreCacheMissCacheTime,omitempty" mapping:"JwkStoreCacheMissCacheTime"`
	JwkSet                                                 *WrappedValue[string] `json:"jwkSet,omitempty" mapping:"JwkSet"`
	IDTokenEncryptionMethod                                *WrappedValue[string] `json:"idTokenEncryptionMethod,omitempty" mapping:"IDTokenEncryptionMethod"`
	JwksUri                                                *WrappedValue[string] `json:"jwksUri,omitempty" mapping:"JwksUri"`
	TokenIntrospectionEncryptedResponseAlg                 *WrappedValue[string] `json:"tokenIntrospectionEncryptedResponseAlg,omitempty" mapping:"TokenIntrospectionEncryptedResponseAlg"`
	AuthorizationResponseEncryptionMethod                  *WrappedValue[string] `json:"authorizationResponseEncryptionMethod,omitempty" mapping:"AuthorizationResponseEncryptionMethod"`
	MTLSCertificateBoundAccessTokens                       *WrappedValue[bool]   `json:"mTLSCertificateBoundAccessTokens,omitempty" mapping:"MTLSCertificateBoundAccessTokens"`
	UserinfoResponseFormat                                 *WrappedValue[string] `json:"userinfoResponseFormat,omitempty" mapping:"UserinfoResponseFormat"`
	PublicKeyLocation                                      *WrappedValue[string] `json:"publicKeyLocation,omitempty" mapping:"PublicKeyLocation"`
	TokenIntrospectionResponseFormat                       *WrappedValue[string] `json:"tokenIntrospectionResponseFormat,omitempty" mapping:"TokenIntrospectionResponseFormat"`
	RequestParameterEncryptedEncryptionAlgorithm           *WrappedValue[string] `json:"requestParameterEncryptedEncryptionAlgorithm,omitempty" mapping:"RequestParameterEncryptedEncryptionAlgorithm"`
	UserinfoSignedResponseAlg                              *WrappedValue[string] `json:"userinfoSignedResponseAlg,omitempty" mapping:"UserinfoSignedResponseAlg"`
	IDTokenEncryptionAlgorithm                             *WrappedValue[string] `json:"idTokenEncryptionAlgorithm,omitempty" mapping:"IDTokenEncryptionAlgorithm"`
	RequestParameterEncryptedAlg                           *WrappedValue[string] `json:"requestParameterEncryptedAlg,omitempty" mapping:"RequestParameterEncryptedAlg"`
	AuthorizationResponseEncryptionAlgorithm               *WrappedValue[string] `json:"authorizationResponseEncryptionAlgorithm,omitempty" mapping:"AuthorizationResponseEncryptionAlgorithm"`
	MTLSTrustedCert                                        *WrappedValue[string] `json:"mTLSTrustedCert,omitempty" mapping:"MTLSTrustedCert"`
	JwksCacheTimeout                                       *WrappedValue[int64]  `json:"jwksCacheTimeout,omitempty" mapping:"JwksCacheTimeout"`
	UserinfoEncryptedResponseAlg                           *WrappedValue[string] `json:"userinfoEncryptedResponseAlg,omitempty" mapping:"UserinfoEncryptedResponseAlg"`
	IDTokenSignedResponseAlg                               *WrappedValue[string] `json:"idTokenSignedResponseAlg,omitempty" mapping:"IDTokenSignedResponseAlg"`
	UserinfoEncryptedResponseEncryptionAlgorithm           *WrappedValue[string] `json:"userinfoEncryptedResponseEncryptionAlgorithm,omitempty" mapping:"UserinfoEncryptedResponseEncryptionAlgorithm"`
	TokenIntrospectionSignedResponseAlg                    *WrappedValue[string] `json:"tokenIntrospectionSignedResponseAlg,omitempty" mapping:"TokenIntrospectionSignedResponseAlg"`
}

type CoreUmaClientConfig struct {
	ClaimsRedirectionUris *WrappedValue[[]string] `json:"claimsRedirectionUris,omitempty" mapping:"ClaimsRedirectionUris"`
}

type OverrideOAuth2ClientConfig struct {
	AccessTokenMayActScript                 string   `json:"accessTokenMayActScript,omitempty" mapping:"AccessTokenMayActScript"`
	AccessTokenModificationPluginType       string   `json:"accessTokenModificationPluginType,omitempty" mapping:"AccessTokenModificationPluginType"`
	AccessTokenModificationScript           string   `json:"accessTokenModificationScript,omitempty" mapping:"AccessTokenModificationScript"`
	AccessTokenModifierClass                string   `json:"accessTokenModifierClass,omitempty" mapping:"AccessTokenModifierClass"`
	AuthorizeEndpointDataProviderClass      string   `json:"authorizeEndpointDataProviderClass,omitempty" mapping:"AuthorizeEndpointDataProviderClass"`
	AuthorizeEndpointDataProviderPluginType string   `json:"authorizeEndpointDataProviderPluginType,omitempty" mapping:"AuthorizeEndpointDataProviderPluginType"`
	AuthorizeEndpointDataProviderScript     string   `json:"authorizeEndpointDataProviderScript,omitempty" mapping:"AuthorizeEndpointDataProviderScript"`
	ClientsCanSkipConsent                   bool     `json:"clientsCanSkipConsent,omitempty" mapping:"ClientsCanSkipConsent"`
	CustomLoginUrlTemplate                  string   `json:"customLoginUrlTemplate,omitempty" mapping:"CustomLoginUrlTemplate"`
	EnableRemoteConsent                     bool     `json:"enableRemoteConsent,omitempty" mapping:"EnableRemoteConsent"`
	EvaluateScopeClass                      string   `json:"evaluateScopeClass,omitempty" mapping:"EvaluateScopeClass"`
	EvaluateScopePluginType                 string   `json:"evaluateScopePluginType,omitempty" mapping:"EvaluateScopePluginType"`
	EvaluateScopeScript                     string   `json:"evaluateScopeScript,omitempty" mapping:"EvaluateScopeScript"`
	IssueRefreshToken                       bool     `json:"issueRefreshToken,omitempty" mapping:"IssueRefreshToken"`
	IssueRefreshTokenOnRefreshedToken       bool     `json:"issueRefreshTokenOnRefreshedToken,omitempty" mapping:"IssueRefreshTokenOnRefreshedToken"`
	OidcClaimsClass                         string   `json:"oidcClaimsClass,omitempty" mapping:"OidcClaimsClass"`
	OidcClaimsPluginType                    string   `json:"oidcClaimsPluginType,omitempty" mapping:"OidcClaimsPluginType"`
	OidcClaimsScript                        string   `json:"oidcClaimsScript,omitempty" mapping:"OidcClaimsScript"`
	OidcMayActScript                        string   `json:"oidcMayActScript,omitempty" mapping:"OidcMayActScript"`
	OverrideableOIDCClaims                  []string `json:"overrideableOIDCClaims,omitempty" mapping:"OverrideableOIDCClaims"`
	ProviderOverridesEnabled                bool     `json:"providerOverridesEnabled,omitempty" mapping:"ProviderOverridesEnabled"`
	RemoteConsentServiceId                  string   `json:"remoteConsentServiceId,omitempty" mapping:"RemoteConsentServiceId"`
	ScopesPolicySet                         string   `json:"scopesPolicySet,omitempty" mapping:"ScopesPolicySet"`
	StatelessTokensEnabled                  bool     `json:"statelessTokensEnabled,omitempty" mapping:"StatelessTokensEnabled"`
	TokenEncryptionEnabled                  bool     `json:"tokenEncryptionEnabled,omitempty" mapping:"TokenEncryptionEnabled"`
	UsePolicyEngineForScope                 bool     `json:"usePolicyEngineForScope,omitempty" mapping:"UsePolicyEngineForScope"`
	ValidateScopeClass                      string   `json:"validateScopeClass,omitempty" mapping:"ValidateScopeClass"`
	ValidateScopePluginType                 string   `json:"validateScopePluginType,omitempty" mapping:"ValidateScopePluginType"`
	ValidateScopeScript                     string   `json:"validateScopeScript,omitempty" mapping:"ValidateScopeScript"`
}

type WrappedValue[T any] struct {
	Inherited bool `json:"inherited"`
	Value     T    `json:"value"`
}

type WrappedValueNotTyped struct {
	Inherited bool `json:"inherited"`
	Value     any  `json:"value"`
}
