package responses

type CreateSolutionResponse struct {
	Data struct {
		Id string `json:"id"`
		SolutionId string `json:"solutionId"`
		SolutionName string `json:"solutionName"`
	} `json:"data"`
}

type GetSolutionResponse struct {
	Data struct {
		Id string `json:"id"`
		SolutionId string `json:"solutionId"`
		SolutionName string `json:"solutionName"`
		SolutionArdId int32 `json:"ardId"`
	} `json:"data"`
}

type UpdateSolutionResponse struct {
	Data struct {
		Id string `json:"id"`
		SolutionId string `json:"solutionId"`
		SolutionName string `json:"solutionName"`
	} `json:"data"`
}

type DeleteSolutionResponse struct {
	Data struct {
	} `json:"data"`
}