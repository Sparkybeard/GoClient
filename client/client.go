package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	oauthclient "github.com/Sparkybeard/GoClient/internal/authentication"
	data "github.com/Sparkybeard/GoClient/internal/contracts/payloads"
	"github.com/Sparkybeard/GoClient/internal/interfaces"
)

type Options struct {
	ApiUrl  string
	Verbose bool
}

type Client struct {
	httpClient *http.Client
	options    *Options
}

type metadata struct {
	MessageId                   string `json:"messageId"`
	MessageCreatorId            string `json:"messageCreatorId"`
	MessageCreationTimestampUTC string `json:"messageCreationTimestampUTC"`
	CorrelationId               string `json:"correlationId"`
	DistributedTracingContext   string `json:"distributedTracingContext"`
	AuthenticationValue         string `json:"authenticationValue"`
	ValidateOnly                bool   `json:"validateOnly"`
	LanguageCode                string `json:"languageCode"`
}

type operations interface {
	data.InstanciateDbPayload | data.GetDbPayload | data.DeleteDbPayload |
		data.CreateApplicationPayload | data.GetApplicationPayload | data.DeleteApplicationPayload | data.UpdateApplicationPayload |
		data.CreateSolutionPayload | data.GetSolutionPayload | data.DeleteSolutionPayload | data.UpdateSolutionPayload |
		data.CreateTeamPayload | data.GetTeamPayload | data.DeleteTeamPayload | data.UpdateTeamPayload |
		data.CreateUserPayload | data.GetUserPayload | data.DeleteUserPayload | data.UpdateUserPayload
}

type payload[T operations] struct {
	Metadata         metadata `json:"metadata"`
	OperationPayload T        `json:"data"`
}

func NewClient(ops Options) (interfaces.WMclient, error) {

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	httpclient := http.Client{
		Transport: tr,
	}

	var client interfaces.WMclient = &Client{
		httpClient: &httpclient,
		options:    &ops,
	}

	return client, nil
}

func doAPIRequest[T operations](payload payload[T], wmclient *Client, actionUrl string) (*http.Response, error) {
	// get oauth token
	token, err := oauthclient.GetAuthToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get authentication token")
	}

	// configure metadata
	payload.Metadata.AuthenticationValue = token
	payload.Metadata.MessageCreationTimestampUTC = time.Now().UTC().Format("2006-01-02T15:04:05.0000000Z")
	payload.Metadata.ValidateOnly = false

	// parse payload to json

	json_data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to parse payload, %w", err)
	}

	er := bytes.NewBuffer(json_data)
	fmt.Printf("\n" + string(json_data))
	// create request
	req, err := http.NewRequest(http.MethodPost, wmclient.options.ApiUrl+actionUrl, er)
	if err != nil {
		return nil, fmt.Errorf("failed to create request, %w", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")

	// execute request
	resp, err := wmclient.httpClient.Do(req)
	if err != nil {
		return resp, fmt.Errorf("failed to retrieve database data. %w", err)
	}

	return resp, nil
}
