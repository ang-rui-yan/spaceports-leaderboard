package main

import (
	"net/http"
	"spaceports-leaderboard/cmd"
	"spaceports-leaderboard/database"

	_ "github.com/lib/pq"
)

func main() {
	database.ConnectDb()
	
	routes := cmd.SetupRoutes()

	http.ListenAndServe(":5100", routes)
}