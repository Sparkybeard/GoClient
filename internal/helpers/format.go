package helpers

func FormatDatabaseName(applicationName string, solutionName string, environmentName string) string {
	return applicationName+"-"+solutionName+"_"+environmentName
}
