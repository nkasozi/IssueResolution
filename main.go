package main

import (
	"IssueResolution/api"
	"IssueResolution/db"
	"IssueResolution/logic"
	"IssueResolution/server/http"
)

func main() {
	//initialize the DB
	database := db.InitInMemoryDb()

	//initialize the logic which depends on the DB
	coreLogic := logic.Logic{Db: database}

	//initialize the endpoints which depends on the logic
	endpoints := api.ApiEndpoints{CoreLogic: &coreLogic}

	//initialize the server which hosts the endpoints
	httpServer := http.Server{Endpoints: endpoints}

	//run the server
	httpServer.Run(9090)
}
