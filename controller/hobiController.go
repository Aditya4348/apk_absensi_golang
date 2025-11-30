package controller

import (
	"apk_absesnsi_harian/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetHobiHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Menggunakan fungsi baru GetAllHobi yang mengembalikan slice.
		hobies, err := models.GetAllHobi(db)
		if err != nil {
			http.Error(w, "Failed to get hobi: "+err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(hobies)
	}
}

func GetHobiByIDHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Mengambil ID dari URL query parameter
		idsr := r.URL.Query().Get("id")
		if idsr == "" {
			http.Error(w, "ID parameter is required", http.StatusBadRequest)
			return
		}

		// Konversi ID dari string ke int64
		id, err := strconv.Atoi(idsr)
		if err != nil {
			http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
			return
		}

		hobi, err := models.GetHobiByID(db, int64(id))
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Hobi not found", http.StatusNotFound)
			} else {
				http.Error(w, "Failed to get hobi", http.StatusInternalServerError)
			}
			return
		}

		json.NewEncoder(w).Encode(hobi)
	}
}
