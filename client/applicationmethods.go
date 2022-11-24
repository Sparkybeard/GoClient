package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	//helperfmt "wmclient/internal/helpers"
	"github.com/Sparkybeard/GoClient/internal/consts"
	"github.com/Sparkybeard/GoClient/internal/contracts/payloads"
	"github.com/Sparkybeard/GoClient/internal/contracts/responses"
)

func (c *Client) CreateApplication(ctx context.Context, applicationName string, ardId int64, endOfLife time.Time, monitor string, partOf string, workload string) (responses.CreateApplicationResponse, error) {
	// TODO explore get solutions with ardId instead of other solution parameters
	var result responses.CreateApplicationResponse

	// create payload
	reqPayload := payload[payloads.CreateApplicationPayload]{metadata{}, payloads.CreateApplicationPayload{}}
	reqPayload.OperationPayload.ApplicationName = applicationName
	reqPayload.OperationPayload.SolutionArdId = ardId
	reqPayload.OperationPayload.EndOfLife = endOfLife.UTC()
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

	return result, fmt.Errorf("failed to create application: %s", err.Error())
}

func (c *Client) GetApplication(ctx context.Context, applicationId string, applicationName string, 
	solutionArdId int64, solutionId string) (responses.GetApplicationResponse, error) {
	var result responses.GetApplicationResponse

	// create payload
	reqPayload := payload[payloads.GetApplicationPayload]{metadata{}, payloads.GetApplicationPayload{}}
	reqPayload.OperationPayload.ApplicationName = applicationName
	reqPayload.OperationPayload.SolutionArdId = solutionArdId
	reqPayload.OperationPayload.SolutionId = solutionId
	reqPayload.OperationPayload.ApplicationId = applicationId

	// execute request
	resp, err := doAPIRequest(reqPayload, c, consts.GetApplicationActionPath)
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

func (c *Client) DeleteApplication(ctx context.Context, applicationName string, ardId int64) (bool, error) {
	var result responses.DeleteApplicationResponse
	// TODO: incorporate result in function return

	// create payload
	reqPayload := payload[payloads.DeleteApplicationPayload]{metadata{}, payloads.DeleteApplicationPayload{}}
	reqPayload.OperationPayload.ApplicationName = applicationName
	reqPayload.OperationPayload.SolutionArdId = ardId

	// execute request
	resp, err := doAPIRequest(reqPayload, c, consts.DeleteApplicationActionPath)
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
