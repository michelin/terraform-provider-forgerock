package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	Username   string
	Password   string
	ApiUrl     string
	RealmPath  string
	Token      string
	HttpClient *http.Client
}


func NewClient(apiUrl string, realmPath string, username string, password string) (*Client, error) {

	client := &Client{
		Username: username,
		Password: password,
		ApiUrl:   apiUrl,
		RealmPath:   realmPath,
	}

	client.HttpClient = &http.Client{Timeout: 10 * time.Second}

	err := client.authenticate()

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (client *Client) GetFullPath() string {
	return fmt.Sprintf("%s%s", client.ApiUrl, client.RealmPath)
}

func (client *Client) DoGenericRequest(req *http.Request) ([]byte, error) {

	req.Header.Set("Content-Type", "application/json")
	if client.Token != "" {
		req.Header.Add("Cookie", "amlbcookie=01; devmcookie="+client.Token)
	}

	res, err := client.HttpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		var prettyJSON bytes.Buffer

		if req.Body != nil {
			bodyRequestReader, err := req.GetBody()
			if err != nil {
				return nil, err
			}
			bodyRequest, err := io.ReadAll(bodyRequestReader)
			if err != nil {
				return nil, err
			}

			_ = json.Indent(&prettyJSON, bodyRequest, "", "\t")

		}

		prettyJSON.Reset()

		_ = json.Indent(&prettyJSON, body, "", "\t")

		switch res.StatusCode {
		case 404:
			return nil, fmt.Errorf("Http Code: %d, object not found in FORGEROCK", res.StatusCode)
		case 500:
			return nil, fmt.Errorf("Http Code: %d, Internal server ERROR", res.StatusCode)
		case 400:
			return nil, fmt.Errorf("Http Code: %d, Bad Request", res.StatusCode)
		case 401:
			return nil, fmt.Errorf("Http Code: %d, Auth KO, please check your token", res.StatusCode)
		case 403:
			return nil, fmt.Errorf("Http Code: %d, Forbidden, please check permissions", res.StatusCode)
		case 409:
			return nil, fmt.Errorf("Http Code: %d, object already exists, please check.", res.StatusCode)
		default:
			if res.StatusCode < 200 && res.StatusCode >= 300 {
				return nil, fmt.Errorf("Http Code: %d, Unknown Error", res.StatusCode)
			}
		}
	}

	return body, err
}
