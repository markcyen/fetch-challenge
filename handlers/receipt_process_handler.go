package handlers

import (
	"encoding/json"
	"fetch-challenge/models"
	"fetch-challenge/utils"
	"log"
	"net/http"
	"sync"
)

var (
	receipts = make(map[string]models.Receipt)
	receiptHashes = make(map[string]string)
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

	// Generate hash for receipt to capture any duplicates
	receiptHash := utils.GenerateReceiptHash(receipt)

	// Store receipt id and details in in-memory storage
	mu.Lock()
	defer	mu.Unlock()

	// Check if existing ID exists for a receipt that has already been processed
	if existingID, found := receiptHashes[receiptHash]; found {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"id": existingID})
		return
	}

	// Generate unique id for new receipt
	id := utils.GenerateUniqueID()
	receipts[id] = receipt
	receiptHashes[receiptHash] = id

	log.Printf("\nReceipt saved: ID=%s, Hash: %s, Data=%+v\n", id, receiptHash, receipt)

	response := models.ReceiptResponse{ID: id}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
