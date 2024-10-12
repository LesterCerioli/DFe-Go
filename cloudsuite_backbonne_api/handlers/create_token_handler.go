package handlers

import (
	"cloudsuite_backbonne_api/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateTokenHandler(w http.ResponseWriter, r *http.Request) {
	var token models.Token

	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := token.Validate(); err != nil {
		http.Error(w, fmt.Sprintf("Validation failed: %v", err), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

func GetTokenHandler(w http.ResponseWriter, r *http.Request) {

	tokenValue := r.URL.Query().Get("value")

	if tokenValue == "" {
		http.Error(w, "Missing token value", http.StatusBadRequest)
		return
	}

	token := models.Token{
		Value:     tokenValue,
		IPAddress: "127.0.0.1", // Placeholder value
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}
