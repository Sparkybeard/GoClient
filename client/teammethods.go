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

func (c *Client) CreateTeam(ctx context.Context, costCenter string, teamName string) (responses.CreateTeamResponse, error) {
	var result responses.CreateTeamResponse

	// create payload
	reqPayload := payload[payloads.CreateTeamPayload]{metadata{}, payloads.CreateTeamPayload{}}
	reqPayload.OperationPayload.TeamName = teamName 
	reqPayload.OperationPayload.CostCenter = costCenter 

	// execute request
	resp, err := doAPIRequest(reqPayload, c, consts.CreateTeamActionPath)
	if err != nil {
		return result, fmt.Errorf("failed to create application %+v: %w", teamName, err)
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
	if result.Data.CostCenter != "" && result.Data.TeamName != "" {
		return result, nil
	}

	return result, fmt.Errorf("failed to create team: %s", err.Error())
	// TODO: add errors to add details to validation failure
}

func (c *Client) GetTeam(ctx context.Context, teamName string) (responses.GetTeamResponse, error) {
	var result responses.GetTeamResponse

	// create payload
	reqPayload := payload[payloads.GetTeamPayload]{metadata{}, payloads.GetTeamPayload{}}
	reqPayload.OperationPayload.TeamName = teamName 

	// execute request
	resp, err := doAPIRequest(reqPayload, c, consts.GetTeamActionPath)
	if err != nil {
		return result, fmt.Errorf("failed to retrieve team data. \n%w", err)
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
	if (result.Data.TeamName == teamName && result.Data.CostCenter != "") {
		return result, nil
	}

	return result, fmt.Errorf("wrong data received from api")
}

func (c *Client) DeleteTeam(ctx context.Context, teamName string) (bool, error) {
	var result responses.DeleteTeamResponse
	// TODO: incorporate result in function return

	// create payload
	reqPayload := payload[payloads.DeleteTeamPayload]{metadata{}, payloads.DeleteTeamPayload{}}
	reqPayload.OperationPayload.TeamName = teamName

	// execute request
	resp, err := doAPIRequest(reqPayload, c, consts.DeleteTeamActionPath)
	if err != nil {
		return false, fmt.Errorf("failed to delete team. \n%w", err)
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

func (c *Client) UpdateTeam() (responses.UpdateTeamResponse, error) {
	var result responses.UpdateTeamResponse

	return result, nil
}