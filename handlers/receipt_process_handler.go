package handlers

import (
	"encoding/json"
	"fetch-challenge/models"
	"fetch-challenge/utils"
	"net/http"
	"sync"
)

var (
	receipts = make(map[string]models.Receipt)
	mu       sync.Mutex
)

// ProcessReceiptHandler takes a receipt and generates an id for it
func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var receipt models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	id := utils.GenerateUniqueID()
	mu.Lock()
	receipts[id] = receipt
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}
