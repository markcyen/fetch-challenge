package handlers

import (
	"encoding/json"
	"fetch-challenge/services"
	"net/http"

	"github.com/gorilla/mux"
)

// GetPointsHandler handles the points response
// GET /receipts/{id}/points
func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the receipt ID from URL path (/receipts/{id}/points)
	vars := mux.Vars(r)
	id := vars["id"]

	// Check the existance of an id
	if id == "" {
		http.Error(w, "No receipt found for that ID.", http.StatusNotFound)
		return
	}

	// Look up receipt by id in the in-memory storage
	mu.Lock()
	receipt, exists := receipts[id]
	mu.Unlock()

	if !exists {
		http.Error(w, "No receipt found for that ID.", http.StatusNotFound)
		return
	}

	// Calculate points
	points := services.CalculatePoints(receipt)

	// Respond with the points
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"points": points})
}
