package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"spaceports-leaderboard/database"
	"spaceports-leaderboard/models"
)

func ListLeaderboard(w http.ResponseWriter, r *http.Request) {
	limitParam := r.URL.Query().Get("limit")
	limit := 0

	if limitParam != "" {
		limit, _ = strconv.Atoi(limitParam)
		if limit <= 0 {
			http.Error(w, "Invalid limit value", http.StatusBadRequest)
			return
		}
	}

	leaderboard := []models.Leaderboard{}
	query := database.DB.Db

	if limit > 0 {
		query = query.Order("score desc").Limit(limit)
	}

	query.Find(&leaderboard)
    w.Header().Set("Content-Type", "application/json")
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
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(leaderboard)
}
