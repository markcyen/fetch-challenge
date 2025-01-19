package utils

import (
	"crypto/sha256"
	"encoding/json"
	"fetch-challenge/models"
	"fmt"
)

// Generate a hash for a receipt to check for duplicates
func GenerateReceiptHash(receipt models.Receipt) string {
	data, _ := json.Marshal(receipt)
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash)
}
