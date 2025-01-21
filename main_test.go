package main

import (
	"bytes"
	"fetch-challenge/handlers"
	"fetch-challenge/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	router := setupRouter()

	testReceiptID := "123"
	handlers.Receipts[testReceiptID] = models.Receipt{
		Retailer:     "Walgreens",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "08:13",
		Items: []models.Item{
			{ShortDescription: "Pepsi 12PK", Price: "1.25"},
		},
		Total: "1.25",
	}

	tests := []struct {
		name           string
		method         string
		path           string
		requestBody    string
		expectedStatus int
	}{
		{
			name:           "Valid POST /receipts/process",
			method:         "POST",
			path:           "/receipts/process",
			requestBody:    `{"retailer":"Walgreens","purchaseDate":"2022-01-02","purchaseTime":"08:13","items":[{"shortDescription":"Pepsi 12PK","price":"1.25"}],"total":"1.25"}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid Method for /receipts/process",
			method:         "GET",
			path:           "/receipts/process",
			requestBody:    "",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Valid GET /receipts/{id}/points",
			method:         "GET",
			path:           "/receipts/123/points",
			requestBody:    "",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid Path",
			method:         "GET",
			path:           "/invalid/path",
			requestBody:    "",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "POST with Empty Body",
			method:         "POST",
			path:           "/receipts/process",
			requestBody:    "",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var req *http.Request
			var err error

			if tc.requestBody != "" {
				req, err = http.NewRequest(tc.method, tc.path, bytes.NewBuffer([]byte(tc.requestBody)))
			} else {
				req, err = http.NewRequest(tc.method, tc.path, bytes.NewBuffer([]byte("{}")))
			}
			assert.NoError(t, err, "Failed to create request")

			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedStatus, rr.Code, "Unexpected status code for %s", tc.name)
		})
	}
}
