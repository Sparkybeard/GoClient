package responses

type CreateTeamResponse struct {
	Data struct {
		Id         string `json:"id"`
		TeamName   string `json:"name"`
		CostCenter string `json:"costCenter"`
	} `json:"data"`
}

type GetTeamResponse struct {
	Data struct {
		Id         string `json:"id"`
		TeamName   string `json:"name"`
		CostCenter string `json:"costCenterName"`
	} `json:"data"`
}

type UpdateTeamResponse struct {
	Data struct {
		Id         string `json:"id"`
		TeamName   string `json:"teamName"`
		CostCenter string `json:"costCenterName"`
	} `json:"data"`
}

type DeleteTeamResponse struct {
	Data struct {
	} `json:"data"`
}
