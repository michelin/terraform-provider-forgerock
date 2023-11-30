package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (client *Client) authenticate() error {

	url := fmt.Sprintf("%s/%s/authenticate" ,client.ApiUrl, "realms/root")

	req, err := http.NewRequest("POST", url, nil)

	req.Header.Add("Accept-API-Version", "resource=2.1")
	req.Header.Add("X-OpenAM-Username", client.Username)
	req.Header.Add("X-OpenAM-Password", client.Password)

	response, err := client.DoGenericRequest(req)
	if err != nil {
		return err
	}

	var authResponse AuthResponse
	err = json.Unmarshal(response, &authResponse)
	if err != nil {
		return err
	}

	client.Token = authResponse.TokenId
	return nil
}
