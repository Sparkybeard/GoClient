package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Sparkybeard/GoClient/internal/consts"
	"github.com/Sparkybeard/GoClient/internal/contracts/payloads"
	"github.com/Sparkybeard/GoClient/internal/contracts/responses"
	"github.com/Sparkybeard/GoClient/internal/helpers"
)

func (c *Client) CreateUser(ctx context.Context, userName string) (responses.CreateUserResponse, error) {
	var result responses.CreateUserResponse

	// create payload
	reqPayload := payload[payloads.CreateUserPayload]{metadata{}, payloads.CreateUserPayload{}}
	reqPayload.OperationPayload.UserName = userName

	// execute request
	resp, err := doAPIRequest(reqPayload, c, consts.CreateUserActionPath)
	if err != nil {
		return result, fmt.Errorf("failed to create application %+v: %w", userName, err)
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
	if result.Data.Id != "" && result.Data.UserName == userName {
		return result, nil
	}

	return result, fmt.Errorf("data received from api doesn't match data from request: %s", err.Error())
	// TODO: add errors to add details to validation failure
}

func (c *Client) GetUser(ctx context.Context, userName string, userId string) (responses.GetUserResponse, error) {
	var result responses.GetUserResponse

	// create payload
	reqPayload := payload[payloads.GetUserPayload]{metadata{}, payloads.GetUserPayload{}}
	reqPayload.OperationPayload.UserName = userName
	reqPayload.OperationPayload.UserId = helpers.FormatGuid(userId)

	// execute request
	resp, err := doAPIRequest(reqPayload, c, consts.GetUserActionPath)
	if err != nil {
		return result, fmt.Errorf("failed to retrieve user data. \n%w", err)
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
	if result.Data.UserName == userName && result.Data.Id != "" {
		return result, nil
	}

	return result, fmt.Errorf("wrong data received from api")
}

func (c *Client) DeleteUser(ctx context.Context, userName string) (bool, error) {
	var result responses.DeleteUserResponse
	// TODO: incorporate result in function return

	// create payload
	reqPayload := payload[payloads.DeleteUserPayload]{metadata{}, payloads.DeleteUserPayload{}}
	reqPayload.OperationPayload.UserName = userName

	// execute request
	resp, err := doAPIRequest(reqPayload, c, consts.DeleteUserActionPath)
	if err != nil {
		return false, fmt.Errorf("failed to delete user. \n%w", err)
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

func (c *Client) UpdateUser() (responses.UpdateUserResponse, error) {
	var result responses.UpdateUserResponse

	return result, nil
}
