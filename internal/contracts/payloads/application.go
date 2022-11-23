package payloads

type CreateApplicationPayload struct {
	ApplicationName string `json:"name"`
	EndOfLife	string `json:"endOfLife"`
	Workload string `json:"workload"`
	PartOf string `json:"partOf"`
	Monitor	string `json:"monitor"`
	SolutionId string `json:"solutionId"`
}

type GetApplicationPayload struct {
	ApplicationName string `json:"name"`
	SolutionId string `json:"solutionId"`
}

type UpdateApplicationPayload struct {
	ApplicationName string `json:"name"`
	EndOfLife	string `json:"endOfLife"`
	Workload string `json:"workload"`
	PartOf string `json:"partOf"`
	Monitor	string `json:"monitor"`
	SolutionId string `json:"solutionId"`
}

type DeleteApplicationPayload struct {
	ApplicationName string `json:"name"`
	SolutionId string `json:"solutionId"`
}
