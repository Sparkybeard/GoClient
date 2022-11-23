package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	//helperfmt "wmclient/internal/helpers"
	"github.com/Sparkybeard/GoClient/internal/consts"
	"github.com/Sparkybeard/GoClient/internal/contracts/responses"
	"github.com/Sparkybeard/GoClient/internal/contracts/payloads"
)

func (c *Client) CreateApplication(ctx context.Context, applicationName string, solutionId string, endOfLife string, monitor string, partOf string, workload string) (responses.CreateApplicationResponse, error) {
	// TODO explore get solutions with ardId instead of other solution parameters
	var result responses.CreateApplicationResponse

	// create payload
	reqPayload := payload[payloads.CreateApplicationPayload]{metadata{}, payloads.CreateApplicationPayload{}}
	reqPayload.OperationPayload.ApplicationName = applicationName
	reqPayload.OperationPayload.SolutionId = solutionId
	reqPayload.OperationPayload.EndOfLife = endOfLife
	reqPayload.OperationPayload.Monitor = monitor 
	reqPayload.OperationPayload.PartOf = partOf
	reqPayload.OperationPayload.Workload = workload

	// execute request
	resp, err := doAPIRequest(reqPayload, c, consts.CreateApplicationActionPath)
	if err != nil {
		return result, fmt.Errorf("failed to create application %+v: %w", applicationName, err)
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
	if result.Data.ApplicationName != "" && result.Data.ApplicationId != "" {
		return result, nil
	}

	return result, fmt.Errorf("failed to create application %+v: %w", err)
}

func (c *Client) GetApplication(ctx context.Context, applicationName string, solutionName string) (responses.GetApplicationResponse, error) {
	var result responses.GetApplicationResponse

	// create payload
	reqPayload := payload[payloads.GetDbPayload]{metadata{}, payloads.GetDbPayload{}}
	reqPayload.OperationPayload.ApplicationName = applicationName
	reqPayload.OperationPayload.SolutionName = solutionName

	// execute request
	resp, err := doAPIRequest(reqPayload, c, consts.GetDbActionPath)
	if err != nil {
		return result, fmt.Errorf("failed to retrieve application data. \n%w", err)
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
	if result.Data.ApplicationName == applicationName {
		return result, nil
	}

	return result, fmt.Errorf("application name from requested does not match")
}

func (c *Client) DeleteApplication(ctx context.Context, applicationName string, solutionName string) (bool, error) {
	var result responses.DeleteApplicationResponse
	// TODO: incorporate result in function return

	// create payload
	reqPayload := payload[payloads.DeleteApplicationPayload]{metadata{}, payloads.DeleteApplicationPayload{}}
	reqPayload.OperationPayload.ApplicationName = applicationName
	reqPayload.OperationPayload.SolutionId = solutionName

	// execute request
	resp, err := doAPIRequest(reqPayload, c, consts.DeleteDbActionPath)
	if err != nil {
		return false, fmt.Errorf("failed to delete application. \n%w", err)
	}

	// read output
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("error parsing body \n%w", err)
	}
	err = json.Unmarshal([]byte(body), &result)

	// validate response
	if err != nil {
		return false, fmt.Errorf("wrong data retrieved. \n%w", err)
	}

	return true, nil
}

func (c *Client) UpdateApplication() (responses.UpdateApplicationResponse, error) {
	var result responses.UpdateApplicationResponse
	return result, nil	
}