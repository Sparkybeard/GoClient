package payloads

type CreateUserPayload struct {
	UserName string `json:"name"`
}

type GetUserPayload struct {
	UserName string `json:"name"`
}

type UpdateUserPayload struct {
	UserName string `json:"name"`
}

type DeleteUserPayload struct {
	UserName string `json:"name"`
}
