package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// http.HandlerFunc("/receipts/process", handlers.ProcessReceiptHandler)
	// http.HandlerFunc("/receipts/", handlers.GetPointsHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Faild to start server: %v", err)
	}
}
