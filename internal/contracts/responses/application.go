package responses

type CreateApplicationResponse struct {
	Data struct {
		ApplicationId	string `json:"applicationId"`
		ApplicationName string `json:"applicationName"`
		EndOfLife string `json:"endOfLife"`
		Monitor string `json:"monitor"`
		Interval string `json:"interval"`
		PartOf string `json:"partOf"`
		SolutionId string `json:"solutionId"`
	} `json:"data"`
}

type GetApplicationResponse struct {
	Data struct {
		ApplicationId	string `json:"applicationId"`
		ApplicationName string `json:"applicationName"`
		EndOfLife string `json:"endOfLife"`
		Monitor string `json:"monitor"`
		Interval string `json:"interval"`
		PartOf string `json:"partOf"`
		SolutionId string `json:"solutionId"`
	} `json:"data"`
}

type UpdateApplicationResponse struct {
	Data struct {
		ApplicationId	string `json:"applicationId"`
		ApplicationName string `json:"applicationName"`
		EndOfLife string `json:"endOfLife"`
		Monitor string `json:"monitor"`
		Interval string `json:"interval"`
		PartOf string `json:"partOf"`
		SolutionId string `json:"solutionId"`
	} `json:"data"`
}

type DeleteApplicationResponse struct {
	Data struct {
	} `json:"data"`
}
