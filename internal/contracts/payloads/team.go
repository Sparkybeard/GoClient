package payloads

type CreateTeamPayload struct {
	TeamName   string `json:"name"`
	CostCenter string `json:"costCenter"`
}

type GetTeamPayload struct {
	TeamName string `json:"name"`
}

type UpdateTeamPayload struct {
	TeamName   string `json:"name"`
	CostCenter string `json:"costCenter"`
}

type DeleteTeamPayload struct {
	TeamName string `json:"name"`
}
