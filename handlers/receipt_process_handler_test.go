package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProcessReceiptHandler(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid Request",
			requestBody:    `{"retailer":"Target","purchaseDate":"2025-01-15","purchaseTime":"12:00:00","items":[{"description":"Apples","price":2.99},{"description":"Bananas","price":1.49}],"total":4.48}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `"id":`,
		},
		{
			name:           "Empty Body",
			requestBody:    "",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid JSON payload\n",
		},
		{
			name:           "Invalid JSON",
			requestBody:    `{"retailer":"Target",`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid JSON payload\n",
		},
	}

	// Iterate over test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create request with the specified body
			req, err := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer([]byte(tc.requestBody)))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			// Record the response
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(ProcessReceiptHandler)
			handler.ServeHTTP(rr, req)

			// Assert the status code
			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("Test %q failed: got status %v, want %v", tc.name, status, tc.expectedStatus)
			}

			// Assert the response body (partial match for flexibility)
			if !strings.Contains(rr.Body.String(), tc.expectedBody) {
				t.Errorf("Test %q failed: response body = %q, want partial match with %q", tc.name, rr.Body.String(), tc.expectedBody)
			}
		})
	}
}
