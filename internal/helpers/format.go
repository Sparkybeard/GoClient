package helpers

func FormatDatabaseName(applicationName string, solutionName string, environmentName string) string {
	return applicationName + "-" + solutionName + "_" + environmentName
}

func FormatGuid(guid string) string {
	if guid == "" {
		guid = "00000000-0000-0000-0000-000000000000"
	}

	return guid
}
