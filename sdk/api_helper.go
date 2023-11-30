package sdk

import (
	"encoding/json"
	"fmt"

	jsonpatch "github.com/evanphx/json-patch"
)

/*
Converts the OAuth2Client struct to a complete JSON and merges any missing values with default values
from the ForgeRock template endpoint.
*/
func oAuth2ClientToFullJSON(oAuth2Client *OAuth2Client, client *Client) ([]byte, error) {
	JSONByteFromTemplate, err := client.GetJSONClientTemplate()
	if err != nil {
		return nil, err
	}

	JSONByteFromContext, err := json.Marshal(oAuth2Client)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(JSONByteFromContext))
	fmt.Println("===================")
	fmt.Println(string(JSONByteFromTemplate))
	fmt.Println("===================")

	// Priority for our jsonProvided, else, missing fields imported from templates
	mergedJson, err := jsonpatch.MergeMergePatches(JSONByteFromTemplate, JSONByteFromContext)
	if err != nil {
		return nil, err
	}

	return mergedJson, nil
}

func wrapModelsFields(unwrappedOauth2Client []byte) ([]byte, error) {
	var result map[string]interface{}
	err := json.Unmarshal(unwrappedOauth2Client, &result)
	if err != nil {
		return nil, err
	}

	for key, configObject := range result {
		configObjectMap := configObject.(map[string]interface{})

		for fieldName, fieldValue := range configObjectMap {
			// Patch for overrideOAuth2ClientConfig that is not wrapped ...
			if key == "overrideOAuth2ClientConfig" {
				configObjectMap[fieldName] = fieldValue
			} else {
				if fieldValue == nil {
					configObjectMap[fieldName] = WrappedValueNotTyped{Inherited: false, Value: nil}
					continue
				}

				// Forgerock template API return array with empty string instead of empty array
				if arr, ok := fieldValue.([]interface{}); ok && len(arr) == 1 && arr[0] == "" {
					configObjectMap[fieldName] = WrappedValueNotTyped{Inherited: false, Value: []string{}}
					continue
				}

				configObjectMap[fieldName] = WrappedValueNotTyped{Inherited: false, Value: fieldValue}
			}
		}

	}

	wrappedData, err := json.Marshal(result)

	return wrappedData, nil
}
