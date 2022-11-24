package responses

import "time"

type CreateApplicationResponse struct {
	Data struct {
		ApplicationId	string `json:"applicationId"`
		ApplicationName string `json:"applicationName"`
		EndOfLife time.Time `json:"endOfLife"`
		Monitor string `json:"monitor"`
		Interval string `json:"interval"`
		PartOf string `json:"partOf"`
		SolutionId string `json:"solutionId"`
		SolutionName string `json:"solutionName"`
		SolutionArdId int64 `json:"solutionArdId"`
	} `json:"data"`
}

type GetApplicationResponse struct {
	Data struct {
		ApplicationId	string `json:"applicationId"`
		ApplicationName string `json:"applicationName"`
		EndOfLife time.Time `json:"endOfLife"`
		Monitor string `json:"monitor"`
		Interval string `json:"interval"`
		PartOf string `json:"partOf"`
		Workload string `json:"workload"`
		SolutionId string `json:"solutionId"`
		SolutionName string `json:"solutionName"`
		SolutionArdId int64 `json:"solutionArdId"`
	} `json:"data"`
}

type UpdateApplicationResponse struct {
	Data struct {
		ApplicationId	string `json:"applicationId"`
		ApplicationName string `json:"applicationName"`
		EndOfLife time.Time `json:"endOfLife"`
		Monitor string `json:"monitor"`
		Interval string `json:"interval"`
		PartOf string `json:"partOf"`
		Workload string `json:"workload"`
		SolutionId string `json:"solutionId"`
		SolutionName string `json:"solutionName"`
		SolutionArdId int64 `json:"solutionArdId"`
	} `json:"data"`
}

type DeleteApplicationResponse struct {
	Data struct {
	} `json:"data"`
}
