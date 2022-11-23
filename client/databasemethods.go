package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	helperfmt "github.com/Sparkybeard/GoClient/internal/helpers"
	"github.com/Sparkybeard/GoClient/internal/consts"
	"github.com/Sparkybeard/GoClient/internal/contracts/responses"
	"github.com/Sparkybeard/GoClient/internal/contracts/payloads"
)

func (c *Client) InstanciateDb(ctx context.Context, applicationName string, solutionName string, environmentName string, dbSku string) (responses.InstanciateDbResponse, error) {
	var result responses.InstanciateDbResponse

	// create payload
	reqPayload := payload[payloads.InstanciateDbPayload]{metadata{}, payloads.InstanciateDbPayload{}}
	reqPayload.OperationPayload.ApplicationName = applicationName
	reqPayload.OperationPayload.SolutionName = solutionName
	reqPayload.OperationPayload.Sku = dbSku
	reqPayload.OperationPayload.EnvironmentName = environmentName

	// execute request
	resp, err := doAPIRequest(reqPayload, c, consts.InstanciateDbActionPath)
	if err != nil {
		return result, fmt.Errorf("failed to instanciate database %+v: %w", applicationName+"-"+solutionName, err)
	}

	// read output
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("error parsing body. \n%w", err)
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, fmt.Errorf("error unmarshaling data from request. \n%w", err)
	}

	// validate response
	if result.Data.ApplicationConnectionString != "" && result.Data.OpsConnectionString != "" {
		return result, nil
	}

	return result, fmt.Errorf("failed to instanciate database %+v: %w", helperfmt.FormatDatabaseName(applicationName, solutionName, environmentName), err)
}

func (c *Client) GetDb(ctx context.Context, applicationName string, solutionName string, environmentName string, serverName string) (responses.GetDbResponse, error) {
	var result responses.GetDbResponse

	// create payload
	reqPayload := payload[payloads.GetDbPayload]{metadata{}, payloads.GetDbPayload{}}
	reqPayload.OperationPayload.ApplicationName = applicationName
	reqPayload.OperationPayload.SolutionName = solutionName
	reqPayload.OperationPayload.EnvironmentName = environmentName
	reqPayload.OperationPayload.ServerName = serverName

	// execute request
	resp, err := doAPIRequest(reqPayload, c, consts.GetDbActionPath)
	if err != nil {
		return result, fmt.Errorf("failed to retrieve database data. \n%w", err)
	}

	// read output
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("error parsing body \n%w", err)
	}
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return result, fmt.Errorf("failed to unmarshall payload \n%w", err)
	}

	// validate response
	// if incorrect return empty string to instanciate
	if result.Data.DatabaseName == helperfmt.FormatDatabaseName(applicationName, solutionName, environmentName) {
		return result, nil
	}

	return result, nil
}

func (c *Client) DeleteDb(ctx context.Context, applicationName string, solutionName string, environmentName string, serverName string) (responses.DeleteDbResponse, error) {
	var result responses.DeleteDbResponse

	// create payload
	reqPayload := payload[payloads.DeleteDbPayload]{metadata{}, payloads.DeleteDbPayload{}}
	reqPayload.OperationPayload.ApplicationName = applicationName
	reqPayload.OperationPayload.SolutionName = solutionName
	reqPayload.OperationPayload.EnvironmentName = environmentName
	reqPayload.OperationPayload.ServerName = serverName

	// execute request
	resp, err := doAPIRequest(reqPayload, c, consts.DeleteDbActionPath)
	if err != nil {
		return result, fmt.Errorf("failed to retrieve database data. \n%w", err)
	}

	// read output
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("error parsing body \n%w", err)
	}
	err = json.Unmarshal([]byte(body), &result)

	// validate response
	if err != nil {
		return result, fmt.Errorf("wrong data retrieved. \n%w", err)
	}

	return result, nil
}

func (c *Client) UpdateDb() (responses.UpdateDbResponse, error) {
	var result responses.UpdateDbResponse
	return result, nil
}