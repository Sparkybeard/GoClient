package payloads

type CreateUserPayload struct {
	UserName string `json:"name"`
}

type GetUserPayload struct {
	UserName string `json:"name"`
	UserId	 string	`json:"id"`
}

type UpdateUserPayload struct {
	UserName string `json:"name"`
	UserId	 string	`json:"id"`
}

type DeleteUserPayload struct {
	UserName string `json:"name"`
}
