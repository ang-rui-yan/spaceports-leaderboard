package cmd

import (
	"encoding/json"
	"log"
	"net/http"

	"spaceports-leaderboard/database"
	"spaceports-leaderboard/handlers"

	"gorm.io/gorm"
)

func SetupRoutes() *http.ServeMux {
	routes := http.NewServeMux()

	routes.HandleFunc("/health", pingHandler(database.DB.Db))
	routes.HandleFunc("/jwt", getJWTKey())
	routes.HandleFunc("/api/v1/scores", insertScoreHandler())
	routes.HandleFunc("/api/v1/leaderboard", viewLeaderboardHandler())
	
	return routes
}


func pingHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
            // If not, return a 405 Method Not Allowed error
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }

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


func insertScoreHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
            // If not, return a 405 Method Not Allowed error
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }
		handlers.InsertScore(w, r)
	}
}

func viewLeaderboardHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
            // If not, return a 405 Method Not Allowed error
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }

		handlers.ListLeaderboard(w, r)
	}
}

func getJWTKey() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			// If not, return a 405 Method Not Allowed error
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		handlers.GetJWTKey(w, r)
	}
}