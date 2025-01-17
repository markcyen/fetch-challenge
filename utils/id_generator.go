package utils

import (
	"crypto/sha256"
	"encoding/json"
	"fetch-challenge/models"
	"fmt"
	"math/rand"
	"time"
)

// Generate a unique ID
func GenerateUniqueID() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%x", rand.Int63())
}

// Generate a hash for a receipt to check for duplicates
func GenerateReceiptHash(receipt models.Receipt) string {
	data, _ := json.Marshal(receipt)
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash)
}
