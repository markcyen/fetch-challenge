package handlers

import (
	"crypto/sha256"
	"encoding/json"
	"fetch-challenge/models"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync"

	"github.com/google/uuid"
)

var (
	receipts      = make(map[string]models.Receipt)
	receiptHashes = make(map[string]string)
	mu            sync.Mutex
)

// ProcessReceiptHandler takes a receipt and generates an id for it
// POST /receipts/process
func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "The receipt is invalid.", http.StatusBadRequest)
		return
	}

	if !isValidRetailer(receipt.Retailer) || !isValidCurrency(receipt.Total) || len(receipt.Items) < 1 || !isValidPurchaseDate(receipt.PurchaseDate) || !isValidPurchaseTime(receipt.PurchaseTime) {
		http.Error(w, "The receipt is invalid.", http.StatusBadRequest)
		return
	}

	for _, item := range receipt.Items {
		if !isValidShortDescription(item.ShortDescription) || !isValidCurrency(item.Price) {
			http.Error(w, "The receipt is invalid.", http.StatusBadRequest)
			return
		}
	}

	// Generate hash for receipt to capture any duplicates
	receiptHash := generateReceiptHash(receipt)

	// Store receipt id and details in in-memory storage
	mu.Lock()
	defer mu.Unlock()

	// Check if existing ID exists for a receipt that has already been processed
	if existingID, found := receiptHashes[receiptHash]; found {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"id": existingID})
		return
	}

	// Generate unique id for new receipt
	id := uuid.New().String()
	receipts[id] = receipt
	receiptHashes[receiptHash] = id

	log.Printf("\nReceipt saved: ID=%s, Hash: %s, Data=%+v\n", id, receiptHash, receipt)

	response := models.ReceiptResponse{ID: id}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "The receipt is invalid.", http.StatusBadRequest)
		return
	}
}

// isValidRetailer checks for specified regular expression pattern on the retailer name
// For instance, "Super-Market" and "Shop & Save" are valid, but "Super@Market!" is not
func isValidRetailer(retailer string) bool {
	pattern := "^[\\w\\s\\-&]+$"
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(retailer)
}

// isValidPurchaseDate checks for specified regular expression pattern on the purchase date
// For instance, "2022-01-17" is valid, but "2022-17-01" or "2022/01/17" is not
func isValidPurchaseDate(purchaseDate string) bool {
	pattern := "^\\d{4}-\\d{2}-\\d{2}$"
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(purchaseDate)
}

// isValidPurchaseTime checks for specified regular expression pattern on the purchase time
// For instance, "13:01" is valid, but " 13:01" or "7:12" is not
func isValidPurchaseTime(purchaseTime string) bool {
	pattern := "^([01]\\d|2[0-3]):[0-5]\\d$"
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(purchaseTime)
}

// isValidCurrency checks for specified regular expression pattern on total and price
// For instance, "1.23" is valid, but "1.2", "123", ".12", and "1.234" are not
func isValidCurrency(amount string) bool {
	pattern := "^\\d+\\.\\d{2}$"
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(amount)
}

// isValidShortDescription checks for specified regular expression pattern on the short description
// For instance, "Pepsi 12-oz PK" is valid, but "#Pepsi & Coke!" is not
func isValidShortDescription(shortDescription string) bool {
	pattern := `^[\w\s-]+$`
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(shortDescription)
}

// generateReceiptHash is a helper function to generate a hash 
// for a receipt to check for duplicates
func generateReceiptHash(receipt models.Receipt) string {
	data, _ := json.Marshal(receipt)
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash)
}
