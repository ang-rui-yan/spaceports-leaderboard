package cmd

import (
	"encoding/json"
	"log"
	"net/http"

	"spaceports-leaderboard/database"

	"gorm.io/gorm"
)

func SetupRoutes() *http.ServeMux {
	routes := http.NewServeMux()

	routes.HandleFunc("/health", pingHandler(database.DB.Db))

	return routes
}


func pingHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbConn, err := db.DB()
		if err != nil {
			log.Fatal(err)
		}

		err = dbConn.Ping()
		if err != nil {
			log.Fatal(err)
		}

		response := map[string]string{
			"message": "Service is healthy",
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
	
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	}
}