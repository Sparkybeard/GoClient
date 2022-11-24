package interfaces

import (
	"context"

	"github.com/Sparkybeard/GoClient/internal/contracts/responses"
)

type WMclient interface {
	GetDb(ctx context.Context, applicationName string, solutionName string, 
		environmentName string, serverName string) (responses.GetDbResponse, error)
	InstanciateDb(ctx context.Context, applicationName string, solutionName string, 
		environmentName string, dbSku string) (responses.InstanciateDbResponse, error)
	DeleteDb(ctx context.Context, applicationName string, solutionName string, 
		environmentName string, serverName string) (responses.DeleteDbResponse, error)
	UpdateDb() (responses.UpdateDbResponse, error)

	GetSolution(ctx context.Context, ardId int64) (responses.GetSolutionResponse, error)
	CreateSolution(ctx context.Context, solutionId string, solutionName string, ardId int64) (responses.CreateSolutionResponse, error)
	DeleteSolution(ctx context.Context, solutionName string) (bool, error)
	UpdateSolution() (responses.UpdateSolutionResponse, error)

	GetApplication(ctx context.Context, applicationId string, applicationName string, 
		solutionArdId int64, solutionId string) (responses.GetApplicationResponse, error)
	CreateApplication(ctx context.Context, applicationName string, ardId int64, endOfLife string, monitor string, partOf string, workload string) (responses.CreateApplicationResponse, error)
	DeleteApplication(ctx context.Context, applicationName string, ardId int64) (bool, error)
	UpdateApplication() (responses.UpdateApplicationResponse, error)

	GetUser(ctx context.Context, userName string) (responses.GetUserResponse, error)
	CreateUser(ctx context.Context, userName string) (responses.CreateUserResponse, error)
	DeleteUser(ctx context.Context, userName string) (bool, error)
	UpdateUser() (responses.UpdateUserResponse, error)

	GetTeam(ctx context.Context, teamName string) (responses.GetTeamResponse, error)
	CreateTeam(ctx context.Context, costCenter string, teamName string) (responses.CreateTeamResponse, error)
	DeleteTeam(ctx context.Context, teamName string) (bool, error)
	UpdateTeam() (responses.UpdateTeamResponse, error)
}
