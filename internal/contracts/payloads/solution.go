package payloads

type CreateSolutionPayload struct {
	SolutionName  string `json:"name"`
	SolutionId    string `json:"solutionId"`
	SolutionArdId int64  `json:"ardId"`
}

type GetSolutionPayload struct {
	SolutionArdId int64 `json:"solutionArdId"`
}

type UpdateSolutionPayload struct {
	SolutionName  string `json:"name"`
	SolutionId    string `json:"solutionId"`
	SolutionArdId int64 `json:"ardId"`
}

type DeleteSolutionPayload struct {
	SolutionName string `json:"solutionName"`
}
