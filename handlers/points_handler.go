package handlers

import (
	"encoding/json"
	"fetch-challenge/services"
	"net/http"
	"strings"
)

func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	// Extract the receipt ID from URL path (/receipts/:id/points)
	id := strings.TrimPrefix(r.URL.Path, "/receipts/")
	id = strings.TrimSuffix(id, "/points")

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
	points := services.CalculatePoints(receipt)

	// Respond with the points
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"points": points})
}
