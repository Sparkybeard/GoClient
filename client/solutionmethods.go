package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	//helperfmt "wmclient/internal/helpers"
	"wmclient/internal/consts"
	"wmclient/internal/contracts/responses"
	"wmclient/internal/contracts/payloads"
)

func (c *Client) CreateSolution(ctx context.Context, solutionId string, solutionName string) (responses.CreateSolutionResponse, error) {
	var result responses.CreateSolutionResponse

	// create payload
	reqPayload := payload[payloads.CreateSolutionPayload]{metadata{}, payloads.CreateSolutionPayload{}}
	reqPayload.OperationPayload.SolutionName = solutionName
	reqPayload.OperationPayload.SolutionId = solutionId

	// execute request
	resp, err := doAPIRequest(reqPayload, c, consts.CreateApplicationActionPath)
	if err != nil {
		return result, fmt.Errorf("failed to create application %+v: %w", solutionName, err)
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
	if result.Data.SolutionName != "" && result.Data.SolutionId != "" {
		return result, nil
	}

	return result, fmt.Errorf("failed to create solution %+v: %w", err)
}

func (c *Client) GetSolution(ctx context.Context, solutionName string) (responses.GetSolutionResponse, error) {
	var result responses.GetSolutionResponse

	// create payload
	reqPayload := payload[payloads.GetSolutionPayload]{metadata{}, payloads.GetSolutionPayload{}}
	reqPayload.OperationPayload.SolutionName = solutionName

	// execute request
	resp, err := doAPIRequest(reqPayload, c, consts.GetDbActionPath)
	if err != nil {
		return result, fmt.Errorf("failed to retrieve solution data. \n%w", err)
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
	if result.Data.SolutionName == solutionName {
		return result, nil
	}

	return result, fmt.Errorf("solution name from requested does not match")
}

func (c *Client) DeleteSolution(ctx context.Context, solutionName string) (bool, error) {
	var result responses.DeleteSolutionResponse
	// TODO: incorporate result in function return

	// create payload
	reqPayload := payload[payloads.DeleteSolutionPayload]{metadata{}, payloads.DeleteSolutionPayload{}}
	reqPayload.OperationPayload.SolutionName = solutionName

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

func (c *Client) UpdateSolution() (responses.UpdateSolutionResponse, error) {
	var result responses.UpdateSolutionResponse
	return result, nil
}