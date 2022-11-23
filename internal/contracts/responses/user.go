package responses

type CreateUserResponse struct {
	Data struct {
		Id string `json:"id"`
		UserName string `json:"name"`
	} `json:"data"`	
}

type GetUserResponse struct {
	Data struct {
		Id string `json:"id"`
		UserName string `json:"name"`
	} `json:"data"`
}

type UpdateUserResponse struct {
	Data struct {
		Id string `json:"id"`
		UserName string `json:"name"`
	} `json:"data"`
}

type DeleteUserResponse struct {
	Data struct {
		Id string `json:"id"`
		UserName string `json:"name"`
	} `json:"data"`
}
