package payloads

type CreateTeamPayload struct {
	TeamName   string `json:"name"`
	CostCenter string `json:"costCenterName"`
}

type GetTeamPayload struct {
	TeamName string `json:"name"`
}

type UpdateTeamPayload struct {
	TeamName   string `json:"name"`
	CostCenter string `json:"costCenterName"`
}

type DeleteTeamPayload struct {
	TeamName string `json:"name"`
}

type GetTeamTypesPayload struct {
}
