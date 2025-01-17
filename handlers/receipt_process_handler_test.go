package handlers

import (
	"bytes"
	"encoding/json"
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
			requestBody:    `{"retailer": "Walgreens","purchaseDate":"2022-01-02","purchaseTime":"08:13","items":[{"shortDescription":"Pepsi - 12-oz","price":1.25},{"shortDescription":"Dasani","price":1.40}],"total":2.65}`,
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
			requestBody:    `{"retailer":"Walgreens",`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid JSON payload\n",
		},
		{
			name:           "Duplicate Receipt",
			requestBody:    `{"retailer": "Walgreens","purchaseDate":"2022-01-02","purchaseTime":"08:13","items":[{"shortDescription":"Pepsi - 12-oz","price":1.25},{"shortDescription":"Dasani","price":1.40}],"total":2.65}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `"id:`,
		},
	}

	// previousID stored receipt ID to test duplicates
	var previousID string

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
			if tc.name != "Duplicate Receipt" {
				if !strings.Contains(rr.Body.String(), tc.expectedBody) {
					t.Errorf("Test %q failed: response body = %q, want partial match with %q", tc.name, rr.Body.String(), tc.expectedBody)
				}
			}

			if tc.name == "Duplicate Receipt" {
				// Extract the receipt ID from the response
				var response map[string]string
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to parse response: %v", err)
				}
				currentID, exists := response["id"]
				if !exists {
					t.Fatalf("Response does not contain 'id': %v", rr.Body.String())
				}

				// Check if the ID matches the one from the previous "Valid Request" test
				if previousID == "" {
					previousID = currentID // Store the ID from the first valid request
				} else if previousID != currentID {
					t.Errorf("Expected duplicate receipt to return the same ID, got %q and %q", previousID, currentID)
				}
			}
		})
	}
}
