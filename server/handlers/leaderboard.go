package handlers

import (
	"encoding/json"
	"net/http"

	"spaceports-leaderboard/database"
	"spaceports-leaderboard/models"
)

func ListLeaderboard(w http.ResponseWriter, r *http.Request) {
	leaderboard := []models.Leaderboard{}
	database.DB.Db.Find(&leaderboard)

	json.NewEncoder(w).Encode(leaderboard)
}

func InsertScore(w http.ResponseWriter, r *http.Request) {
	leaderboard := new(models.Leaderboard)
	err := json.NewDecoder(r.Body).Decode(leaderboard)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	database.DB.Db.Create(&leaderboard)

	json.NewEncoder(w).Encode(leaderboard)
}
