package responses

type InstanciateDbResponse struct {
	Data struct {
		ApplicationConnectionString	string `json:"appConnectionString"`
		OpsConnectionString 		string `json:"opsConnectionString"`
		ServerName 					string `json:"serverName"`
	} `json:"data"`
}

type GetDbResponse struct {
	Data struct {
		DatabaseName	string	`json:"databaseName"`
		AppUserName		string	`json:"appUserName"`
		OpsUserName 	string	`json:"opsUserName"`
		DbState 		int32	`json:"dbState"`
	} `json:"data"`
}

type DeleteDbResponse struct {
	Data struct {
		Success bool
	} `json:"data"`
}

type UpdateDbResponse struct {
}