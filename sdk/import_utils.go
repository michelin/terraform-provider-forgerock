package sdk

import (
	"encoding/json"
	"fmt"

	jsonpatch "github.com/evanphx/json-patch"
)


func ExtractChangedData(reference *OAuth2Client, currentValue *OAuth2Client) (*OAuth2Client, error){
	refbyte, err := json.Marshal(reference)
	if(err != nil){
		return nil, err
	}

	curbyte, err := json.Marshal(currentValue)
	if(err != nil){
		return nil, err
	}


	patch, err := jsonpatch.CreateMergePatch(refbyte, curbyte)
	if err != nil {
		return nil, err
	}


	fmt.Println(string(patch))

	var newClient OAuth2Client
	err = json.Unmarshal(patch, &newClient)
	if err != nil {
		return nil, err
	}

	return &newClient, nil
}