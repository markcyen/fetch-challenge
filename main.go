package main

import (
	"fetch-challenge/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	http.HandleFunc("/receipts/process", handlers.ProcessReceiptHandler)
	r.HandleFunc("/receipts/{id}/points", handlers.GetPointsHandler).Methods("GET")

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Faild to start server: %v", err)
	}
}
