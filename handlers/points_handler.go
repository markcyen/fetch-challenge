package handlers

import (
	"net/http"
)

func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	// Extract the receipt ID from URL path (/receipts/:id/points)

	// Check the existance of an id

	// Look up receipt by id in the in-memory storage

	// Calculate points

	// Respond with the points
}
