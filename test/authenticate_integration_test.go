//go:build integration_tests
// +build integration_tests

package test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	TestInit(m)
}

func TestCheckNotExistClientIntegration(t *testing.T) {

	must_not_exist_clients_names := []string{
		"INTEGRATION_TEST_AUTO_CLIENT_1",
		"INTEGRATION_TEST_AUTO_CLIENT_2",
		"INTEGRATION_TEST_AUTO_CLIENT_3"}

	for _, name := range must_not_exist_clients_names {
		_, err := sdkClient.ReadClient(name)
		if err != nil && strings.Contains(err.Error(), "404") {
			continue
		}
		t.Errorf("Client should not exist %s %s", name, err)
	}
}

func TestClientDefaultValues(t *testing.T) {
	executeCommand("terraform", "-chdir=./usecase1", "plan")
	executeCommand("terraform", "-chdir=./usecase1", "apply", "-auto-approve")

	client, err := sdkClient.ReadClient("INTEGRATION_TEST_AUTO_CLIENT_1")
	if err != nil {
		t.Errorf("Client should exist %s", err)
	}

	clientTemplate, err := sdkClient.GetClientTemplate()
	if err != nil {
		t.Errorf("Fail get template %s", err)
	}

	clientTemplate.CoreOAuth2ClientConfig.Userpassword = client.CoreOAuth2ClientConfig.Userpassword
	clientTemplate.CoreOAuth2ClientConfig.Status = client.CoreOAuth2ClientConfig.Status
	clientTemplate.CoreOAuth2ClientConfig.Agentgroup = client.CoreOAuth2ClientConfig.Agentgroup

	assert.Equal(t, client, clientTemplate)
}

func TestClientCustomValues(t *testing.T) {
	executeCommand("terraform", "-chdir=./usecase1", "plan")
	executeCommand("terraform", "-chdir=./usecase1", "apply", "-auto-approve")

	client, err := sdkClient.ReadClient("INTEGRATION_TEST_AUTO_CLIENT_2")
	if err != nil {
		t.Errorf("Client should exist %s", err)
	}

	assert.Equal(t, client.CoreOAuth2ClientConfig.Status.Value, "Active")
	assert.Equal(t, client.CoreOAuth2ClientConfig.Scopes.Value, []string{"profile", "openid", "test2"})
	assert.Equal(t, client.CoreOAuth2ClientConfig.ClientType.Value, "Confidential")

	assert.Equal(t, client.AdvancedOAuth2ClientConfig.TokenEndpointAuthMethod.Value, "private_key_jwt")
	assert.Equal(t, client.AdvancedOAuth2ClientConfig.GrantTypes.Value, []string{"client_credentials"})

	assert.Equal(t, client.SignEncOAuth2ClientConfig.PublicKeyLocation.Value, "jwks")
	assert.Equal(t, client.SignEncOAuth2ClientConfig.JwkSet.Value, "{   \"keys\": [       {           \"kty\": \"RSA\",           \"e\": \"AQAB\",            \"use\": \"sig\",           \"kid\": \"23c50663-6902-4cdb-9ef9-a15f9bb10eeb\",          \"n\": \"xsczx0C6tMxxZquC-bsAfCd7HGx41c8aje5pn5Oys6lSUKEJ7mqZFWQi86IkOWwjoorSpczP1xOKhwN5_80_Yi_zAYs7iDeENDXt-O5bjNdagC3nxgGYoefSaJgmKmK3Da20b_YcIWGEddS_IK4QRtgLEcY3wh6-9fUvHsbCSarPGdm34E4F1jAaiuC1dTyT5qUiDroiK8qig27iiIOHXGUz2TpSrpHB5bWvTP6nELLN2m05dG5gF0EA8H3WjCfMrVPM11avgLt5TOOKNR8u5lQZvNVoUY_f1X_cUfhRyNuTpnJY6WOVQy-lbG0XQp4Wbuske3-6hlAW_JSFIsySWw\"       }   ] }")

}

func TestDuplicatesClient(t *testing.T) {
	executeCommand("terraform", "-chdir=./usecase2", "plan")
	err := executeCommand("terraform", "-chdir=./usecase2", "apply", "-auto-approve")
	if err == nil {
		t.Errorf("Should be error")
	}
}

func TestClientCustomUpdateValues(t *testing.T) {
	executeCommand("terraform", "-chdir=./usecase3", "plan")
	executeCommand("terraform", "-chdir=./usecase3", "apply", "-auto-approve")

	client, err := sdkClient.ReadClient("INTEGRATION_TEST_AUTO_CLIENT_2")
	if err != nil {
		t.Errorf("Client should exist %s", err)
	}

	assert.Equal(t, client.CoreOAuth2ClientConfig.Status.Value, "Active")
	assert.Equal(t, client.CoreOAuth2ClientConfig.Scopes.Value, []string{"profile", "openid"})
	assert.Equal(t, client.CoreOAuth2ClientConfig.RedirectionUris.Value, []string{"http://localhost:4200"})
	assert.Equal(t, client.CoreOAuth2ClientConfig.ClientType.Value, "Confidential")

	assert.Equal(t, client.AdvancedOAuth2ClientConfig.TokenEndpointAuthMethod.Value, "private_key_jwt")
	assert.Equal(t, client.AdvancedOAuth2ClientConfig.GrantTypes.Value, []string{"client_credentials"})

	assert.Equal(t, client.SignEncOAuth2ClientConfig.PublicKeyLocation.Value, "jwks")
	assert.Equal(t, client.SignEncOAuth2ClientConfig.JwkSet.Value, "{   \"keys\": [       {           \"kty\": \"RSA\",           \"e\": \"AQAB\",            \"use\": \"sig\",           \"kid\": \"23c50663-6902-4cdb-9ef9-a15f9bb10eeb\",          \"n\": \"xsczx0C6tMxxZquC-bsAfCd7HGx41c8aje5pn5Oys6lSUKEJ7mqZFWQi86IkOWwjoorSpczP1xOKhwN5_80_Yi_zAYs7iDeENDXt-O5bjNdagC3nxgGYoefSaJgmKmK3Da20b_YcIWGEddS_IK4QRtgLEcY3wh6-9fUvHsbCSarPGdm34E4F1jAaiuC1dTyT5qUiDroiK8qig27iiIOHXGUz2TpSrpHB5bWvTP6nELLN2m05dG5gF0EA8H3WjCfMrVPM11avgLt5TOOKNR8u5lQZvNVoUY_f1X_cUfhRyNuTpnJY6WOVQy-lbG0XQp4Wbuske3-6hlAW_JSFIsySWw\"       }   ] }")

}

func TestDestroyClients(t *testing.T) {
	executeCommand("terraform", "-chdir=./usecase4", "plan")
	err := executeCommand("terraform", "-chdir=./usecase4", "apply", "-auto-approve")
	if err != nil {
		t.Errorf("Should not be error %s", err)
	}

	must_not_exist_clients_names := []string{
		"INTEGRATION_TEST_AUTO_CLIENT_1",
		"INTEGRATION_TEST_AUTO_CLIENT_2",
		"INTEGRATION_TEST_AUTO_CLIENT_3"}

	for _, name := range must_not_exist_clients_names {
		_, err := sdkClient.ReadClient(name)
		if err != nil && strings.Contains(err.Error(), "404") {
			continue
		}
		t.Errorf("Client should not exist %s %s", name, err)
	}
}
