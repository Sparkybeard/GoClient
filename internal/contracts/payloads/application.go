package payloads

type CreateApplicationPayload struct {
	ApplicationName string `json:"name"`
	EndOfLife	string `json:"endOfLife"`
	Workload string `json:"workload"`
	PartOf string `json:"partOf"`
	Monitor	string `json:"monitor"`
	SolutionArdId int64 `json:"solutionArdId"`
}

type GetApplicationPayload struct {
	ApplicationId 	string `json:"id"`
	ApplicationName string `json:"name"`
	SolutionArdId int64 `json:"solutionArdId"`
	SolutionId	string `json:"solutionId"`
}

type UpdateApplicationPayload struct {
	ApplicationName string `json:"name"`
	EndOfLife	string `json:"endOfLife"`
	Workload string `json:"workload"`
	PartOf string `json:"partOf"`
	Monitor	string `json:"monitor"`
	SolutionArdId int64 `json:"solutionArdId"`
}

type DeleteApplicationPayload struct {
	ApplicationName string `json:"name"`
	SolutionArdId int64 `json:"solutionArdId"`
}
