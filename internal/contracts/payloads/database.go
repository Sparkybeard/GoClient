package payloads

type InstanciateDbPayload struct {
	ApplicationName string `json:"application"`
	SolutionName    string `json:"solution"`
	EnvironmentName string `json:"environment"`
	Sku             string `json:"sku"`
	ServerName      string `json:"serverName"`
}

type GetDbPayload struct {
	ApplicationName string `json:"application"`
	SolutionName    string `json:"solution"`
	EnvironmentName string `json:"environment"`
	ServerName      string `json:"serverName"`
}

type DeleteDbPayload struct {
	ApplicationName string `json:"application"`
	SolutionName    string `json:"solution"`
	EnvironmentName string `json:"environment"`
	ServerName      string `json:"serverName"`
}

type UpdateDbPayload struct {
}