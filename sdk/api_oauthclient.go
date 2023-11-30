package sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (client *Client) ReadClient(clientId string) (*OAuth2Client, error) {

	if clientId == "" {
		return nil, errors.New("client not specified")
	}

	url := fmt.Sprintf("%s/realm-config/agents/OAuth2Client/%s", client.GetFullPath(), clientId)

	req, err := http.NewRequest("GET", url, nil)

	response, err := client.DoGenericRequest(req)
	if err != nil {
		return nil, err
	}

	var oAuth2Client OAuth2Client
	err = json.Unmarshal(response, &oAuth2Client)
	if err != nil {
		return nil, err
	}

	return &oAuth2Client, nil
}

func (client *Client) UpdateClient(clientId string, oAuth2Client OAuth2Client) (*OAuth2Client, error) {

	fullJson, err := oAuth2ClientToFullJSON(&oAuth2Client, client)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/realm-config/agents/OAuth2Client/%s", client.GetFullPath(), clientId)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(fullJson))
	req.Header.Add("Accept-API-Version", "protocol=2.0")

	response, err := client.DoGenericRequest(req)
	if err != nil {
		return nil, err
	}

	var oAuth2ClientConfig OAuth2Client
	err = json.Unmarshal(response, &oAuth2ClientConfig)
	if err != nil {
		return nil, err
	}

	return &oAuth2ClientConfig, nil
}

func (client *Client) DeleteClient(clientId string) error {
	url := fmt.Sprintf("%s/realm-config/agents/OAuth2Client/%s", client.GetFullPath(), clientId)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept-API-Version", "protocol=2.0")

	_, err = client.DoGenericRequest(req)
	if err != nil {
		return err
	}

	return nil
}

/*
This endpoint provides fields that are not wrapped in its responses.
Therefore, it is necessary to wrap all of its fields to match the format of the model in the SDK.
However, please note that we cannot use the model as a reference, as the template may include fields not anticipated in our SDK.
*/
func (client *Client) GetJSONClientTemplate() ([]byte, error) {

	url := fmt.Sprintf("%s/realm-config/agents/OAuth2Client?_action=template", client.GetFullPath())

	// Yes ... it's a post, forgerock is a bit strange
	req, err := http.NewRequest("POST", url, nil)
	req.Header.Add("X-Requested-With", "")

	response, err := client.DoGenericRequest(req)
	if err != nil {
		return nil, err
	}

	wrapepdResponse, err := wrapModelsFields(response)

	return wrapepdResponse, nil
}

func (client *Client) GetClientTemplate() (*OAuth2Client, error) {
	jsonTemplate, err := client.GetJSONClientTemplate()
	if err != nil {
		return nil, err
	}

	var template OAuth2Client
	err = json.Unmarshal(jsonTemplate, &template)
	if err != nil {
		return nil, err
	}

	return &template, nil
}