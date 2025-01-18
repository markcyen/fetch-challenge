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

	r.HandleFunc("/receipts/process", handlers.ProcessReceiptHandler).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", handlers.GetPointsHandler).Methods("GET")

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Faild to start server: %v", err)
	}
}
