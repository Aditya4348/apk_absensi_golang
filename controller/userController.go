package controller

import (
	"database/sql"
	"apk_absesnsi_harian/models"
	"encoding/json"
	"net/http"
)

func GetUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		user, err := models.GetUser(db)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "User not found", http.StatusNotFound)
			} else {
				http.Error(w, "Failed to get user", http.StatusInternalServerError)
			}
			return	
		}

		json.NewEncoder(w).Encode(user)
	}
}