package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			requestBody:    `{"retailer": "Walgreens","purchaseDate":"2022-01-02","purchaseTime":"08:13","items":[{"shortDescription":"Pepsi 12PK","price":"1.25"},{"shortDescription":"Dasani","price":"1.40"}],"total":"2.65"}`,
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
			requestBody:    `{"retailer": "Walgreens","purchaseDate":"2022-01-02","purchaseTime":"08:13","items":[{"shortDescription":"Pepsi 12PK","price":"1.25"},{"shortDescription":"Dasani","price":"1.40"}],"total":"2.65"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `"id:`,
		},
	}

	// previousID stored receipt ID to test duplicates
	var previousID string

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer([]byte(tc.requestBody)))
			require.NoError(t, err, "Failed to create request")
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			ProcessReceiptHandler(rr, req)

			assert.Equal(t, tc.expectedStatus, rr.Code, "Unexpected status code for test case %q", tc.name)

			if tc.name != "Duplicate Receipt" {
				assert.Contains(t, rr.Body.String(), tc.expectedBody, "Response body does not contain expected substring for test case %q", tc.name)
			}

			if tc.name == "Duplicate Receipt" {
				var response map[string]string
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to parse response: %v", err)
				}
				currentID, exists := response["id"]
				require.True(t, exists, "Response does not contain 'id' for test case %q", tc.name)

				if previousID == "" {
					previousID = currentID 
				} else if previousID != currentID {
					t.Errorf("Expected duplicate receipt to return the same ID, got %q and %q", previousID, currentID)
				}
			}
		})
	}
}
