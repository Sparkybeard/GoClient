package payloads

type CreateSolutionPayload struct {
	SolutionName  string `json:"name"`
	SolutionId    string `json:"solutionId"`
	SolutionArdId int32  `json:"ardId"`
}

type GetSolutionPayload struct {
	SolutionName string `json:"solutionName"`
}

type UpdateSolutionPayload struct {
	SolutionName  string `json:"name"`
	SolutionId    string `json:"solutionId"`
	SolutionArdId string `json:"ardId"`
}

type DeleteSolutionPayload struct {
	SolutionName string `json:"solutionName"`
}
