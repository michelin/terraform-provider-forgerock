package sdk

type AuthResponse struct {
	TokenId    string `json:"tokenId"`
	SuccessUrl string `json:"successUrl"`
	Realm      string `json:"realm"`
}
