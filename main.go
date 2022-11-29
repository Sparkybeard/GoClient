package main

import (
	"context"
	"fmt"

	"github.com/Sparkybeard/GoClient/client"
)

func main() {
	ops := client.Options{
		ApiUrl:  "http://localhost:5000/api/rpc",
		Verbose: true,
	}
	wmclient, err := client.NewClient(ops)
	if err != nil {
		return
	}
	ctx := context.Background()

	// createTeam, err := wmclient.CreateTeam(ctx, "costCenterTest", "testTeamName")
	// if err != nil {
	// 	return
	// }

	// fmt.Printf(createTeam.Data.TeamName)

	getTeam, err := wmclient.GetTeam(ctx, "testTeamName")
	if err != nil {
		return
	}

	fmt.Printf(getTeam.Data.CostCenter)

	booleanito, err := wmclient.DeleteTeam(ctx, "testTeamName")
	if err != nil {
		return
	}

	if booleanito {
		fmt.Printf("tteam2")
	}
}
