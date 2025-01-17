package main

import (
	"fetch-challenge/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/receipts/process", handlers.ProcessReceiptHandler)
	http.HandleFunc("/receipts/:id/points", handlers.GetPointsHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Faild to start server: %v", err)
	}
}
