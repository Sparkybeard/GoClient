package authentication

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	azOauthURL   = "https://login.microsoftonline.com/81381255-50e7-4e13-8b38-ec769644c349/oauth2/v2.0/token"
	grantType    = "client_credentials"
	scope        = "api://bcp.dev.ccoe-workload-manager/.default"
	clientSecret = "3LN8Q~CHjK9VLk10anxX3bz1khHVhLPb2NFK2bKE"
	clientId     = "b22eaea6-508f-46aa-b367-d34ff0044e31"
)

type Result struct {
	TokenType    string `json:"token_type"`
	ExpiresInd   int    `json:"expires_in"`
	ExtExpiresId int    `json:"ext_expires_in"`
	AccessToken  string `json:"access_token"`
}

func GetAuthToken() (string, error) {
	var result Result
	payload := url.Values{}
	payload.Add("grant_type", grantType)
	payload.Add("client_id", clientId)
	payload.Add("scope", scope)
	payload.Add("client_secret", clientSecret)

	resp, err := http.Post(azOauthURL, "application/x-www-form-urlencoded", strings.NewReader(payload.Encode()))
	if err != nil {
		return "failed to get credentials", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &result)
	return result.AccessToken, err
}
