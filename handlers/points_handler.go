package handlers

import (
	"encoding/json"
	"fetch-challenge/services"
	"net/http"

	"github.com/gorilla/mux"
)

func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the receipt ID from URL path (/receipts/{id}/points)
	vars := mux.Vars(r)
	id := vars["id"]

	// Check the existance of an id
	if id == "" {
		http.Error(w, "Missing receipt ID", http.StatusBadRequest)
		return
	}

	// Look up receipt by id in the in-memory storage
	mu.Lock()
	receipt, exists := receipts[id]
	mu.Unlock()

	if !exists {
		http.Error(w, "Receipt ID not found", http.StatusNotFound)
		return
	}

	// Calculate points
	points, err := services.CalculatePoints(receipt)
	if err != nil {
		http.Error(w, "Error in points calculation", http.StatusUnprocessableEntity)
		return
	}

	// Respond with the points
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"points": points})
}
