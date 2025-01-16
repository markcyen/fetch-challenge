package handlers

import (
	"fetch-challenge/models"
	"net/http"
	"sync"
)

var (
	receipts = make(map[string]models.Receipt)
	mu       sync.Mutex
)

func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {

}
