package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Restful-Perpustakaan-API/database" // Ganti dengan path yang sesuai
	"github.com/gorilla/mux"
)

// GetRecommendations mengambil rekomendasi buku umum dari database.
func GetRecommendations(w http.ResponseWriter, r *http.Request) {
	recommendations, err := database.GetRecommendations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recommendations)
}

// GetPersonalizedRecommendations mengambil rekomendasi buku yang dipersonalisasi berdasarkan ID anggota.
func GetPersonalizedRecommendations(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	memberID, err := strconv.Atoi(params["memberId"])
	if err != nil {
		http.Error(w, "Invalid member ID", http.StatusBadRequest)
		return
	}

	recommendations, err := database.GetPersonalizedRecommendations(memberID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recommendations)
}
