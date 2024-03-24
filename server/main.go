package main

import (
	"log"
	"net/http"
	"os"
	"spaceports-leaderboard/cmd"
	"spaceports-leaderboard/database"

	_ "github.com/lib/pq"
)

func main() {
	database.ConnectDb()

	routes := cmd.SetupRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":"+port, routes)
	if err != nil {
		log.Fatal(err)
	}
}