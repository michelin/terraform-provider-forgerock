package forgerockprovider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func createOAuthClientSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"admin_mail": schema.StringAttribute{
				Required: true,
			},
			"user_password_version":         defaultInt64SchemaAttribute(),
			"advanced_oauth2_client_config": createAdvancedOAuth2ClientConfigSchema(),
			"core_open_id_client_config":    createCoreOpenIDClientConfigSchema(),
			"core_oauth2_client_config":     createCoreOAuth2ClientConfigSchema(),
			"sign_enc_oauth2_client_config": createSignEncOAuth2ClientConfigSchema(),
			"core_uma_client_config":        createCoreUmaClientConfigSchema(),
			"override_oauth2_client_config": createOverrideOAuth2ClientConfigSchema(),
		},
	}
}

func createAdvancedOAuth2ClientConfigSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"logo_uri":                              defaultListSchemaAttribute(),
			"subject_type":                          defaultStringSchemaAttribute(),
			"client_uri":                            defaultListSchemaAttribute(),
			"token_exchange_auth_level":             defaultInt64SchemaAttribute(),
			"response_types":                        defaultListSchemaAttribute(),
			"mix_up_mitigation":                     defaultBoolSchemaAttribute(),
			"custom_properties":                     defaultListSchemaAttribute(),
			"javascript_origins":                    defaultListSchemaAttribute(),
			"policy_uri":                            defaultListSchemaAttribute(),
			"software_version":                      defaultStringSchemaAttribute(),
			"tos_uri":                               defaultListSchemaAttribute(),
			"sector_identifier_uri":                 defaultStringSchemaAttribute(),
			"token_endpoint_auth_method":            defaultStringSchemaAttribute(),
			"refresh_token_grace_period":            defaultInt64SchemaAttribute(),
			"is_consent_implied":                    defaultBoolSchemaAttribute(),
			"software_identity":                     defaultStringSchemaAttribute(),
			"grant_types":                           defaultListSchemaAttribute(),
			"require_pushed_authorization_requests": defaultBoolSchemaAttribute(),
			"descriptions":                          defaultListSchemaAttribute(),
			"request_uris":                          defaultListSchemaAttribute(),
			"name":                                  defaultListSchemaAttribute(),
			"contacts":                              defaultListSchemaAttribute(),
			"update_access_token":                   defaultStringSchemaAttribute(),
		},
	}
}

func createCoreOAuth2ClientConfigSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"status":                         defaultStringSchemaAttribute(),
			"client_name":                    defaultListSchemaAttribute(),
			"client_type":                    defaultStringSchemaAttribute(),
			"loopback_interface_redirection": defaultBoolSchemaAttribute(),
			"default_scopes":                 defaultListSchemaAttribute(),
			"agentgroup":                     defaultStringSchemaAttribute(),
			"refresh_token_lifetime":         defaultInt64SchemaAttribute(),
			"scopes":                         defaultListSchemaAttribute(),
			"access_token_lifetime":          defaultInt64SchemaAttribute(),
			"redirection_uris":               defaultListSchemaAttribute(),
			"authorization_code_lifetime":    defaultInt64SchemaAttribute(),
		},
	}
}

func createCoreOpenIDClientConfigSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"claims":                              defaultListSchemaAttribute(),
			"backchannel_logout_uri":              defaultStringSchemaAttribute(),
			"default_acr_values":                  defaultListSchemaAttribute(),
			"jwt_token_lifetime":                  defaultInt64SchemaAttribute(),
			"default_max_age_enabled":             defaultBoolSchemaAttribute(),
			"client_session_uri":                  defaultStringSchemaAttribute(),
			"default_max_age":                     defaultInt64SchemaAttribute(),
			"post_logout_redirect_uri":            defaultListSchemaAttribute(),
			"backchannel_logout_session_required": defaultBoolSchemaAttribute(),
		},
	}
}

func createSignEncOAuth2ClientConfigSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"authorization_response_encryption_algorithm":                 defaultStringSchemaAttribute(),
			"authorization_response_encryption_method":                    defaultStringSchemaAttribute(),
			"authorization_response_signing_algorithm":                    defaultStringSchemaAttribute(),
			"client_jwt_public_key":                                       defaultStringSchemaAttribute(),
			"id_token_encryption_algorithm":                               defaultStringSchemaAttribute(),
			"id_token_encryption_enabled":                                 defaultBoolSchemaAttribute(),
			"id_token_encryption_method":                                  defaultStringSchemaAttribute(),
			"id_token_public_encryption_key":                              defaultStringSchemaAttribute(),
			"id_token_signed_response_alg":                                defaultStringSchemaAttribute(),
			"jwk_set":                                                     defaultStringSchemaAttribute(),
			"jwk_store_cache_miss_cache_time":                             defaultInt64SchemaAttribute(),
			"jwks_cache_timeout":                                          defaultInt64SchemaAttribute(),
			"jwks_uri":                                                    defaultStringSchemaAttribute(),
			"mtls_certificate_bound_access_tokens":                        defaultBoolSchemaAttribute(),
			"mtls_subject_dn":                                             defaultStringSchemaAttribute(),
			"mtls_trusted_cert":                                           defaultStringSchemaAttribute(),
			"public_key_location":                                         defaultStringSchemaAttribute(),
			"request_parameter_encrypted_alg":                             defaultStringSchemaAttribute(),
			"request_parameter_encrypted_encryption_algorithm":            defaultStringSchemaAttribute(),
			"request_parameter_signed_alg":                                defaultStringSchemaAttribute(),
			"token_endpoint_auth_signing_algorithm":                       defaultStringSchemaAttribute(),
			"token_introspection_encrypted_response_alg":                  defaultStringSchemaAttribute(),
			"token_introspection_encrypted_response_encryption_algorithm": defaultStringSchemaAttribute(),
			"token_introspection_response_format":                         defaultStringSchemaAttribute(),
			"token_introspection_signed_response_alg":                     defaultStringSchemaAttribute(),
			"userinfo_encrypted_response_alg":                             defaultStringSchemaAttribute(),
			"userinfo_encrypted_response_encryption_algorithm":            defaultStringSchemaAttribute(),
			"userinfo_response_format":                                    defaultStringSchemaAttribute(),
			"userinfo_signed_response_alg":                                defaultStringSchemaAttribute(),
		},
	}
}

func createCoreUmaClientConfigSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"claims_redirection_uris": defaultListSchemaAttribute(),
		},
	}
}

func createOverrideOAuth2ClientConfigSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"access_token_may_act_script":                  defaultStringSchemaAttribute(),
			"access_token_modification_plugin_type":        defaultStringSchemaAttribute(),
			"access_token_modification_script":             defaultStringSchemaAttribute(),
			"access_token_modifier_class":                  defaultStringSchemaAttribute(),
			"authorize_endpoint_data_provider_plugin_type": defaultStringSchemaAttribute(),
			"authorize_endpoint_data_provider_class":       defaultStringSchemaAttribute(),
			"authorize_endpoint_data_provider_script":      defaultStringSchemaAttribute(),
			"clients_can_skip_consent":                     defaultBoolSchemaAttribute(),
			"custom_login_url_template":                    defaultStringSchemaAttribute(),
			"enable_remote_consent":                        defaultBoolSchemaAttribute(),
			"evaluate_scope_class":                         defaultStringSchemaAttribute(),
			"evaluate_scope_plugin_type":                   defaultStringSchemaAttribute(),
			"evaluate_scope_script":                        defaultStringSchemaAttribute(),
			"issue_refresh_token":                          defaultBoolSchemaAttribute(),
			"issue_refresh_token_on_refreshed_token":       defaultBoolSchemaAttribute(),
			"oidc_claims_class":                            defaultStringSchemaAttribute(),
			"oidc_claims_plugin_type":                      defaultStringSchemaAttribute(),
			"oidc_claims_script":                           defaultStringSchemaAttribute(),
			"oidc_may_act_script":                          defaultStringSchemaAttribute(),
			"overrideable_oidc_claims":                     defaultListSchemaAttribute(),
			"provider_overrides_enabled":                   defaultBoolSchemaAttribute(),
			"remote_consent_service_id":                    defaultStringSchemaAttribute(),
			"scopes_policy_set":                            defaultStringSchemaAttribute(),
			"stateless_tokens_enabled":                     defaultBoolSchemaAttribute(),
			"token_encryption_enabled":                     defaultBoolSchemaAttribute(),
			"use_policy_engine_for_scope":                  defaultBoolSchemaAttribute(),
			"validate_scope_class":                         defaultStringSchemaAttribute(),
			"validate_scope_plugin_type":                   defaultStringSchemaAttribute(),
			"validate_scope_script":                        defaultStringSchemaAttribute(),
		},
	}
}

/**
* Overview of default*SchemaAttribute functions:
* By default, all attributes are set as optional and computed. This allows fields in Terraform files to be left empty.
* When a resource is created, default values from Forgerock are assigned and then saved to the state file.
* On subsequent updates, these default values in the state will be utilized by the planModifier and passed to the Forgerock API, preventing the sending of null values.
**/

func defaultStringSchemaAttribute() schema.Attribute {
	return schema.StringAttribute{
		Optional: true,
	}
}

func defaultBoolSchemaAttribute() schema.Attribute {
	return schema.BoolAttribute{
		Optional: true,
	}
}

func defaultInt64SchemaAttribute() schema.Attribute {
	return schema.Int64Attribute{
		Optional: true,
	}
}

func defaultListSchemaAttribute() schema.ListAttribute {
	return schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	}
}
